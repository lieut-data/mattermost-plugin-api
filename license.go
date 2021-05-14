package pluginapi

import (
	"github.com/mattermost/mattermost-server/v5/model"
)

// IsEnterpriseLicensedOrDevelopment returns true when the server is licensed with any Mattermost
// Enterprise License, or has `EnableDeveloper` and `EnableTesting` configuration settings
// enabled signaling a non-production, developer mode.
func IsEnterpriseLicensedOrDevelopment(config *model.Config, license *model.License) bool {
	if license != nil {
		return true
	}

	return IsConfiguredForDevelopment(config)
}

// IsE10LicensedOrDevelopment returns true when the server is licensed with a Mattermost
// Enterprise E10 License, or has `EnableDeveloper` and `EnableTesting` configuration settings
// enabled, signaling a non-production, developer mode.
func IsE10LicensedOrDevelopment(config *model.Config, license *model.License) bool {
	if license != nil &&
		license.Features != nil &&
		license.Features.LDAP != nil &&
		*license.Features.LDAP {
		return true
	}

	return IsConfiguredForDevelopment(config)
}

// IsE20LicensedOrDevelopment returns true when the server is licensed with a Mattermost
// Enterprise E20 License, or has `EnableDeveloper` and `EnableTesting` configuration settings
// enabled, signaling a non-production, developer mode.
func IsE20LicensedOrDevelopment(config *model.Config, license *model.License) bool {
	if license != nil &&
		license.Features != nil &&
		license.Features.FutureFeatures != nil &&
		*license.Features.FutureFeatures {
		return true
	}

	return IsConfiguredForDevelopment(config)
}

// IsConfiguredForDevelopment returns true when the server has `EnableDeveloper` and `EnableTesting`
// configuration settings enabled, signaling a non-production, developer mode.
func IsConfiguredForDevelopment(config *model.Config) bool {
	if config != nil &&
		config.ServiceSettings.EnableTesting != nil &&
		*config.ServiceSettings.EnableTesting &&
		config.ServiceSettings.EnableDeveloper != nil &&
		*config.ServiceSettings.EnableDeveloper {
		return true
	}

	return false
}
