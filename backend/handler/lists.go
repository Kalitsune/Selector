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

func returnFileContent(ctx *fiber.Ctx, srv *drive.Service, fileid string) error {

	list, err := api.IdToList(srv, fileid)
	if err != nil {
		if strings.HasPrefix(err.Error(), "googleapi: Error 404: File not found:") {
			return fiber.ErrNotFound
		}
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(list)
}

// GetLists is a handler for the /api/lists route that returns an array of the lists
func GetLists(ctx *fiber.Ctx) error {

	token := ctx.Locals("token").(*oauth2.Token)

	//get a Google Drive service
	srv, err := api.TokenToDriveClientService(token)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	//get the list of files
	fileList, err := srv.Files.List().Spaces("appDataFolder").Do()
	if err != nil {
		return fiber.ErrInternalServerError
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

	id := ctx.Params("id")
	token := ctx.Locals("token").(*oauth2.Token)

	//get a Google Drive service
	srv, err := api.TokenToDriveClientService(token)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return returnFileContent(ctx, srv, id)
}

// PostList is a handler for the /api/list/ route that takes a json list object and
// creates a new file in the user's Google Drive.
func PostList(ctx *fiber.Ctx) error {

	//get the token from the context
	token := ctx.Locals("token").(*oauth2.Token)

	//get a Google Drive service and handle potential server errors
	srv, err := api.TokenToDriveClientService(token)
	if err != nil {
		println("token2service error")
		return fiber.ErrInternalServerError
	}

	// Create a list from the go type structures.List, parse the json body and handle potential errors
	list := &structures.List{}
	err = list.FromJSON(ctx.Body())
	if err != nil {
		return fiber.ErrBadRequest
	}

	// Create a drive.File object from the list
	file := &drive.File{Name: list.Name, MimeType: "application/json", Parents: []string{"appDataFolder"}}

	// Marshal the list to json
	listData, err := list.ToJSON()
	if err != nil {
		return fiber.ErrInternalServerError
	}
	// create a reader from the json data
	reader := bytes.NewReader(listData)

	//save the file in the user's Google Drive and handle potential errors
	file, err = srv.Files.Create(file).Media(reader).Do()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	list.Id = file.Id

	//update the file on Google Drive to give it its ID
	reader = bytes.NewReader(listData)
	_, err = srv.Files.Update(list.Id, &drive.File{}).Media(reader).Do()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(list)
}

// PatchList is a handler for the /api/list/ route that takes a json list object and update it on the web
func PatchList(ctx *fiber.Ctx) error {
	// get the token from the context
	token := ctx.Locals("token").(*oauth2.Token)
	id := ctx.Params("id")

	//get a Google Drive service and handle potential server errors
	srv, err := api.TokenToDriveClientService(token)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	// Create a list from the go type structures.List, parse the json body and handle potential errors
	list := &structures.List{}
	err = list.FromJSON(ctx.Body())
	if err != nil {
		return fiber.ErrBadRequest
	}
	list.Id = id

	// Marshal the list to json
	listData, err := list.ToJSON()
	if err != nil {
		return fiber.ErrInternalServerError
	}
	// create a reader from the json data
	reader := bytes.NewReader(listData)

	//update the file on Google Drive and handle potential errors
	_, err = srv.Files.Update(list.Id, &drive.File{Name: list.Name}).Media(reader).Do()
	if err != nil {
		//handle 404 errors
		if strings.HasPrefix(err.Error(), "googleapi: Error 404: File not found:") {
			return fiber.ErrNotFound
		}
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(list)
}

// DeleteList is a handler for the /api/list/ route that takes a list id and deletes it from the user's Google Drive
func DeleteList(ctx *fiber.Ctx) error {
	// get the token from the context
	token := ctx.Locals("token").(*oauth2.Token)
	id := ctx.Params("id")

	//get a Google Drive service and handle potential server errors
	srv, err := api.TokenToDriveClientService(token)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	//delete the file from Google Drive and handle potential errors
	err = srv.Files.Delete(id).Do()
	if err != nil {
		//handle 404 errors
		if strings.HasPrefix(err.Error(), "googleapi: Error 404: File not found:") {
			return fiber.ErrNotFound
		}
		return fiber.ErrInternalServerError
	}

	return ctx.SendStatus(fiber.StatusOK)
}
