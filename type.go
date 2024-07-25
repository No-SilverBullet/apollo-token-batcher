package apollotokenbatcher

type AppInfo struct {
	AppID                      string `json:"appId"`
	Name                       string `json:"name"`
	OrgId                      string `json:"orgId"`
	OrgName                    string `json:"orgName"`
	Id                         int32  `json:"id"`
	IsDeleted                  bool   `json:"isDeleted"`
	DataChangeCreatedBy        string `json:"dataChangeCreatedBy"`
	DataChangeCreatedTime      string `json:"dataChangeCreatedTime"`
	DataChangeLastModifiedBy   string `json:"dataChangeLastModifiedBy"`
	DataChangeLastModifiedTime string `json:"dataChangeLastModifiedTime"`
}

type GrantAppAccess2TokenRequest struct {
	AppID string `json:"appId"`
}
