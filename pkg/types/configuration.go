package types

import "fmt"

func ToStr(config Configuration) string {
	return fmt.Sprintf("Config is %#v", config)
}

type Configuration struct {
	Environment environment `mapstructure:"ENVIRONMENT"`
	Hosts       hosts       `mapstructure:"HOSTS"`
	Identity    identity    `mapstructure:"IDENTITY"`
	Ig          ig          `mapstructure:"IG"`
	Users       users       `mapstructure:"USERS"`
}

type hosts struct {
	BaseDomain           string `mapstructure:"BASE_DOMAIN"`
	RcsUiFQDN            string `mapstructure:"RCS_UI_FQDN"`
	IgFQDN               string `mapstructure:"IG_FQDN"`
	IdentityPlatformFQDN string `mapstructure:"IDENTITY_PLATFORM_FQDN"`
	Scheme               string `mapstructure:"SCHEME"`
}

type identity struct {
	AmRealm                    string `mapstructure:"AM_REALM"`
	IdmClientId                string `mapstructure:"IDM_CLIENT_ID"`
	IdmClientSecret            string `mapstructure:"IDM_CLIENT_SECRET"`
	RemoteConsentId            string `mapstructure:"REMOTE_CONSENT_ID"`
	ObriSoftwarePublisherAgent string `mapstructure:"OBRI_SOFTWARE_PUBLISHER_AGENT_NAME"`
	TestSoftwarePublisherAgent string `mapstructure:"TEST_SOFTWARE_PUBLISHER_AGENT_NAME"`
	ServiceAccountPolicy       string `mapstructure:"SERVICE_ACCOUNT_POLICY"`
}

type ig struct {
	IgClientId      string `mapstructure:"IG_CLIENT_ID"`
	IgClientSecret  string `mapstructure:"IG_CLIENT_SECRET"`
	IgRcsSecret     string `mapstructure:"IG_RCS_SECRET"`
	IgSsaSecret     string `mapstructure:"IG_SSA_SECRET"`
	IgIdmUser       string `mapstructure:"IG_IDM_USER"`
	IgIdmPassword   string `mapstructure:"IG_IDM_PASSWORD"`
	IgAgentId       string `mapstructure:"IG_AGENT_ID"`
	IgAgentPassword string `mapstructure:"IG_AGENT_PASSWORD"`
}
type environment struct {
	Verbose bool   `mapstructure:"VERBOSE"`
	Strict  bool   `mapstructure:"STRICT"`
	Type    string `mapstructure:"TYPE"`
	Paths   paths  `mapstructure:"PATHS"`
}

type paths struct {
	ConfigBaseDirectory    string `mapstructure:"CONFIG_BASE_DIRECTORY"`
	ConfigSecureBanking    string `mapstructure:"CONFIG_SECURE_BANKING"`
	ConfigIdentityPlatform string `mapstructure:"CONFIG_IDENTITY_PLATFORM"`
}

type users struct {
	CdmAdminUsername string `mapstructure:"FR_PLATFORM_ADMIN_USERNAME"`
	CdmAdminPassword string `mapstructure:"FR_PLATFORM_ADMIN_PASSWORD"`
}
