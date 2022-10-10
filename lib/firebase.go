package lib

import (
	"context"
	"encoding/json"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"qualifighting.backend.de/models"
)

var firebaseAuthClient *auth.Client

func initFirebase() (*auth.Client, error) {
	firebaseJSON := &models.Config{
		Type:                    appConfig.Firebase.Type,
		ProjectID:               appConfig.Firebase.ProjectID,
		PrivateKeyID:            appConfig.Firebase.PrivateKeyID,
		PrivateKey:              appConfig.Firebase.PrivateKey,
		ClientEmail:             appConfig.Firebase.ClientEmail,
		ClientID:                appConfig.Firebase.ClientID,
		AuthURI:                 appConfig.Firebase.AuthURI,
		TokenURI:                appConfig.Firebase.TokenURI,
		AuthProviderX509CertURL: appConfig.Firebase.AuthProviderX509CertURL,
		ClientX509CertURL:       appConfig.Firebase.ClientX509CertURL,
	}

	jsonByte, err := json.Marshal(&firebaseJSON)
	if err != nil {
		log.Fatalln(err)
	}

	opt := option.WithCredentialsJSON(jsonByte)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalln("Couldn't authorize Firebase: " + err.Error())
	}
	return auth, nil

}

func init() {
	var err error
	firebaseAuthClient, err = initFirebase()
	if err != nil {
		log.Fatalln(err)
	}

}

func GetFirebaseAuth() *auth.Client {
	return firebaseAuthClient
}
