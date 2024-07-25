package apollotokenbatcher

import (
	"apollo-token-batcher/openapi"
	"testing"
)

func TestGrantAppAccess2Token(t *testing.T) {
	conf := &openapi.ApolloClientConfig{
		PortalAddr:   "http://apollo-hw.qa.xgimi.com",
		Username:     "apollo",
		Password:     "*xDIivRZfvXieWaDb#DFeq&^RLoP4$Nn",
		OpenapiToken: "ea2ac40af3cd4212a3041e9240eac7aa1c4678e3",
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
		PortalAddr:   "http://apollo-hw.qa.xgimi.com",
		Username:     "apollo",
		Password:     "*xDIivRZfvXieWaDb#DFeq&^RLoP4$Nn",
		OpenapiToken: "ea2ac40af3cd4212a3041e9240eac7aa1c4678e3",
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
