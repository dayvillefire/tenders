package config

import (
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	Config           *AppConfig
	cachedConfigPath string
)

type AppConfig struct {
	Debug bool `yaml:"debug"`
	Port  int  `yaml:"port"`
	Paths struct {
		BasePath string `yaml:"base-path"`
	} `yaml:"paths"`
	Database struct {
		User          string `yaml:"user"`
		Pass          string `yaml:"password"`
		Host          string `yaml:"host"`
		Name          string `yaml:"name"`
		Configuration string `yaml:"configuration"`
	} `yaml:"database"`
	Apm struct {
		Host           string   `yaml:"host"`
		Enabled        bool     `yaml:"enabled"`
		ServiceName    string   `yaml:"service-name"`
		Environment    string   `yaml:"environment"`
		IgnorePatterns []string `yaml:"ignore-patterns"`
		SanitizeFields []string `yaml:"sanitize-fields"`
	} `yaml:"apm"`
	Auth struct {
		ClientID     string `yaml:"id"`
		ClientSecret string `yaml:"secret"`
		StoreSecret  string `yaml:"store-secret"`
	} `yaml:"auth"`
}

func (c *AppConfig) SetDefaults() {
	c.Port = 8000
	c.Paths.BasePath = "."
	c.Apm.Host = "http://localhost:8200"
	c.Apm.Enabled = false
	c.Apm.ServiceName = "deployer"
	c.Apm.Environment = "production"
	c.Apm.IgnorePatterns = []string{}
	c.Apm.SanitizeFields = []string{"password"}
}

func LoadConfigWithDefaults(configPath string) (*AppConfig, error) {
	cachedConfigPath = configPath
	c := &AppConfig{}
	c.SetDefaults()
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal([]byte(data), c)
	if err != nil {
		return c, err
	}

	err = setApmEnvironmental(c)

	return c, err
}

func ConfigReload() error {
	c := &AppConfig{}
	c.SetDefaults()
	data, err := ioutil.ReadFile(cachedConfigPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(data), c)
	if err == nil {
		Config = c

		err = setApmEnvironmental(c)
	}
	return err
}

func setApmEnvironmental(c *AppConfig) error {
	var err error
	{
		err = os.Setenv("ELASTIC_APM_SERVER_URL", c.Apm.Host)
		if err != nil {
			return err
		}
	}
	{
		err = os.Setenv("ELASTIC_APM_SERVICE_NAME", c.Apm.ServiceName)
		if err != nil {
			return err
		}
	}
	{
		err = os.Setenv("ELASTIC_APM_ENVIRONMENT", c.Apm.Environment)
		if err != nil {
			return err
		}
	}
	if !c.Apm.Enabled {
		err = os.Setenv("ELASTIC_APM_ACTIVE", "false")
		if err != nil {
			return err
		}
	}
	if len(c.Apm.IgnorePatterns) > 0 {
		err = os.Setenv("ELASTIC_APM_IGNORE_URLS", strings.Join(c.Apm.IgnorePatterns, "|"))
		if err != nil {
			return err
		}
	}
	if len(c.Apm.SanitizeFields) > 0 {
		err = os.Setenv("ELASTIC_APM_SANITIZE_FIELD_NAMES", strings.Join(c.Apm.SanitizeFields, "|"))
		if err != nil {
			return err
		}
	}
	return nil
}
