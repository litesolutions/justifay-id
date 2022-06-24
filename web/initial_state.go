package web

import (
	"github.com/litesolutions/justifay-api/model"
	"github.com/litesolutions/justifay-id/config"
	"github.com/litesolutions/justifay-id/session"
	"github.com/resonatecoop/user-api-client/models"
)

// Profile user public profile
type Profile struct {
	ID          string `json:"id"`
	Role        string `json:"role"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	// FullName               string                                 `json:"fullName"`
	// FirstName              string                                 `json:"firstName"`
	// LastName               string                                 `json:"lastName"`
	Country                string                                 `json:"country"`
	NewsletterNotification bool                                   `json:"newsletterNotification"`
	EmailConfirmed         bool                                   `json:"emailConfirmed"`
	Member                 bool                                   `json:"member"`
	Complete               bool                                   `json:"complete"`
	Usergroups             []*models.UserUserGroupPrivateResponse `json:"usergroups"`
}

// NewProfile
func NewProfile(
	user *model.User,
	usergroups []*models.UserUserGroupPrivateResponse,
	isUserAccountComplete bool,
	role string,
) *Profile {
	displayName := ""

	if len(usergroups) > 0 {
		displayName = usergroups[0].DisplayName
	}

	return &Profile{
		ID:                     user.ID.String(),
		Complete:               isUserAccountComplete,
		Country:                user.Country,
		DisplayName:            displayName,
		Role:                   role,
		Email:                  user.Username,
		EmailConfirmed:         user.EmailConfirmed,
		Member:                 user.Member,
		NewsletterNotification: user.NewsletterNotification,
		Usergroups:             usergroups,
	}
}

type InitialState struct {
	ApplicationName string                `json:"applicationName"`
	ClientID        string                `json:"clientID"`
	UserGroup       string                `json:"usergroup"`
	Token           string                `json:"token"`
	Clients         []config.ClientConfig `json:"clients"`
	Profile         *Profile              `json:"profile"`
	CSRFToken       string                `json:"csrfToken"`
	CountryList     []Country             `json:"countries"`
}

func NewInitialState(
	cnf *config.Config,
	client *model.Client,
	user *model.User,
	userSession *session.UserSession,
	isUserAccountComplete bool,
	usergroups []*models.UserUserGroupPrivateResponse,
	csrfToken string,
	countryList []Country,
) *InitialState {
	accessToken := ""

	if userSession != nil {
		accessToken = userSession.AccessToken
	}

	profile := NewProfile(
		user,
		usergroups,
		isUserAccountComplete,
		userSession.Role,
	)

	if len(usergroups) > 0 {
		profile.DisplayName = usergroups[0].DisplayName
	}

	return &InitialState{
		ApplicationName: client.ApplicationName.String,
		ClientID:        client.Key,
		Clients:         cnf.Clients,
		Profile:         profile,
		Token:           accessToken,
		CSRFToken:       csrfToken,
		CountryList:     countryList,
	}
}
