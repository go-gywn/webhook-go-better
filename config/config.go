package config

import (
	"flag"
	"fmt"
	"github.com/jinzhu/configor"
	"os"
	"webhook-better/helpers"
	"webhook-better/models"
)

// Config config
var Config = struct {
	Base     string `yaml:"base" default:"webhook"`
	Bind     string `yaml:"bind" default:"0.0.0.0:52802"`
	Timezone string `yaml:"timezone" default:"Asia/Seoul"`
	Key      string `yaml:"key" default:"03a73f3e7c9a7b38d196cd34c072567e"`
	Database struct {
		Host   string `yaml:"host" default:"127.0.0.1:3306"`
		User   string `yaml:"user" default:"dbadmin"`
		Pass   string `yaml:"pass" default:"l-6ILJ3Y6yahD7ibKwNe-t12rt1ahMUU6mI="`
		Schema string `yaml:"schema" default:"dbadmin"`
	} `yaml:"database" required`
	Webhooks map[string]models.Webhook
}{}

func InitConfig(cfg string) {
	var config, password string
	flag.StringVar(&config, "config", "configure.yml", "configuration")
	flag.StringVar(&password, "password", "", "password")
	flag.Parse()

	if err := configor.Load(&Config, config); err != nil {
		fmt.Println(config, "not exists, use default.")
	}

	if password != "" {
		fmt.Printf("<Encrypted>\n%s\n", helpers.CryptoHelper().EncryptAES(password, Config.Key))
		os.Exit(0)
	}

}
