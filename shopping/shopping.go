package shopping

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"ebayapi"
)

//Include applicable values:
//•   Details
//Include most available fields in the response (except fields that significantly affect the call's performance).
//•   Description
//Include the Description field in the response. (This can affect the call's performance.)
//•   TextDescription
//Include the text Description (no html tag) field in the response. (This can affect the call's performance.)
//•   ShippingCosts
//Include basic shipping costs in the response. (Use GetShippingCosts to retrieve more details.)
//•   ItemSpecifics
//Include ItemSpecifics in the response.
//•   Variations
//Include Variations in the response.
//•   Compatibility
//Include Compatibility in the response.

//goland:noinspection GoUnusedExportedFunction
func GetSingleItem(itm string, includeSelector []string, t *ebayapi.Token) (item Item, apiErrors []ErrorMsg, err error) {
	if time.Now().Unix() > t.ExpiredAt.Unix() {
		return item, apiErrors, fmt.Errorf(ebayapi.EXPIRETOKENERROR)
	}
	r, errR := buildRequest(itm, includeSelector, t)
	if errR != nil {
		return item, apiErrors, errR
	}
	c := &http.Client{}
	resp, errC := c.Do(r)
	if errC != nil {
		return item, apiErrors, errC
	}
	if resp.StatusCode != http.StatusOK {
		return item, apiErrors, fmt.Errorf(resp.Status)
	}
	return parseResponse(resp)
}

func parseResponse(resp *http.Response) (item Item, apiErrors []ErrorMsg, err error) {
	apiResp := &APIResponse{}
	body, errB := io.ReadAll(resp.Body)
	if errB != nil {
		return item, nil, errB
	}
	err = json.Unmarshal(body, apiResp)
	if err != nil {
		return
	}
	apiErrors = apiResp.Errors
	if len(apiResp.Errors) > 0 {
		ee := make([]string, 0)
		for _, e := range apiResp.Errors {
			ee = append(ee, e.LongMessage)
		}
		err = fmt.Errorf("%s", strings.Join(ee, ","))
		return
	}
	item = apiResp.Item
	return
}

func buildRequest(itm string, includeSelector []string, t *ebayapi.Token) (r *http.Request, err error) {
	const (
		TOKENHEADER = `X-EBAY-API-IAF-TOKEN`
		ROUTE       = `shopping`
	)
	r, err = http.NewRequest(http.MethodGet, ebayapi.APIURL+ROUTE, nil)
	if err != nil {
		return nil, err
	}
	//set header
	r.Header.Add(TOKENHEADER, t.AccessToken)
	q := r.URL.Query()
	q.Set("callname", "GetSingleItem")
	q.Set("responseencoding", "JSON")
	q.Set("version", "1199")
	q.Set("IncludeSelector", strings.Join(includeSelector, ","))
	q.Set("ItemID", itm)
	r.URL.RawQuery = q.Encode()
	return r, nil
}
