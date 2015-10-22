package utilities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// rest client wrapper.

type RestClient struct {
	Address string
	Client  http.Client
}

func NewRestClient(address string) *RestClient {
	var c RestClient
	c.Address = address
	return &c
}

func (c *RestClient) Get(url string) (*http.Response, error) {
	resp, err := c.Client.Get(c.Address + url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// The obtained JSON is deserialised into v
func (c *RestClient) GetJSON(url string, v interface{}) error {
	resp, err := c.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		DumpResponse(resp)
		return fmt.Errorf("Invalid response from %s: %s", url, resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return err
	}
	return nil
}

// The interface v is serialised as JSON
func (c *RestClient) Post(url string, v interface{}) (*http.Response, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(v)
	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Post(c.Address+url, "application/json", &buf)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// The interface v is serialised as JSON
func (c *RestClient) PostOK(url string, v interface{}) error {
	resp, err := c.Post(url, v)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		DumpResponse(resp)
		return fmt.Errorf("Invalid response from %s: %s", url, resp.Status)
	}
	return nil
}

// The interface v is serialised as JSON
func (c *RestClient) PutOK(url string, v interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(v)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("PUT", url, &buf)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := c.Client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		DumpResponse(resp)
		return fmt.Errorf("Invalid response from %s: %s", url, resp.Status)
	}
	return nil
}

// The interface req (if not nil) is serialised as JSON content in the delete request and any response content is deserialised into the interface res (if not nil)
func (c *RestClient) DeleteOK(url string) error {
	// Construct a DELETE request
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	// Issue the DELETE request and check the response
	response, err := c.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		DumpResponse(response)
		return fmt.Errorf("Invalid response from %s: %s", url, response.Status)
	}
	return nil
}
