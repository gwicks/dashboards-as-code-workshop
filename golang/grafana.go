package main

import (
	"net/url"
	"strings"

	"github.com/go-openapi/strfmt"
	gapi "github.com/grafana/grafana-openapi-client-go/client"
	"github.com/grafana/grafana-openapi-client-go/client/folders"
	"github.com/grafana/grafana-openapi-client-go/models"
)

func grafanaClient(cfg config) *gapi.GrafanaHTTPAPI {
	return gapi.NewHTTPClientWithConfig(strfmt.Default, &gapi.TransportConfig{
		// Host is the domain name or IP address of the host that serves the API.
		Host: cfg.GrafanaHost,
		// BasePath is the URL prefix for all API paths, relative to the host root.
		BasePath: "/api",
		// Schemes are the transfer protocols used by the API (http or https).
		Schemes: []string{"http"},
		// BasicAuth is optional basic auth credentials.
		// TODO: use the SA token from the config instead
		BasicAuth: url.UserPassword("admin", "admin"),
	})
}

func findOrCreateFolder(gapi *gapi.GrafanaHTTPAPI, folderName string) (string, error) {
	// TODO: this doesn't handle pagination.
	// It will misbehave if the target Grafana instance has >1000 folders.
	getParams := folders.NewGetFoldersParams()
	response, err := gapi.Folders.GetFolders(getParams)
	if err != nil {
		return "", err
	}

	for _, folder := range response.Payload {
		if strings.EqualFold(folder.Title, folderName) {
			return folder.UID, nil
		}
	}

	// The folder doesn't exist: create it.
	createResponse, err := gapi.Folders.CreateFolder(&models.CreateFolderCommand{
		Title: folderName,
	})
	if err != nil {
		return "", err
	}

	return createResponse.Payload.UID, nil
}
