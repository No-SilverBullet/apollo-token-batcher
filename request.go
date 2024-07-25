package apollotokenbatcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// return the cookies
func (s *ApolloClient) GetApolloCookies() ([]*http.Cookie, error) {
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

func (s *ApolloClient) GetAllAppInfos() ([]*AppInfo, error) {
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
func (s *ApolloClient) GrantAppAccess2Token(r *GrantAppAccess2TokenRequest) error {
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
	fmt.Printf("cookies:%v", cookies)
	req.Header.Add("content-type", "application/json")
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%s", string(body))
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to grant access, status code: %d", resp.StatusCode)
	}
	return nil
}
