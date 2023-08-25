package main // import "github.com/dayvillefire/tenders"

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
	"strings"
	"sync"
	"syscall"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	_ "github.com/dayvillefire/tenders/api"
	"github.com/dayvillefire/tenders/auth"
	"github.com/dayvillefire/tenders/common"
	"github.com/dayvillefire/tenders/config"
	"github.com/dayvillefire/tenders/models"
	_ "github.com/dayvillefire/tenders/models"

	//"github.com/elastic/apm-agent-go/module/apmgin"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
)

var (
	//Apm        = flag.Bool("apm", false, "Use apm")
	ConfigFile = flag.String("config-file", "./tenders.yml", "App configuration file")
	Debug      = flag.Bool("debug", false, "Enable debugging (overrides config)")
	Daemonize  = flag.Bool("daemon", false, "Run as daemon")
	LogFile    = flag.String("log", "./tenders.log", "Log file (when run as daemon)")
	CPUProfile = flag.String("cpu-profile", "", "Write cpu profile to file")

	cacheStatusChan     = make(chan bool)
	cacheStatusQuitChan = make(chan bool)
	shutdownChannel     = make(chan os.Signal, 1)
	hostname            string
	Version             string
	VERSION             string

	htmlCache     map[string]htmlCacheEntry
	htmlCacheLock *sync.Mutex
)

type htmlCacheEntry struct {
	data   []byte
	expire time.Time
}

func init() {
	htmlCache = make(map[string]htmlCacheEntry)
	htmlCacheLock = new(sync.Mutex)
}

func main() {
	flag.Parse()

	Version = VERSION

	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)

	if *CPUProfile != "" {
		f, err := os.Create(*CPUProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if *Daemonize && *LogFile != "" {
		// Fix logging
		log.SetOutput(LockedWriter{&lumberjack.Logger{
			Filename:   *LogFile,
			MaxSize:    1024, // megabytes
			MaxBackups: 3,
			MaxAge:     28, //days
		}, &sync.Mutex{}})
	}

	c, err := config.LoadConfigWithDefaults(*ConfigFile)
	if err != nil {
		panic(err)
	}
	if c == nil {
		panic("UNABLE TO LOAD CONFIG")
	}
	config.Config = c

	if *Debug {
		log.Print("Overriding existing debug configuration")
		config.Config.Debug = true
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	common.InitDB()

	application()
}

func application() {
	//hostname, _ := os.Hostname()

	//oauth.InitializeOauth()

	log.Printf("Initializing web services")
	m := gin.Default()
	//m.Use(sessions.Sessions("tenderssession", oauth.Store))
	m.Use(gin.Logger())
	//if *Apm {
	//	m.Use(apmgin.Middleware(m))
	//} else {
	m.Use(gin.Recovery())
	//}

	// OAuth:
	//m.GET("/login", oauth.OAuthLoginHandler)
	//m.GET("/auth", oauth.OAuthHandler)

	// Enable gzip compression
	m.Use(gzip.Gzip(gzip.DefaultCompression))

	a := m.Group("/api")

	// Iterate through initializing API maps
	for k, v := range common.ApiMap {
		f := make([]string, 0)
		f = append(f, "AUTH")

		log.Printf("Adding handler /api/%s [%s]", k, strings.Join(f, ","))
		v(a.Group("/" + k))
	}

	// Redirection for short code URLs, based on whether we have an active session or not

	// Dashboard
	m.GET("/d/:shortcode", func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		if claims == nil || claims[auth.IdentityKey] == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/dashboard.html?code="+c.Param("shortcode"))
			return
		}
		_, err := models.UserByShortCode(claims[auth.IdentityKey].(string))
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/dashboard.html?code="+c.Param("shortcode"))
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, "/login.html?code="+c.Param("shortcode"))
	})

	// Redirect for login
	m.GET("/r/:shortcode", func(c *gin.Context) {
		// Propagate
		//c.SetCookie("shortcode", c.Param("shortcode"), 1, "/", "", false, false)

		claims := jwt.ExtractClaims(c)
		if claims == nil || claims[auth.IdentityKey] == nil {
			page, err := getPage("login.html")
			if err != nil {
				c.HTML(http.StatusInternalServerError, "", err.Error())
			}
			page = []byte(strings.Replace(
				string(page),
				`id="shortcode" value=""`,
				fmt.Sprintf(`id="shortcode" value="%s"`, c.Param("shortcode")),
				-1,
			),
			)
			//log.Printf("shortcode = %s, page = %s", c.Param("shortcode"), string(page))
			c.Data(http.StatusOK, "text/html", page)
			return
		}
		_, err := models.UserByShortCode(claims[auth.IdentityKey].(string))
		if err != nil {
			page, err := getPage("login.html")
			if err != nil {
				c.HTML(http.StatusInternalServerError, "", err.Error())
			}
			page = []byte(strings.Replace(
				string(page),
				`id="shortcode" value=""`,
				fmt.Sprintf(`id="shortcode" value="%s"`, c.Param("shortcode")),
				-1,
			),
			)
			c.Data(http.StatusOK, "text/html", page)
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, "/d/"+c.Param("shortcode"))
	})

	m.POST("/login", auth.GetAuthMiddleware().LoginHandler)

	log.Print("[static] Initializing with local resources")
	m.Use(static.Serve("/", static.LocalFile(config.Config.Paths.BasePath+string(os.PathSeparator)+"ui", false)))
	m.StaticFile("/", config.Config.Paths.BasePath+string(os.PathSeparator)+"ui"+string(os.PathSeparator)+"index.html")

	go func() {
		log.Printf("Initializing display on :%d", config.Config.Port)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), m); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		for {
			select {
			default:
				time.Sleep(time.Duration(5) * time.Second)
			case <-shutdownChannel:
				// stop
				return
			}
		}
	}()

	// Catch signals and termination
	signal.Notify(shutdownChannel, os.Interrupt)
	signal.Notify(shutdownChannel, syscall.SIGTERM)
	log.Println(<-shutdownChannel)
}

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		//c.Set("example", "12345")

		// before request
		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func getPage(path string) ([]byte, error) {
	htmlCacheLock.Lock()
	defer htmlCacheLock.Unlock()

	entry, found := htmlCache[path]
	if !found {
		b, err := ioutil.ReadFile(config.Config.Paths.BasePath + string(os.PathSeparator) + "ui" + string(os.PathSeparator) + path)
		if err != nil {
			return b, err
		}
		htmlCache[path] = htmlCacheEntry{data: b, expire: time.Now().Add(10 * time.Minute)}
		return b, err
	}

	if entry.expire.After(time.Now()) {
		b, err := ioutil.ReadFile(config.Config.Paths.BasePath + string(os.PathSeparator) + "ui" + string(os.PathSeparator) + path)
		if err != nil {
			return b, err
		}
		htmlCache[path] = htmlCacheEntry{data: b, expire: time.Now().Add(10 * time.Minute)}
		return b, err
	}

	return entry.data, nil
}
