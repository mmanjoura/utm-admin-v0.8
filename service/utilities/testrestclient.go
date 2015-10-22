package utilities

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

// rest client wrapper.

type TestRestClient struct {
	RestClient
}

func NewTestRestClient(address string, client http.Client) *TestRestClient {
	var c TestRestClient
	c.Address = address
	c.Client = client
	return &c
}

func (c *TestRestClient) GetJSONunchecked(url string, v interface{}) (*http.Response, error) {
	resp, err := c.Get(url)
	if err != nil {
		return resp, err
	}
	if (resp.StatusCode >= 200) && (resp.StatusCode <= 299) {
		// Only attempt to decode the JSON if the call was successful
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}

func (c *TestRestClient) GetStringBody(url string) (string, error) {
	response, err := c.Get(url)
	if err != nil {
		return "", err
	}
	if response.StatusCode != http.StatusOK {
		DumpResponse(response)
		return "", fmt.Errorf("invalid response from %s: %s", url, response.Status)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	str := buf.String()
	return str, nil
}

func (c *TestRestClient) PostOmitField(url string, v interface{}, fieldToOmit string) (*http.Response, error) {
	// Encode the object into JSON in a bytes object
	var buf1 bytes.Buffer
	err := json.NewEncoder(&buf1).Encode(v)
	if err != nil {
		return nil, err
	}

	// Convert the bytes object to a string, remove the field and convert it back
	inTxt := buf1.String()
	i := strings.Index(inTxt, fieldToOmit)
	if i < 2 {
		return nil, errors.New("Cannot find field '" + fieldToOmit + "' in JSON '" + inTxt + "'")
	}
	leftTxt := inTxt[:i-1]
	rightTxt := inTxt[i+len(fieldToOmit):]
	i = strings.Index(rightTxt, ",")
	j := strings.Index(rightTxt, "}")
	if (i >= 0) && ((j < 0) || (i < j)) {
		// There is a trailing comma and (there is no "}" or it is before the "}") so eliminate it
		rightTxt = rightTxt[i+1:]
	} else if j >= 0 {
		// There is trailing "}" so keep it and remove any preceeding comma
		rightTxt = rightTxt[j:]
		// Reverse the leftTxt and search it to remove a preceeding comma
		revTxt := reverse(leftTxt)
		i = strings.Index(revTxt, ",")
		j = strings.Index(revTxt, "{")
		if (i >= 0) && ((j < 0) || (i < j)) {
			// There is a preceeding comma and (there is no "{" or it is before the "{") so eliminate it
			leftTxt = reverse(revTxt[i+1:])
		} else if j >= 0 {
			// There is a preceeding "{" so keep it
			leftTxt = reverse(revTxt[j:])
		} else {
			return nil, errors.New("Cannot find either a ',' or '{' before field '" + fieldToOmit + "' in JSON '" + inTxt + "'")
		}
	} else {
		return nil, errors.New("Cannot find either a ',' or '}' after field '" + fieldToOmit + "' in JSON '" + inTxt + "'")
	}
	txt := leftTxt + rightTxt
	buf2 := bytes.NewBufferString(txt)

	// POST the request
	resp, err := c.Client.Post(c.Address+url, "application/json", buf2)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *TestRestClient) PostEmptyBody(url string) (*http.Response, error) {
	resp, err := c.Post(url, "")
	return resp, err
}

func (c *TestRestClient) PostEmptyBodyOK(url string) (*http.Response, error) {
	resp, err := c.PostEmptyBody(url)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode != http.StatusOK {
		DumpResponse(resp)
		return resp, fmt.Errorf("invalid response from %s: %s", url, resp.Status)
	}
	return resp, nil
}

func (c *TestRestClient) Delete(url string) (*http.Response, error) {
	// Construct a DELETE request
	request, err := http.NewRequest("DELETE", c.Address+url, nil)
	if err != nil {
		return nil, err
	}

	// Issue the DELETE request and check the response
	response, err := c.Client.Do(request)
	if err != nil {
		return response, err
	}
	if response.StatusCode != http.StatusOK {
		DumpResponse(response)
		return response, fmt.Errorf("Invalid response from %s: %s", url, response.Status)
	}

	return response, err
}
