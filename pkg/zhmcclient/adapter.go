/*
 * =============================================================================
 * IBM Confidential
 * © Copyright IBM Corp. 2020, 2021
 *
 * The source code for this program is not published or otherwise divested of
 * its trade secrets, irrespective of what has been deposited with the
 * U.S. Copyright Office.
 * =============================================================================
 */

package zhmcclient

import (
	"encoding/json"
	"net/http"
	"path"
)

// AdapterAPI defines an interface for issuing Adapter requests to ZHMC
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o fakes/adapter.go --fake-name AdapterAPI . AdapterAPI
type AdapterAPI interface {
	ListAdapters(cpcURI string, query map[string]string) ([]Adapter, *HmcError)
	CreateHipersocket(cpcURI string, adaptor *HipersocketPayload) (string, *HmcError)
	DeleteHipersocket(adapterURI string) *HmcError
}

type AdapterManager struct {
	client ClientAPI
}

func NewAdapterManager(client ClientAPI) *AdapterManager {
	return &AdapterManager{
		client: client,
	}
}

/**
* GET /api/cpcs/{cpc-id}/adapters
* @cpcURI the ID of the CPC
* @query the fields can be queried include:
*        name,
*        adapter-id,
*        adapter-family,
*        type,
*        status
* @return adapter array
* Return: 200 and Adapters array
*     or: 400, 404, 409
 */
func (m *AdapterManager) ListAdapters(cpcURI string, query map[string]string) ([]Adapter, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, cpcURI, "/adapters")
	requestUrl = BuildUrlFromQuery(requestUrl, query)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if status == http.StatusOK {
		adapters := &AdaptersArray{}
		err := json.Unmarshal(responseBody, adapters)
		if err != nil {
			return nil, getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return adapters.ADAPTERS, nil
	}

	return nil, GenerateErrorFromResponse(responseBody)
}

/**
* POST /api/cpcs/{cpc-id}/adapters
* @cpcURI the ID of the CPC
* @adaptor the payload includes properties when create Hipersocket
* Return: 201 and body with "object-uri"
*     or: 400, 403, 404, 409, 503
 */
func (m *AdapterManager) CreateHipersocket(cpcURI string, adaptor *HipersocketPayload) (string, *HmcError) {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, cpcURI, "/adapters")

	status, responseBody, err := m.client.ExecuteRequest(http.MethodPost, requestUrl, adaptor)
	if err != nil {
		return "", err
	}

	if status == http.StatusCreated {
		uriObj := HipersocketCreateResponse{}
		err := json.Unmarshal(responseBody, &uriObj)
		if err != nil {
			return "", getHmcErrorFromErr(ERR_CODE_HMC_UNMARSHAL_FAIL, err)
		}
		return uriObj.URI, nil
	}

	return "", GenerateErrorFromResponse(responseBody)
}

/**
* DELETE /api/adapters/{adapter-id}
* @adapterURI the adapter ID to be deleted
* Return: 204
*     or: 400, 403, 404, 409, 503
 */
func (m *AdapterManager) DeleteHipersocket(adapterURI string) *HmcError {
	requestUrl := m.client.CloneEndpointURL()
	requestUrl.Path = path.Join(requestUrl.Path, adapterURI)

	status, responseBody, err := m.client.ExecuteRequest(http.MethodDelete, requestUrl, nil)
	if err != nil {
		return err
	}

	if status == http.StatusNoContent {
		return nil
	}

	return GenerateErrorFromResponse(responseBody)
}
