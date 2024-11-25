package openapi

import (
	"net/http"
	"time"

	validator "github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

// ApolloClient is the client for Apollo, which is used to grant the token access
type apolloClient struct {
	client       *http.Client
	portalAddr   string
	openapiToken string
	username     string
	password     string
}

// ApolloClientConfig is the configuration for ApolloClient
type ApolloClientConfig struct {
	PortalAddr   string `json:"portalAddr" validate:"required"`   // Apollo portal address,example: http://apollo.xxx.com
	Username     string `json:"username" validate:"required"`     // Apollo username
	Password     string `json:"password" validate:"required"`     // Apollo password
	OpenapiToken string `json:"openapiToken" validate:"required"` // Apollo openapi token
}

func (conf *ApolloClientConfig) Validate() error {
	return validate.Struct(conf)
}

// NewApolloClient creates a new ApolloClient
func NewApolloClient(conf *ApolloClientConfig) (*apolloClient, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}
	c := &apolloClient{
		client: &http.Client{
			Timeout: time.Second * 100,
		},
		portalAddr:   conf.PortalAddr,
		username:     conf.Username,
		password:     conf.Password,
		openapiToken: conf.OpenapiToken,
	}
	return c, nil
}
