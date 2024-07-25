package apollotokenbatcher

import "testing"

func TestGrantAppAccess2Token(t *testing.T) {
	conf := &ApolloClientConfig{
		PortalAddr:   "http://apollo-hw.qa.xgimi.com",
		Username:     "apollo",
		Password:     "*xDIivRZfvXieWaDb#DFeq&^RLoP4$Nn",
		OpenapiToken: "ea2ac40af3cd4212a3041e9240eac7aa1c4678e3",
	}
	c, err := New(conf)
	if err != nil {
		t.Fatal(err)
	}
	err = c.GrantAppAccess2Token(&GrantAppAccess2TokenRequest{
		AppID: "app-cms-service",
	})
	if err != nil {
		t.Fatal(err)
	}
}
func TestGetAllAppInfo(t *testing.T) {
	conf := &ApolloClientConfig{
		PortalAddr:   "http://apollo-hw.qa.xgimi.com",
		Username:     "apollo",
		Password:     "*xDIivRZfvXieWaDb#DFeq&^RLoP4$Nn",
		OpenapiToken: "ea2ac40af3cd4212a3041e9240eac7aa1c4678e3",
	}
	c, err := New(conf)
	if err != nil {
		t.Fatal(err)
	}
	appInfos, err := c.GetAllAppInfos()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("app count :%d", len(appInfos))
	// for i := 0; i < len(appInfos); i++ {
	// 	t.Logf("app info :%v", appInfos[i])
	// }
}
