package config

import (
	"fmt"

	"github.com/ErwinSalas/go-grpc-tls/pkg/gogrpctls"
	"github.com/spf13/viper"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSvcUrl    string `mapstructure:"AUTH_SVC_URL"`
	ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
	OrderSvcUrl   string `mapstructure:"ORDER_SVC_URL"`
	CertM         gogrpctls.CertManager
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	fmt.Println(c.AuthSvcUrl)
	fmt.Println(c.ProductSvcUrl)
	fmt.Println(c.OrderSvcUrl)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
