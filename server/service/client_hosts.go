package service

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// GetHosts retrieves the list of all Hosts
func (c *Client) GetHosts() ([]HostResponse, error) {
	response, err := c.AuthenticatedDo("GET", "/api/v1/kolide/hosts", nil)
	if err != nil {
		return nil, errors.Wrap(err, "GET /api/v1/kolide/hosts")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.Errorf(
			"get hosts received status %d %s",
			response.StatusCode,
			extractServerErrorText(response.Body),
		)
	}
	var responseBody listHostsResponse
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return nil, errors.Wrap(err, "decode list hosts response")
	}
	if responseBody.Err != nil {
		return nil, errors.Errorf("list hosts: %s", responseBody.Err)
	}

	return responseBody.Hosts, nil
}

// HostByIdentifier retrieves a host by the uuid, osquery_host_id, hostname, or
// node_key.
func (c *Client) HostByIdentifier(identifier string) (*HostResponse, error) {
	response, err := c.AuthenticatedDo("GET", "/api/v1/kolide/hosts/identifier/"+identifier, nil)
	if err != nil {
		return nil, errors.Wrap(err, "GET /api/v1/kolide/hosts/identifier")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.Errorf(
			"get host by identifier received status %d %s",
			response.StatusCode,
			extractServerErrorText(response.Body),
		)
	}
	var responseBody getHostResponse
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return nil, errors.Wrap(err, "decode host response")
	}
	if responseBody.Err != nil {
		return nil, errors.Errorf("get host by identifier: %s", responseBody.Err)
	}

	return responseBody.Host, nil
}
