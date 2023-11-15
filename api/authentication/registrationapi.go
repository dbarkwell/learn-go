package authentication

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
	"net/http"
)

func (api *API) cleanup(userID uuid.UUID) {
	api.AuthenticationService.RemoveCredentialSession(userID)
}

func (api *API) BeginRegistration(c *gin.Context) {
	username := c.Param("username")

	webAuthn, err := webauthn.New(getWebAuthnConfig())
	if err != nil {
		fmt.Println(err)
	}

	// TODO handle err
	user, err := api.AuthenticationService.GetUser(username)

	options, session, err := webAuthn.BeginRegistration(user)
	if session == nil {
		return
	}

	// handle errors if present
	// store the sessionData values

	api.AuthenticationService.SaveCredentialSession(user.ID, session)
	c.JSON(http.StatusOK, options) // return the options generated
	// options.publicKey contain our registration options
}

func (api *API) FinishRegistration(c *gin.Context) {
	username := c.Param("username")
	user, err := api.AuthenticationService.GetUserAndCredentials(username)
	defer api.cleanup(user.ID)

	webAuthn, err := webauthn.New(getWebAuthnConfig())
	if err != nil {
		fmt.Println(err)
	}

	session, err := api.AuthenticationService.GetCredentialSession(user.ID)

	credential, err := webAuthn.FinishRegistration(user, session, c.Request)
	if err != nil {
		// Handle Error and return.

		return
	}

	// If creation was successful, store the credential object
	// Pseudocode to add the user credential.
	//user.Credential = append(user.Credential, *credential)
	//datastore.SaveUser(user)
	var transport []string
	for _, v := range credential.Transport {
		transport = append(transport, string(v))
	}
	cred := CredentialDTO{
		ID:              credential.ID,
		PublicKey:       credential.PublicKey,
		AttestationType: credential.AttestationType,
		Transport:       transport,
		Flags: CredentialFlags{
			UserPresent:    credential.Flags.UserPresent,
			UserVerified:   credential.Flags.UserVerified,
			BackupEligible: credential.Flags.BackupEligible,
			BackupState:    credential.Flags.BackupState,
		},
		Authenticator: Authenticator{
			AAGUID:       credential.Authenticator.AAGUID,
			SignCount:    credential.Authenticator.SignCount,
			CloneWarning: credential.Authenticator.CloneWarning,
			Attachment:   string(credential.Authenticator.Attachment),
		},
	}

	api.AuthenticationService.SaveCredential(user.ID, cred)
	c.JSON(http.StatusOK, gin.H{"message": "Registration Success"}) // Handle next steps
}
