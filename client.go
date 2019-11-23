package pluginapi

import "github.com/mattermost/mattermost-server/v5/plugin"

// Client is a streamlined wrapper over the mattermost plugin API.
type Client struct {
	api plugin.API

	User UserService
	File *FileService
}

// NewClient creates a new instance of Client.
func NewClient(api plugin.API) *Client {
	return &Client{
		api:  api,
		User: UserService{api},
		File: &FileService{api},
	}
}
