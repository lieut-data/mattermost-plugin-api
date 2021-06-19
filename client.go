package pluginapi

import (
	"github.com/mattermost/mattermost-server/v5/plugin"
)

// Client is a streamlined wrapper over the mattermost plugin API.
type Client struct {
	api plugin.API

	Bot           BotService
	Configuration ConfigurationService
	Channel       ChannelService
	SlashCommand  SlashCommandService
	Emoji         EmojiService
	File          FileService
	Frontend      FrontendService
	Group         GroupService
	KV            KVService
	Log           LogService
	Mail          MailService
	Plugin        PluginService
	Post          PostService
	Session       SessionService
	Store         *StoreService
	System        SystemService
	Team          TeamService
	User          UserService
	AppsCache     AppsCacheService
}

// NewClient creates a new instance of Client.
//
// This client must only be created once per plugin to
// prevent reacquiring of resources, such as database connections.
func NewClient(api plugin.API) *Client {
	return &Client{
		api: api,

		Bot:           BotService{api: api},
		Channel:       ChannelService{api: api},
		Configuration: ConfigurationService{api: api},
		SlashCommand:  SlashCommandService{api: api},
		Emoji:         EmojiService{api: api},
		File:          FileService{api: api},
		Frontend:      FrontendService{api: api},
		Group:         GroupService{api: api},
		KV:            KVService{api: api},
		Log:           LogService{api: api},
		Mail:          MailService{api: api},
		Plugin:        PluginService{api: api},
		Post:          PostService{api: api},
		Session:       SessionService{api: api},
		Store:         &StoreService{api: api},
		System:        SystemService{api: api},
		Team:          TeamService{api: api},
		User:          UserService{api: api},
		AppsCache:     AppsCacheService{api: api},
	}
}
