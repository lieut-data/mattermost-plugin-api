package pluginapi

import "github.com/mattermost/mattermost-server/v5/model"

// normalizeAppError returns a truly nil error if appErr is nil.
//
// This doesn't happen automatically when a *model.AppError is cast to an error, since the
// resulting error interface has a concrete type with a nil value. This leads to the seemingly
// impossible:
//
//     var err error
//     err = func() *model.AppError { return nil }()
//     if err != nil {
//         panic("err != nil, which surprises most")
//     }
//
// Fix this problem for all plugin authors by normalizing to special case the handling of a nil
// *model.AppError. See https://golang.org/doc/faq#nil_error for more details.
func normalizeAppErr(appErr *model.AppError) error {
	if appErr == nil {
		return nil
	}
	return appErr
}
