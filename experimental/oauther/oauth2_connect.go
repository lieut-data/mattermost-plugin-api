// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package oauther

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"
	"golang.org/x/oauth2"
)

func (o *oAuther) oauth2Connect(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-ID")
	if userID == "" {
		o.logger.Debugf("oauth2Connect: reached by non authed user")
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	state := fmt.Sprintf("%v_%v", model.NewId()[0:15], userID)
	err := o.store.SetWithExpiry(o.getStateKey(userID), state, o.oAuth2StateTimeToLive)
	if err != nil {
		o.logger.Errorf("oauth2Connect: failed to store state, err=%s", err.Error())
		http.Error(w, "failed to store token state", http.StatusInternalServerError)
		return
	}

	redirectURL := o.config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(w, r, redirectURL, http.StatusFound)
}
