package api

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"github.com/kalitsune/selector/structures"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"log"
	"net/http"
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
	config, err = google.ConfigFromJSON(b, drive.DriveAppdataScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
}

// GetAuthUrl returns the URL to the Google OAuth2 consent page.
func GetAuthUrl() string {
	return config.AuthCodeURL("state-token", oauth2.AccessTypeOnline)
}

// CodeToToken exchanges a Google authorisation code for a token (read: https://developers.google.com/identity/protocols/oauth2)
func CodeToToken(code string) (*oauth2.Token, error) {
	return config.Exchange(context.TODO(), code)
}

// tokenToClient creates an HTTP client from a token
func tokenToClient(token *oauth2.Token) *http.Client {
	return config.Client(context.Background(), token)
}

// driveClientService creates a Google Drive service from a token
func driveClientService(client *http.Client) (*drive.Service, error) {
	return drive.NewService(context.Background(), option.WithHTTPClient(client))
}

// TokenToDriveClientService creates a Google Drive service from a token
func TokenToDriveClientService(token *oauth2.Token) (*drive.Service, error) {
	client := tokenToClient(token)
	return driveClientService(client)
}

// TokenSerializer serializes a token to a string
func TokenSerializer(token *oauth2.Token) (string, error) {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(token)
	return base64.StdEncoding.EncodeToString(buffer.Bytes()), err
}

// TokenDeserializer deserializes a token from a string
func TokenDeserializer(str string) (*oauth2.Token, error) {
	token := &oauth2.Token{}
	serializedToken, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return token, err
	}

	buffer := bytes.NewBuffer(serializedToken)
	d := gob.NewDecoder(buffer)
	err = d.Decode(token)

	return token, err
}

// IdToList returns a list from a Google Drive file ID
func IdToList(srv *drive.Service, id string) (*structures.List, error) {
	//get the JSON file from Google Drive
	file, err := srv.Files.Get(id).Download()
	if err != nil {
		return &structures.List{}, err
	}

	//Unmarshal the JSON file to a list
	var list structures.List
	err = json.NewDecoder(file.Body).Decode(&list)
	if err != nil {
		return &structures.List{}, err
	}

	return &list, nil
}
