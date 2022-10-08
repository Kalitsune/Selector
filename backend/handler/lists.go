package handler

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/kalitsune/selector/api"
	"github.com/kalitsune/selector/structures"
	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
	"strings"
)

func returnFileContent(ctx *fiber.Ctx, srv *drive.Service, fileId string) error {
	list, err := api.IdToList(srv, fileId)
	if err != nil {
		return handleErrors(err)
	}

	return ctx.JSON(list)
}

func handleErrors(err error) error {
	api.Logger.Warning.Printf("(Probably) A GoogleAPI Error has occurred: %v", err)

	//handle 404 errors
	if strings.HasPrefix(err.Error(), "googleapi: Error 404: File not found:") ||
		strings.HasPrefix(err.Error(), "googleapi: got HTTP response code 404") {
		return fiber.ErrNotFound
	}

	//handle invalid credentials errors
	if strings.HasPrefix(err.Error(), "googleapi: Error 401: Invalid Credentials") {
		return fiber.ErrUnauthorized
	}

	//handle rate limit errors
	if strings.HasPrefix(err.Error(), "googleapi: Error 403: User Rate Limit Exceeded") ||
		strings.HasPrefix(err.Error(), "googleapi: Error 429: Too Many Requests") {
		return fiber.ErrTooManyRequests
	}

	//handle storage quota errors
	if strings.HasPrefix(err.Error(), "googleapi: Error 403: Storage quota exceeded") {
		return fiber.ErrInsufficientStorage
	}

	//handle unauthorized errors
	if strings.HasPrefix(err.Error(), "googleapi: Error 403: The caller does not have permission") {
		return fiber.ErrForbidden
	}

	return fiber.ErrInternalServerError
}

// GetLists is a handler for the /api/lists route that returns an array of the lists
func GetLists(ctx *fiber.Ctx) error {

	api.Logger.Info.Println("handling GET request to /api/lists")

	token := ctx.Locals("token").(*oauth2.Token) // get the token unmarshalled in the api/apiMiddleware.go file

	//get a Google Drive service
	srv, err := api.TokenToDriveClientService(token)
	if err != nil {

		return fiber.ErrInternalServerError
	}

	//get the list of files
	fileList, err := srv.Files.List().Spaces("appDataFolder").Do()
	if err != nil {
		return handleErrors(err)
	}

	//convert the Google Drive files to lists
	var listArray []structures.List
	for _, file := range fileList.Files {
		listArray = append(listArray, structures.List{Id: file.Id, Name: file.Name})
	}

	return ctx.JSON(listArray)
}

// GetList is a handler for the /api/list/ route that returns a json containing
// the random elements lists when providing its ID.
func GetList(ctx *fiber.Ctx) error {
	//init the vars
	token := ctx.Locals("token").(*oauth2.Token) // get the token unmarshalled in the api/apiMiddleware.go file

	id := ctx.Params("id")
	api.Logger.Info.Printf("handling GET request to /api/list/%s", id)

	//get a Google Drive service
	srv, err := api.TokenToDriveClientService(token)
	if err != nil {
		return handleErrors(err)
	}

	return returnFileContent(ctx, srv, id)
}

// PostList is a handler for the /api/list/ route that takes a json list object and
// creates a new file in the user's Google Drive.
func PostList(ctx *fiber.Ctx) error {
	//get the token from the context
	token := ctx.Locals("token").(*oauth2.Token) // get the token unmarshalled in the api/apiMiddleware.go file

	api.Logger.Info.Println("handling POST request to /api/lists")

	//get a Google Drive service and handle potential server errors
	srv, err := api.TokenToDriveClientService(token)
	if err != nil {
		return handleErrors(err)
	}

	// Create a list from the go type structures.List, parse the json body and handle potential errors
	list := &structures.List{}
	err = list.FromJSON(ctx.Body())
	if err != nil {
		return fiber.ErrBadRequest
	}

	// Create a drive.File object from the list
	file := &drive.File{Name: list.Name, MimeType: "application/json", Parents: []string{"appDataFolder"}}

	//ensure that the list ID is empty
	list.Id = ""
	// Marshal the list to json
	listData, err := list.ToJSON()
	if err != nil {
		api.Logger.Error.Printf("handling POST request to /api/lists - Error marshalling list to json: %v", err)
		return fiber.ErrInternalServerError
	}
	// create a reader from the json data
	reader := bytes.NewReader(listData)

	//save the file in the user's Google Drive and handle potential errors
	file, err = srv.Files.Create(file).Media(reader).Do()
	if err != nil {
		return handleErrors(err)
	}
	api.Logger.Info.Printf("handling POST request to /api/lists - created file with id: %s", file.Id)

	list.Id = file.Id

	ctx.Status(fiber.StatusCreated)
	return ctx.JSON(list)
}

// PatchList is a handler for the /api/list/ route that takes a json list object and update it on the web
func PatchList(ctx *fiber.Ctx) error {
	// init the vars
	token := ctx.Locals("token").(*oauth2.Token) // get the token unmarshalled in the api/apiMiddleware.go file
	id := ctx.Params("id")
	api.Logger.Info.Printf("handling PATCH request to /api/list/%s", id)

	//get a Google Drive service and handle potential server errors
	srv, err := api.TokenToDriveClientService(token)
	if err != nil {
		return handleErrors(err)
	}

	// Create a list from the go type structures.List, parse the json body and handle potential errors
	list := &structures.List{}
	err = list.FromJSON(ctx.Body())
	if err != nil {
		return fiber.ErrBadRequest
	}

	// ensure that the id field is empty
	list.Id = ""
	// Marshal the list to json
	listData, err := list.ToJSON()
	if err != nil {
		api.Logger.Error.Printf("handling PATCH request to /api/list/%s - Error marshalling list to json: %v", id, err)
		return fiber.ErrInternalServerError
	}
	// create a reader from the json data
	reader := bytes.NewReader(listData)

	//update the file on Google Drive and handle potential errors
	_, err = srv.Files.Update(list.Id, &drive.File{Name: list.Name}).Media(reader).Do()
	if err != nil {
		return handleErrors(err)
	}

	api.Logger.Info.Printf("handling PATCH request to /api/list/%s - updated file with id: %s", id, list.Id)

	list.Id = id
	return ctx.JSON(list)
}

// DeleteList is a handler for the /api/list/ route that takes a list id and deletes it from the user's Google Drive
func DeleteList(ctx *fiber.Ctx) error {
	// init the vars
	token := ctx.Locals("token").(*oauth2.Token) // get the token unmarshalled in the api/apiMiddleware.go file
	id := ctx.Params("id")
	api.Logger.Info.Printf("handling DELETE request to /api/list/%s", id)

	//get a Google Drive service and handle potential server errors
	srv, err := api.TokenToDriveClientService(token)
	if err != nil {
		return handleErrors(err)
	}

	//delete the file from Google Drive and handle potential errors
	err = srv.Files.Delete(id).Do()
	if err != nil {
		return handleErrors(err)
	}

	api.Logger.Info.Printf("handling DELETE request to /api/list/%s - deleted file with id: %s", id, id)

	return ctx.SendStatus(fiber.StatusOK)
}
