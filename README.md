

## Apollo-Token-Bacher

Used to grant application permissions to Apollo’s openapi token in batches；

## Usage

- First, you need to prepare Apollo’s account and password and the token used to call openApi.

[how to generate a openApi token](https://github.com/apolloconfig/apollo/wiki/Apollo%E5%BC%80%E6%94%BE%E5%B9%B3%E5%8F%B0)

- get all apollo app info and

```go
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
	// print app total count
	println("app total count :%d", len(appInfos))
}

```

So, to grant application permissions to Apollo’s openapi token in batches,u can fisrt get all applications info and then grant to these applications permission; 

```go
	// grant ALL app access to openApi token
	for _, appInfo := range appInfos {
		err = c.GrantAppAccess2Token(&apollo_token.GrantAppAccess2TokenRequest{
			AppID: appInfo.AppID,
		})
		if err != nil {
			panic(err)
		}
	}
```



If this is helpful for u, please give me a star,thanks~ 

