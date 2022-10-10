package lib

import (
	"github.com/JeremyLoy/config"
)

type Config struct {
	Firebase struct {
		Type                    string `config:"TYPE"`
		ProjectID               string `config:"PROJECT_ID"`
		PrivateKeyID            string `config:"PRIVATE_KEY_ID"`
		PrivateKey              string `config:"PRIVATE_KEY"`
		ClientEmail             string `config:"CLIENT_EMAIL"`
		ClientID                string `config:"CLIENT_ID"`
		AuthURI                 string `config:"AUTH_URI"`
		TokenURI                string `config:"TOKEN_URI"`
		AuthProviderX509CertURL string `config:"AUTH_PROVIDER_X509_CERT_URL"`
		ClientX509CertURL       string `config:"CLIENT_X509_CERT_URL"`
	}
	MongoDBURL   string `config:"MONGODB_URL"`
	Digitalocean struct {
		ID       string
		Secret   string
		Endpoint string
		Region   string
		Bucket   string
	}
	Port int
}

var appConfig Config

func init() {
	err := config.FromEnv().To(&appConfig)
	if err != nil {
		panic(err)
	}

	err = config.FromEnv().Sub(&appConfig.Firebase, "FIREBASE")
	if err != nil {
		panic(err)
	}

}

func GetBucketName() string {
	return appConfig.Digitalocean.Bucket
}

func GetAppConfig() Config {
	return appConfig
}
