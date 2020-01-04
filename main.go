package main // import "github.com/dayvillefire/tenders"

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
	"strings"
	"sync"
	"syscall"
	"time"

	_ "github.com/dayvillefire/tenders/api"
	"github.com/dayvillefire/tenders/common"
	"github.com/dayvillefire/tenders/config"
	_ "github.com/dayvillefire/tenders/models"
	"github.com/dayvillefire/tenders/oauth"

	//"github.com/elastic/apm-agent-go/module/apmgin"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/contrib/sessions"
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
)

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

	oauth.InitializeOauth()

	log.Printf("Initializing web services")
	m := gin.Default()
	m.Use(sessions.Sessions("tenderssession", oauth.Store))
	m.Use(gin.Logger())
	//if *Apm {
	//	m.Use(apmgin.Middleware(m))
	//} else {
	m.Use(gin.Recovery())
	//}

	// OAuth:
	m.GET("/login", oauth.OAuthLoginHandler)
	m.GET("/auth", oauth.OAuthHandler)

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
