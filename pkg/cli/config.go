package cli

import (
	"fmt"
	"github.com/nirasan/go-oauth-pkce-code-verifier"
)

type Config struct {
	KeycloakConfig       KeycloakConfig
	EmbeddedServerConfig EmbeddedServerConfig
}

type KeycloakConfig struct {
	KeycloakURL string
	Realm       string
	ClientID    string
}

type EmbeddedServerConfig struct {
	Port                uint32
	CallbackPath        string
	CodeChallengeMethod string
	CodeChallenge       string
	CodeVerifier        *go_oauth_pkce_code_verifier.CodeVerifier
}

func (c *EmbeddedServerConfig) GetCallbackURL() string {
	return fmt.Sprintf("http://localhost:%v/%v", c.Port, c.CallbackPath)
}
