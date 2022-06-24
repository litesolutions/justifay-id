package webhook

import (
	"github.com/gorilla/mux"
	"github.com/litesolutions/justifay-id/config"
	"github.com/litesolutions/justifay-id/oauth"
	"github.com/litesolutions/justifay-id/util/routes"
)

// ServiceInterface defines exported methods
type ServiceInterface interface {
	GetConfig() *config.Config
	GetOauthService() oauth.ServiceInterface
	GetRoutes() []routes.Route
	RegisterRoutes(router *mux.Router, prefix string)
	Close()
}
