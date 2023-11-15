package authentication

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/webauthn"
	"net/http"
)

func (api *API) BeginLogin(c *gin.Context) {
	username := c.Param("username")
	webAuthn, err := webauthn.New(getWebAuthnConfig())
	if err != nil {
		fmt.Println(err)
	}

	//user := datastore.GetUser() // Find the user
	user, err := api.AuthenticationService.GetUserAndCredentials(username)

	options, session, err := webAuthn.BeginLogin(user)
	if err != nil {
		// Handle Error and return.

		return
	}

	// store the session values
	api.AuthenticationService.SaveCredentialSession(user.ID, session)

	c.JSON(http.StatusOK, options)
}

func (api *API) FinishLogin(c *gin.Context) {
	username := c.Param("username")
	webAuthn, err := webauthn.New(getWebAuthnConfig())
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "Something went wrong"})
	}

	user, err := api.AuthenticationService.GetUserAndCredentials(username)

	// Get the session data stored from the function above
	session, err := api.AuthenticationService.GetCredentialSession(user.ID)

	// _ credential
	credential, err := webAuthn.FinishLogin(user, session, c.Request)
	if err != nil {
		// Handle Error and return.
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// If login was successful, handle next steps
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Login Success %s", base64.RawURLEncoding.EncodeToString(credential.ID))})
}
