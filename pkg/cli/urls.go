package cli

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func BuildAuthorizationRequest(config Config) string {
	return fmt.Sprintf("%v/realms/%v/protocol/openid-connect/auth?client_id=%v"+
		"&redirect_uri=%v&response_type=code&code_challenge_method=%v&code_challenge=%v",
		config.KeycloakConfig.KeycloakURL, config.KeycloakConfig.Realm, config.KeycloakConfig.ClientID,
		config.EmbeddedServerConfig.GetCallbackURL(), config.EmbeddedServerConfig.CodeChallengeMethod,
		config.EmbeddedServerConfig.CodeChallenge)
}

func BuildTokenExchangeRequest(config Config, code string, codeVerifier string) (*http.Request, error) {
	tokenURL := fmt.Sprintf("%v/realms/%v/protocol/openid-connect/token", config.KeycloakConfig.KeycloakURL, config.KeycloakConfig.Realm)

	body := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"client_id":     {config.KeycloakConfig.ClientID},
		"redirect_uri":  {config.EmbeddedServerConfig.GetCallbackURL()},
		"code_verifier": {codeVerifier},
	}
	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(body.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, err
}
