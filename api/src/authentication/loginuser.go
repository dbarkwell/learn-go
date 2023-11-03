package authentication

import (
	"github.com/go-webauthn/webauthn/webauthn"
)

type LoginUser struct {
	ID          []byte
	Name        string
	DisplayName string
	Credential  []webauthn.Credential
}

var _ webauthn.User = (*LoginUser)(nil)

func (user *LoginUser) WebAuthnID() []byte {
	return user.ID
}

func (user *LoginUser) WebAuthnName() string {
	return user.Name
}

func (user *LoginUser) WebAuthnDisplayName() string {
	return user.DisplayName
}

func (user *LoginUser) WebAuthnCredentials() []webauthn.Credential {
	return user.Credential
}

func (user *LoginUser) WebAuthnIcon() string {
	return ""
}
