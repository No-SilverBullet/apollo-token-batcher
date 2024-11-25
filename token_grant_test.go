package apollotokenbatcher

import (
	"testing"

	"github.com/No-SilverBullet/apollo-token-batcher/openapi"
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
		PortalAddr:   "http://apollo.com",
		Username:     "apollo",
		Password:     "*xDIivRZfvXieWaDb#DFeq&^RLoP4$Nn",
		OpenapiToken: "7cb47295f303214e2c966483448ebd515737a216",
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
	for i := 0; i < len(appInfos); i++ {
		if appInfos[i].IsDeleted {
			continue
		}
		c.GrantAppAccess2Token(&openapi.GrantAppAccess2TokenRequest{
			AppID: appInfos[i].AppID,
		})
		t.Logf("app【%s】] grant ok\n", appInfos[i].AppID)
	}
}

func TestCreateCluster(t *testing.T) {
	conf := &openapi.ApolloClientConfig{
		PortalAddr:   "http://apollo.com",
		Username:     "apollo",
		Password:     "*xDIivRZfvXieWaDb#DFeq&^RLoP4$Nn",
		OpenapiToken: "ea2ac40af3cd4212a3041e9240eac7aa1c4678e3",
	}
	c, err := openapi.NewApolloClient(conf)
	if err != nil {
		t.Fatal(err)
	}
	err = c.CreateCluster(&openapi.CreateClusterRequest{
		Name:  "stain-test",
		AppId: "aftersale-inner",
		Env:   "FAT",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteCluster(t *testing.T) {
	conf := &openapi.ApolloClientConfig{
		PortalAddr:   "http://apollo.com",
		Username:     "apollo",
		Password:     "*xDIivRZfvXieWaDb#DFeq&^RLoP4$Nn",
		OpenapiToken: "ea2ac40af3cd4212a3041e9240eac7aa1c4678e3",
	}
	c, err := openapi.NewApolloClient(conf)
	if err != nil {
		t.Fatal(err)
	}
	err = c.DeleteCluster(&openapi.DeleteClusterRequest{
		Name:  "stain-test",
		AppId: "aftersale-inner",
		Env:   "FAT",
	})
	if err != nil {
		t.Fatal(err)
	}
}
