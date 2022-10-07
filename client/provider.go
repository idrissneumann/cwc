package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func GetProviderRegions() (*ProviderRegions, error) {

	endpoint := GetDefaultEndpoint()
	if endpoint == "" {
		return nil, fmt.Errorf("endpoint is not set")
	}
	provider := GetDefaultProvider()
	if provider == "" {
		return nil, fmt.Errorf("provider is not set")
	}
	c := NewClient()
	body, err := c.httpRequest(fmt.Sprintf("/provider/%s/region", provider), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	providerRegions := &ProviderRegions{}
	err = json.NewDecoder(body).Decode(providerRegions)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	println(providerRegions)
	return providerRegions, nil
}

func GetProviders() (*AvailableProviders, error) {
	endpoint := GetDefaultEndpoint()
	if endpoint == "" {
		return nil, fmt.Errorf("endpoint is not set")
	}
	c := NewClient()
	body, err := c.httpRequest(fmt.Sprintf("/provider"), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	available_providers := &AvailableProviders{}
	err = json.NewDecoder(body).Decode(available_providers)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return available_providers, nil
}
