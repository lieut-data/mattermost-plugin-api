package settings

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v5/model"
)

type emptySetting struct {
	baseSetting
}

func NewEmptySetting(id, title, description string) Setting {
	return &emptySetting{
		baseSetting: baseSetting{
			id:          id,
			title:       title,
			description: description,
		},
	}
}

func (s *emptySetting) GetSlackAttachments(userID, settingHandler string, disabled bool) (*model.SlackAttachment, error) {
	title := fmt.Sprintf("Setting: %s", s.title)
	sa := model.SlackAttachment{
		Title: title,
		Text:  s.description,
	}

	return &sa, nil
}
