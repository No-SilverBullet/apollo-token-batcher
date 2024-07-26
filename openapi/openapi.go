package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// return the cookies of apollo login
func (s *apolloClient) GetApolloCookies() ([]*http.Cookie, error) {
	loginUrl := fmt.Sprintf("%s/signin", s.portalAddr)
	req, err := http.NewRequest(http.MethodPost, loginUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(s.username, s.password)
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp.Cookies(), nil
}

// get all app info
func (s *apolloClient) GetAllAppInfos() ([]*AppInfo, error) {
	cookies, err := s.GetApolloCookies()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/apps", s.portalAddr), nil)
	if err != nil {
		return nil, err
	}
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var apps []*AppInfo
	if err := json.NewDecoder(resp.Body).Decode(&apps); err != nil {
		return nil, err
	}
	return apps, nil
}

// grant app access to openApi token
func (s *apolloClient) GrantAppAccess2Token(r *GrantAppAccess2TokenRequest) error {
	cookies, err := s.GetApolloCookies()
	if err != nil {
		return err
	}
	grantAccessUrl := fmt.Sprintf("%s/consumers/%s/assign-role?type=AppRole", s.portalAddr, s.openapiToken)
	postData, _ := json.Marshal(r)
	req, err := http.NewRequest(http.MethodPost, grantAccessUrl, bytes.NewReader(postData))
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		failRes := &GrantAppAccess2TokenFailedResponse{}
		err = json.NewDecoder(resp.Body).Decode(failRes)
		if err != nil {
			return err
		}
		return fmt.Errorf("failed to grant access, status code: %d,message: %s", resp.StatusCode, failRes.Message)
	}
	return nil
}
