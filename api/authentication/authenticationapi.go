package authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
	"net/http"
)

var (
	authnConfig AuthnConfig
)

func getWebAuthnConfig() *webauthn.Config {
	return &webauthn.Config{
		RPDisplayName: "Go Webauthn",                  // Display Name for your site
		RPID:          authnConfig.RPID,               // Generally the FQDN for your site
		RPOrigins:     []string{authnConfig.RPOrigin}, // The origin URLs allowed for WebAuthn requests
	}
}

func ProvideAuthenticationAPI(as Service, ac AuthnConfig) API {
	authnConfig = ac
	return API{AuthenticationService: as}
}

func (api *API) saveCredential(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("userID"))
	if err != nil {
	}

	var cred CredentialDTO
	if err := c.ShouldBindJSON(&cred); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	api.AuthenticationService.SaveCredential(userID, cred)
}

type API struct {
	AuthenticationService Service
}

type AuthnConfig struct {
	RPID     string
	RPOrigin string
}
