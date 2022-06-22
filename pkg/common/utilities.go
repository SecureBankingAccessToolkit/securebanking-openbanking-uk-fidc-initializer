package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"path/filepath"
	"secure-banking-uk-initializer/pkg/types"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

// Session containing the access token and cookie
type Session struct {
	authCode     string
	codeVerifier string
	AuthToken    AdminToken
	Cookie       *http.Cookie
}

// AdminToken returned by IDM
type AdminToken struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

var client = resty.New().SetRedirectPolicy(resty.NoRedirectPolicy()).SetError(RestError{})

// Authenticate user against platform, returns the iPlanetDomainPro cookie and access token
func (s *Session) Authenticate() (*http.Cookie, string) {
	s.GetIDMAdminAuthCode()
	s.GetIDMAdminToken()
	return s.Cookie, s.AuthToken.AccessToken
}

// GetIDMAdminAuthCode - get auth code from IDM.
// 		Redirects should be disabled, we expect a 302 status code here
func (s *Session) GetIDMAdminAuthCode() {
	zap.L().Info("Getting Identity Platform admin auth code")
	path := fmt.Sprintf("https://%s/am/oauth2/authorize", Config.Hosts.IdentityPlatformFQDN)
	// Use a generated verifier and challenge
	verifier, _ := CreateCodeVerifierWithLength(50)
	s.codeVerifier = verifier.String()
	codeChallenge := verifier.CodeChallengeS256()
	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetQueryParams(map[string]string{
			"redirect_uri":          fmt.Sprintf("https://%s/platform/appAuthHelperRedirect.html", Config.Hosts.IdentityPlatformFQDN),
			"client_id":             "idmAdminClient",
			"response_type":         "code",
			"scope":                 "fr:idm:*",
			"code_challenge":        codeChallenge,
			"code_challenge_method": "S256",
		}).
		SetCookie(s.Cookie).
		Get(path)
	if err != nil {
		if resp.StatusCode() != 302 {
			zap.S().Fatalw("Expecting 302 status code when getting auth code from Identity platform",
				"statusCode", resp.StatusCode(),
				"advice", "the idmAdminClient must exist in the root realm and redirects must be turned off",
				"error", resp.Error())
		}
	}
	v, err := url.ParseQuery(resp.Header().Get("Location"))
	if err != nil {
		zap.S().Fatalw("Error parsing location header", "statusCode", resp.StatusCode(), "error", err)
	}
	zap.S().Infow("Got Location header from IDM", "Location", v)
	authCode := v[fmt.Sprintf("https://%s/platform/appAuthHelperRedirect.html?code", Config.Hosts.IdentityPlatformFQDN)][0]
	s.authCode = authCode
}

// GetIDMAdminToken - get admin token from IDM
func (s *Session) GetIDMAdminToken() {
	zap.L().Info("Getting Identity Platform admin token")
	token := &AdminToken{}
	path := fmt.Sprintf("https://%s/am/oauth2/access_token", Config.Hosts.IdentityPlatformFQDN)
	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetResult(token).
		SetCookie(s.Cookie).
		SetFormData(map[string]string{
			"grant_type":    "authorization_code",
			"redirect_uri":  fmt.Sprintf("https://%s/platform/appAuthHelperRedirect.html", Config.Hosts.IdentityPlatformFQDN),
			"client_id":     "idmAdminClient",
			"code":          s.authCode,
			"code_verifier": s.codeVerifier,
		}).
		Post(path)

	RaiseForStatus(err, resp.Error(), resp.StatusCode())

	s.AuthToken = *token
}

//
// Unmarshals the file into v after substituting the values of config.
//
// Substitution is performed using syntax as described in: https://pkg.go.dev/text/template
//
func Unmarshal(file string, config *types.Configuration, v any) error {
	var err error
	var b []byte
	buf := bytes.NewBuffer(b)
	tmpl, err := template.New(filepath.Base(file)).ParseFiles(file)
	if err != nil {
		return err
	}
	if err = tmpl.Execute(buf, config); err != nil {
		return err
	}
	if err = json.Unmarshal(b, v); err != nil {
		return err
	}
	return err
}
