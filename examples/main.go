package examples

import (
	apollo_token "github.com/No-SilverBullet/apollo-token-batcher/openapi"
)

func main() {
	//init a apollo client
	c, err := apollo_token.NewApolloClient(&apollo_token.ApolloClientConfig{
		PortalAddr:   "http://apollo-xxx.com",
		Username:     "your_username",
		Password:     "your_password",
		OpenapiToken: "your_openapi_token",
	})
	if err != nil {
		panic(err)
	}
	//grant app access to openApi token
	err = c.GrantAppAccess2Token(&apollo_token.GrantAppAccess2TokenRequest{
		AppID: "xxx",
	})
	if err != nil {
		panic(err)
	}
	// get all app info
	appInfos, err := c.GetAllAppInfos()
	if err != nil {
		panic(err)
	}
	println("app total count :%d", len(appInfos))
	// grant ALL app access to openApi token
	for _, appInfo := range appInfos {
		err = c.GrantAppAccess2Token(&apollo_token.GrantAppAccess2TokenRequest{
			AppID: appInfo.AppID,
		})
		if err != nil {
			panic(err)
		}
	}
}
