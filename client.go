package apollotokenbatcher

import (
	"net/http"
	"time"

	validator "github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

// ApolloClient is the client for Apollo, which is used to grant the token access
type ApolloClient struct {
	client       *http.Client
	portalAddr   string
	openapiToken string
	username     string
	password     string
}

// ApolloClientConfig is the configuration for ApolloClient
type ApolloClientConfig struct {
	PortalAddr   string `json:"portalAddr" validate:"required"`
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	OpenapiToken string `json:"openapiToken" validate:"required"`
}

func (conf *ApolloClientConfig) Validate() error {
	return validate.Struct(conf)
}

func New(conf *ApolloClientConfig) (*ApolloClient, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}
	c := &ApolloClient{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		portalAddr:   conf.PortalAddr,
		username:     conf.Username,
		password:     conf.Password,
		openapiToken: conf.OpenapiToken,
	}
	return c, nil
}
