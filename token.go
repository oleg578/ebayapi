package ebayapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	EXPIRETOKENERROR = "token expired"
)

type Token struct {
	AccessToken string    `json:"access_token"`
	ExpiresIn   int       `json:"expires_in"`
	TokenType   string    `json:"token_type"`
	ExpiredAt   time.Time `json:"-"`
}

type Error struct {
	Message          string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (t *Token) Get() error {
	const (
		MAGICSECOND = 10
	)
	appID, certID, errCr := getCredentials()
	if errCr != nil {
		return errCr
	}
	r, err := buildRequest(appID, certID)
	if err != nil {
		return err
	}
	c := &http.Client{}
	resp, errResp := c.Do(r)
	if errResp != nil {
		return errResp
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode != 200 {
		errP := &Error{}
		if err := errP.parseError(resp); err != nil {
			return err
		}
		return fmt.Errorf("%s: %s", errP.Message, errP.ErrorDescription)
	}
	errT := t.parseToken(resp)
	if errT != nil {
		return errT
	}
	//set ExpiredAt
	t.ExpiredAt = time.Now().Add(time.Duration(t.ExpiresIn-MAGICSECOND) * time.Second)
	return nil
}

func getCredentials() (appID, certID string, err error) {
	var ok bool
	appID, ok = os.LookupEnv(APPIDENV)
	if !ok {
		return "", "", fmt.Errorf(WRONGCREDENTIALS)
	}
	certID, ok = os.LookupEnv(CERTIDENV)
	if !ok {
		return "", "", fmt.Errorf(WRONGCREDENTIALS)
	}
	return
}

func buildRequest(appID, certID string) (r *http.Request, err error) {
	const (
		ROUTE = `identity/v1/oauth2/token`
	)
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("scope", "https://api.ebay.com/oauth/api_scope")
	r, err = http.NewRequest(http.MethodPost, APIURL+ROUTE, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.SetBasicAuth(appID, certID)
	return r, nil
}

func (t *Token) parseToken(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, t)
}

func (e *Error) parseError(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, e)
}
