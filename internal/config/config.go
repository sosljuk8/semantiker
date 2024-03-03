package config

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
)

type MYSQL struct {
	Db       string `env:"MYSQL_DATABASE,required"`
	User     string `env:"MYSQL_ROOT_USER,required"`
	Password string `env:"MYSQL_ROOT_PASSWORD,required"`
	Net      string `env:"MYSQL_NET,required"`
	Host     string `env:"MYSQL_HOST,required"`
	Port     string `env:"MYSQL_PORT,required"`
	Driver   string `env:"MYSQL_DRIVER,required"`
}

type CKey struct {
	KeyFilePath string `env:"KEY_FILE_PATH,required"`
	KeySheet    string `env:"KEY_SHEET,required"`
}

type Promo struct {
	DirPath      string `env:"PROMO_DIR_PATH,required"`
	Sheet        string `env:"PROMO_SHEET,required"`
	TemplatePath string `env:"PROMO_TEMPLATE_PATH,required"`
	StartRow     int    `env:"PROMO_START_ROW,required"`
}

type SemanterConfig struct {
	MySql MYSQL
	CKey  CKey
	Promo Promo
}

func GetConfig() *SemanterConfig {
	// Loading the environment variables from '.env' file.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
		return nil
	}

	cfg := SemanterConfig{}

	err = env.Parse(&cfg) // 👈 Parse environment variables into `Config`
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
		return nil
	}
	//fmt.Println(cfg)
	return &cfg
}
