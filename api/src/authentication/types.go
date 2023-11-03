package authentication

type credential struct {
	ID              []byte
	PublicKey       []byte
	AttestationType string
	Transport       []string
	Flags           credentialFlags
	Authenticator   authenticator
}

type credentialFlags struct {
	UserPresent    bool
	UserVerified   bool
	BackupEligible bool
	BackupState    bool
}

type authenticator struct {
	AAGUID       []byte
	SignCount    uint32
	CloneWarning bool
	Attachment   string
}
