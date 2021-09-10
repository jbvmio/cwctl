package cwctl

// AuthType represents an Authentication Type.
type AuthType int

const (
	AuthTypeLogin = iota
	AuthTypeToken
)

type Auth interface {
	AuthType() AuthType
}

// Credentials contain login details for a CW user.
type Credentials struct {
	Username          string
	Password          string
	TwoFactorPasscode string `json:"TwoFactorPasscode,omitempty"`
}

func (c *Credentials) AuthType() AuthType {
	return AuthTypeLogin
}

func (t *Token) AuthType() AuthType {
	return AuthTypeToken
}
