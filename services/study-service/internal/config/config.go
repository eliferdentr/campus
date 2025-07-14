package config

import (
	"github.com/spf13/viper"
)

// Config, uygulama için tüm yapılandırma değerlerini tutar.
type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
	DBDriver    string `mapstructure:"DB_DRIVER"`
	Environment string `mapstructure:"ENVIRONMENT"`
}

// Load, yapılandırmayı .env dosyasından ve ortam değişkenlerinden yükler.
func LoadConfig(path string) (config Config, err error)  {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	//override values that you read from config files
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return 
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return 
	}

	return 

}
