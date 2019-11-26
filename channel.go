package pluginapi

import (
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

// ChannelService exposes methods to manipulate channels.
type ChannelService struct {
	api plugin.API
}

// Get gets a channel.
//
// Minimum server version: 5.2
func (c *ChannelService) Get(channelID string) (*model.Channel, error) {
	channel, appErr := c.api.GetChannel(channelID)

	return channel, normalizeAppErr(appErr)
}

// GetByName gets a channel by its name, given a team id.
//
// Minimum server version: 5.2
func (c *ChannelService) GetByName(teamID, name string, includeDeleted bool) (*model.Channel, error) {
	channel, appErr := c.api.GetChannelByName(teamID, name, includeDeleted)

	return channel, normalizeAppErr(appErr)
}

// GetDirect gets a direct message channel.
//
// Note that if the channel does not exist it will create it.
//
// Minimum server version: 5.2
func (c *ChannelService) GetDirect(userID1, userID2 string) (*model.Channel, error) {
	channel, appErr := c.api.GetDirectChannel(userID1, userID2)

	return channel, normalizeAppErr(appErr)
}

// GetGroup gets a group message channel.
//
// Note that if the channel does not exist it will create it.
//
// Minimum server version: 5.2
func (c *ChannelService) GetGroup(userIDs []string) (*model.Channel, error) {
	channel, appErr := c.api.GetGroupChannel(userIDs)

	return channel, normalizeAppErr(appErr)
}

// GetByNameForTeamName gets a channel by its name, given a team name.
//
// Minimum server version: 5.2
func (c *ChannelService) GetByNameForTeamName(teamName, channelName string, includeDeleted bool) (*model.Channel, error) {
	channel, appErr := c.api.GetChannelByNameForTeamName(teamName, channelName, includeDeleted)

	return channel, normalizeAppErr(appErr)
}

// ListForTeamForUser gets a list of channels for given user ID in given team ID.
//
// Minimum server version: 5.6
func (c *ChannelService) ListForTeamForUser(teamID, userID string, includeDeleted bool) ([]*model.Channel, error) {
	channels, appErr := c.api.GetChannelsForTeamForUser(teamID, userID, includeDeleted)

	return channels, normalizeAppErr(appErr)
}

// ListPublicChannelsForTeam gets a list of all channels.
//
// Minimum server version: 5.2
func (c *ChannelService) ListPublicChannelsForTeam(teamID string, page, perPage int) ([]*model.Channel, error) {
	channels, appErr := c.api.GetPublicChannelsForTeam(teamID, page, perPage)

	return channels, normalizeAppErr(appErr)
}

// Search returns the channels on a team matching the provided search term.
//
// Minimum server version: 5.6
func (c *ChannelService) Search(teamID string, term string) ([]*model.Channel, error) {
	channels, appErr := c.api.SearchChannels(teamID, term)

	return channels, normalizeAppErr(appErr)
}

// Create creates a channel.
//
// Minimum server version: 5.2
func (c *ChannelService) Create(channel *model.Channel) (*model.Channel, error) {
	channel, appErr := c.api.CreateChannel(channel)

	return channel, normalizeAppErr(appErr)
}

// Update updates a channel.
//
// Minimum server version: 5.2
func (c *ChannelService) Update(channel *model.Channel) (*model.Channel, error) {
	channel, appErr := c.api.UpdateChannel(channel)

	return channel, normalizeAppErr(appErr)
}

// Delete deletes a channel.
//
// Minimum server version: 5.2
func (c *ChannelService) Delete(channelID string) error {
	appErr := c.api.DeleteChannel(channelID)

	return normalizeAppErr(appErr)
}

// GetChannelStats gets statistics for a channel.
//
// Minimum server version: 5.6
func (c *ChannelService) GetChannelStats(channelID string) (*model.ChannelStats, error) {
	channelStats, appErr := c.api.GetChannelStats(channelID)

	return channelStats, normalizeAppErr(appErr)
}

// GetMember gets a channel membership for a user.
//
// Minimum server version: 5.2
func (c *ChannelService) GetMember(channelID, userID string) (*model.ChannelMember, error) {
	channelMember, appErr := c.api.GetChannelMember(channelID, userID)

	return channelMember, normalizeAppErr(appErr)
}

// GetMembers gets a channel membership for all users.
//
// Minimum server version: 5.6
func (c *ChannelService) GetMembers(channelID string, page, perPage int) ([]*model.ChannelMember, error) {
	channelMembers, appErr := c.api.GetChannelMembers(channelID, page, perPage)

	return channelMembersToChannelMemberSlice(channelMembers), normalizeAppErr(appErr)
}

// GetMembersByIDs gets a channel membership for a particular User
//
// Minimum server version: 5.6
func (c *ChannelService) GetMembersByIDs(channelID string, userIDs []string) ([]*model.ChannelMember, error) {
	channelMembers, appErr := c.api.GetChannelMembersByIds(channelID, userIDs)

	return channelMembersToChannelMemberSlice(channelMembers), normalizeAppErr(appErr)
}

// GetMembersForUser returns all channel memberships on a team for a user.
//
// Minimum server version: 5.10
func (c *ChannelService) GetMembersForUser(teamID, userID string, page, perPage int) ([]*model.ChannelMember, error) {
	channelMembers, appErr := c.api.GetChannelMembersForUser(teamID, userID, page, perPage)

	return channelMembers, normalizeAppErr(appErr)
}

// AddMember joins a user to a channel (as if they joined themselves)
// This means the user will not receive notifications for joining the channel.
//
// Minimum server version: 5.2
func (c *ChannelService) AddMember(channelID, userID string) (*model.ChannelMember, error) {
	channelMember, appErr := c.api.AddChannelMember(channelID, userID)

	return channelMember, normalizeAppErr(appErr)
}

// AddUser adds a user to a channel as if the specified user had invited them.
// This means the user will receive the regular notifications for being added to the channel.
//
// Minimum server version: 5.18
func (c *ChannelService) AddUser(channelID, userID, asUserID string) (*model.ChannelMember, error) {
	channelMember, appErr := c.api.AddUserToChannel(channelID, userID, asUserID)

	return channelMember, normalizeAppErr(appErr)
}

// DeleteMember deletes a channel membership for a user.
//
// Minimum server version: 5.2
func (c *ChannelService) DeleteMember(channelID, userID string) error {
	appErr := c.api.DeleteChannelMember(channelID, userID)

	return normalizeAppErr(appErr)
}

// UpdateChannelMemberRoles updates a user's roles for a channel.
//
// Minimum server version: 5.2
func (c *ChannelService) UpdateChannelMemberRoles(channelID, userID, newRoles string) (*model.ChannelMember, error) {
	channelMember, appErr := c.api.UpdateChannelMemberRoles(channelID, userID, newRoles)

	return channelMember, normalizeAppErr(appErr)
}

// UpdateChannelMemberNotifications updates a user's notification properties for a channel.
//
// Minimum server version: 5.2
func (c *ChannelService) UpdateChannelMemberNotifications(channelID, userID string, notifications map[string]string) (*model.ChannelMember, error) {
	channelMember, appErr := c.api.UpdateChannelMemberNotifications(channelID, userID, notifications)

	return channelMember, normalizeAppErr(appErr)
}

func channelMembersToChannelMemberSlice(cm *model.ChannelMembers) []*model.ChannelMember {
	if cm == nil {
		return nil
	}

	cmp := make([]*model.ChannelMember, len(*cm))
	for i := 0; i < len(*cm); i++ {
		cmp[i] = &(*cm)[i]
	}
	return cmp
}