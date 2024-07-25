

## Apollo-Token-Bacher

Used to grant application permissions to Apollo’s openapi token in batches；

## Usage

- First, you need to prepare Apollo’s account and password and the token used to call openApi.

[how to generate a openApi token](https://github.com/apolloconfig/apollo/wiki/Apollo%E5%BC%80%E6%94%BE%E5%B9%B3%E5%8F%B0)

- get all apollo app info 

```go
func TestGetAllAppInfo(t *testing.T) {
	conf := &openapi.ApolloClientConfig{
		PortalAddr:   "http://apollo-xxx.com",//there is no / at the end of the url
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
	
}
```

- grant application permissions to openapi token

```go
func TestGrantAppAccess2Token(t *testing.T) {
	conf := &openapi.ApolloClientConfig{
		PortalAddr:   "http://apollo-xxx.com",//there is no / at the end of the url
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
```

