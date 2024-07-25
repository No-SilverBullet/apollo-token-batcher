package apollotokenbatcher

import (
	"apollo-token-batcher/openapi"
	"testing"
)

func TestGrantAppAccess2Token(t *testing.T) {
	conf := &openapi.ApolloClientConfig{
		PortalAddr:   "http://apollo-xxx.com", //there is no / at the end of the url
		Username:     "your_username",
		Password:     "your_password",
		OpenapiToken: "your_openapi_token",
	}
	c, err := openapi.NewApolloClient(conf)
	if err != nil {
		t.Fatal(err)
	}
	err = c.GrantAppAccess2Token(&openapi.GrantAppAccess2TokenRequest{
		AppID: "xxx",
	})
	if err != nil {
		t.Fatal(err)
	}
}
func TestGetAllAppInfo(t *testing.T) {
	conf := &openapi.ApolloClientConfig{
		PortalAddr:   "http://apollo-xxx.com", //there is no / at the end of the url
		Username:     "your_username",
		Password:     "your_password",
		OpenapiToken: "your_openapi_token",
	}
	c, err := openapi.NewApolloClient(conf)
	if err != nil {
		t.Fatal(err)
	}
	appInfos, err := c.GetAllAppInfos()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("app total count :%d", len(appInfos))
	// for i := 0; i < len(appInfos); i++ {
	// 	t.Logf("app info :%v", appInfos[i])
	// }
}
