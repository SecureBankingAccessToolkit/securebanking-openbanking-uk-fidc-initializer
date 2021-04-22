package serviceaccount

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/am"
	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var client = resty.New().SetRedirectPolicy(resty.NoRedirectPolicy()).SetError(common.RestError{})

// CreateIGServiceUser -
func CreateIGServiceUser() {
	zap.L().Debug("Creating IG service user")

	user := &common.ServiceUser{
		UserName:  viper.GetString("IG_IDM_USER"),
		SN:        "Service Account",
		GivenName: "IG",
		Mail:      "ig@acme.com",
		Password:  viper.GetString("OPEN_AM_PASSWORD"),
		AuthzRole: []common.AuthzRole{
			{
				Ref: "internal/role/openidm-admin",
			},
		},
	}
	path := "/openidm/managed/user/?_action=create"
	s := am.Client.Post(path, user, map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	})

	zap.S().Infow("IG Service User", "statusCode", s)
}

// CreateIGOAuth2Client -
func CreateIGOAuth2Client() {
	zap.L().Debug("Creating IG OAuth2 client")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "ig-oauth2-client.json")
	if err != nil {
		panic(err)
	}

	oauth2Client := &OAuth2Client{}
	err = json.Unmarshal(b, oauth2Client)
	if err != nil {
		panic(err)
	}
	oauth2Client.CoreOAuth2ClientConfig.Userpassword = "password"
	path := "/am/json/alpha/realm-config/agents/OAuth2Client/" + viper.GetString("IG_CLIENT_ID")
	s := am.Client.Put(path, oauth2Client, map[string]string{
		"Accept":           "application/json",
		"Content-Type":     "application/json",
		"Connection":       "keep-alive",
		"X-Requested-With": "ForgeRock Identity Cloud Postman Collection",
	})

	zap.S().Infow("IG OAuth2 Client", "statusCode", s)
}

// CreateIGPolicyAgent -
func CreateIGPolicyAgent() {
	zap.L().Debug("Creating IG Policy agent")
	policyAgent := &PolicyAgent{
		Userpassword: "password",
		IgTokenIntrospection: IgTokenIntrospection{
			Value:     "Realm",
			Inherited: false,
		},
	}
	path := "/am/json/alpha/realm-config/agents/IdentityGatewayAgent/ig-agent"
	s := am.Client.Put(path, policyAgent, map[string]string{
		"Accept":           "application/json",
		"Content-Type":     "application/json",
		"Connection":       "keep-alive",
		"X-Requested-With": "ForgeRock Identity Cloud Postman Collection",
	})

	zap.S().Infow("IG Policy Agent", "statusCode", s)
}

func CreateIDMAdminClient(cookie *http.Cookie) {
	zap.L().Debug("Creating IDM admin oauth2 client")
	b, err := ioutil.ReadFile(viper.GetString("REQUEST_BODY_PATH") + "idm-admin-client.json")
	if err != nil {
		panic(err)
	}
	config := &OAuth2Client{}
	json.Unmarshal(b, config)
	var redirect string
	for _, uri := range config.CoreOAuth2ClientConfig.RedirectionUris.Value {
		redirect = strings.ReplaceAll(uri, "{{IAM_FQDN}}", viper.GetString("IAM_FQDN"))
	}
	config.CoreOAuth2ClientConfig.RedirectionUris.Value = []string{redirect}
	zap.S().Debugw("Admin client request", "body", config)
	path := "https://" + viper.GetString("IAM_FQDN") + "/am/json/realm-config/agents/OAuth2Client/idmAdminClient"
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "keep-alive").
		SetHeader("X-Requested-With", "ForgeRock Identity Cloud Postman Collection").
		SetContentLength(true).
		SetCookie(cookie).
		SetBody(config).
		Put(path)

	common.RaiseForStatus(err, resp.Error())

	zap.S().Infow("IDM Admin Client", "statusCode", resp.StatusCode(), "redirect", config.CoreOAuth2ClientConfig.RedirectionUris.Value)
}
