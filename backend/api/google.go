package api

import (
	"context"
	"golang.org/x/oauth2"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

var config = &oauth2.Config{}

func init() {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err = google.ConfigFromJSON(b, drive.DriveMetadataReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
}

// GetAuthUrl returns the URL to the Google OAuth2 consent page.
func GetAuthUrl() string {
	return config.AuthCodeURL("state-token", oauth2.AccessTypeOnline)
}

func CodeToToken(code string) (*oauth2.Token, error) {
	return config.Exchange(context.TODO(), code)
}
