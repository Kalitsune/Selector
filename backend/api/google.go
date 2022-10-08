package api

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/kalitsune/selector/structures"
	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"net/http"
	"time"
)

var config = &oauth2.Config{}

// GetAuthUrl returns the URL to the Google OAuth2 consent page.
func GetAuthUrl() string {
	return config.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
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
	Logger.Info.Printf("Downloading file [%s] from Google Drive", id)
	file, err := srv.Files.Get(id).Download()
	if err != nil {
		return &structures.List{}, err
	}

	//Unmarshal the JSON file to a list
	var list structures.List
	err = json.NewDecoder(file.Body).Decode(&list)
	if err != nil {
		Logger.Error.Printf("Error while unmarshalling file [%s] from Google Drive: %v", id, err)
		return &structures.List{}, err
	}

	return &list, nil
}

// RefreshToken refreshes a token
func RefreshToken(token *oauth2.Token) (*oauth2.Token, error) {
	return config.TokenSource(context.TODO(), token).Token()
}

// SaveToken saves a token as a serialized cookie
func SaveToken(ctx *fiber.Ctx, token *oauth2.Token) error {
	//serialize the token
	serializedToken, err := TokenSerializer(token)
	if err != nil {
		return err
	}

	//store the cookie
	cookie := &fiber.Cookie{
		Name:    "token",
		Value:   serializedToken,
		Expires: time.Now().Add(time.Hour * 24 * 5), // equivalent to 5d
	}
	ctx.Cookie(cookie)

	return nil
}
