package oauth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dayvillefire/tenders/common"
	"github.com/dayvillefire/tenders/config"
	"github.com/dayvillefire/tenders/models"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Credentials are google credentials
type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

// User is a retrieved and authenticated user
type User struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Gender        string `json:"gender"`
}

var (
	cred  Credentials
	conf  *oauth2.Config
	Store sessions.CookieStore
)

// Initialize oauth
func InitializeOauth() {
	Store = sessions.NewCookieStore([]byte(config.Config.Auth.StoreSecret))
	conf = &oauth2.Config{
		ClientID:     config.Config.Auth.ClientID,
		ClientSecret: config.Config.Auth.ClientSecret,
		RedirectURL:  "http://localhost:8000/auth",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func OAuthLoginHandler(c *gin.Context) {
	state := randToken()
	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()
	c.Redirect(http.StatusTemporaryRedirect, getLoginURL(state))
}

func OAuthHandler(c *gin.Context) {
	session := sessions.Default(c)
	retrievedState := session.Get("state")
	if retrievedState != c.Query("state") {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retrievedState))
		return
	}
	token, err := conf.Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	client := conf.Client(oauth2.NoContext, token)
	email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer email.Body.Close()
	data, _ := ioutil.ReadAll(email.Body)
	log.Println("Email body: ", string(data))
	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Make sure local record exist here!
	if !models.UserExistsByEmail(user.Email) {
		obj := models.User{
			LoginEmail: user.Email,
			FirstName:  user.GivenName,
			LastName:   user.FamilyName,
			BitField:   models.UserBitActive,
		}
		_, err = common.DB.ValidateAndCreate(&obj)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		log.Printf("Created user : %#v", obj)
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
}
