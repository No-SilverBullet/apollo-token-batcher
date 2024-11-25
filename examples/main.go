package examples

import (
	"fmt"

	apollo_token "github.com/No-SilverBullet/apollo-token-batcher/openapi"
)

func main() {
	//init a apollo client
	c, err := apollo_token.NewApolloClient(&apollo_token.ApolloClientConfig{
		PortalAddr:   "http://apollo-prd.xgimi.com",
		Username:     "apollo",
		Password:     "*xDIivRZfvXieWaDb#DFeq&^RLoP4$Nn",
		OpenapiToken: "7cb47295f303214e2c966483448ebd515737a216",
	})
	if err != nil {
		panic(err)
	}
	//grant app access to openApi token
	// err = c.GrantAppAccess2Token(&apollo_token.GrantAppAccess2TokenRequest{
	// 	AppID: "xxx",
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// get all app info
	appInfos, err := c.GetAllAppInfos()
	if err != nil {
		panic(err)
	}
	println("app total count :%d", len(appInfos))
	// grant ALL app access to openApi token
	for index, appInfo := range appInfos {
		if appInfo.IsDeleted {
			continue
		}
		err = c.GrantAppAccess2Token(&apollo_token.GrantAppAccess2TokenRequest{
			AppID: appInfo.AppID,
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("app 【%d】【%s】] grant ok\n", index, appInfo.AppID)
	}
}
