package openapi

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

type GrantAppAccess2TokenFailedResponse struct {
	Exception string `json:"exception"`
	Message   string `json:"message"`
	Status    int32  `json:"status"`
	TimeStamp string `json:"timeStamp"`
}

type CreateClusterRequest struct {
	Name                string `json:"name"`
	AppId               string `json:"appId"`
	DataChangeCreatedBy string `json:"dataChangeCreatedBy"`
	Env                 string `json:"env"`
}

type DeleteClusterRequest struct {
	Name  string `json:"name"`
	AppId string `json:"appId"`
	Env   string `json:"env"`
}
