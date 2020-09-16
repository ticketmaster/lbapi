package client

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Login struct.
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Host struct {
	Name     string
	UserName string
	Password string
}

// Client struct.
type Client struct {
	Host   Host
	Prefix string
	Cookie *http.Cookie
}

// logOut struct.
type logOut struct {
	Logout struct {
	} `json:"logout"`
}

// LogOut kills session.
func (c *Client) LogOut() (err error) {
	payload := logOut{}
	var payloadIO io.Reader
	verb := "POST"
	url := c.Prefix + BaseURI + "logout"
	jsonStr, err := json.Marshal(payload)
	if err != nil {
		return
	}
	payloadIO = bytes.NewBuffer(jsonStr)
	req, err := http.NewRequest(verb, url, payloadIO)
	if err != nil {
		log.Printf("Error creating HTTP request: %v\n", err)
		return
	}
	req.AddCookie(c.Cookie)
	req.Header.Set("Content-Type", "application/json")
	httpTransport := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:       5,
		IdleConnTimeout:    2 * time.Second,
		DisableCompression: false,
	}
	client := &http.Client{
		Transport: httpTransport,
		Timeout:   45 * time.Second,
	}
	_, err = client.Do(req)
	c = nil
	return
}

// NewClient establishes new REST session.
func NewClient(target string, userName string, password string) (r *Client, err error) {
	host := Host{
		Name:     target,
		UserName: userName,
		Password: password}
	r = &Client{Host: host,
		Prefix: "https://" + target,
	}
	// Attempt to initiate session. Return error if unsuccessful.
	err = r.create()
	return
}

// create establishes a simple connection to the schema endpoint purely to validate connectivity and generate a cookie.
func (c *Client) create() (err error) {
	resp := new(interface{})
	route := "?_schema&_return_as_object=1"
	err = c.Get(route, &resp)
	return
}

// basicAuth formats credentials into a base64 string for authentication.
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// restRequest is a helper for creating HTTP requests.
func (c *Client) restRequest(verb string, uri string, payload interface{}) ([]byte, error) {
	var body []byte
	var payloadIO io.Reader
	url := c.Prefix + BaseURI + uri
	//log.Println(url)
	// Disable certificate verification
	httpTransport := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:       5,
		IdleConnTimeout:    2 * time.Second,
		DisableCompression: false,
	}
	// Inspect and load payload.
	if payload != nil {
		jsonStr, err := json.Marshal(payload)
		if err != nil {
			return body, err
		}
		payloadIO = bytes.NewBuffer(jsonStr)
	}
	// Create http request.
	req, err := http.NewRequest(verb, url, payloadIO)
	if err != nil {
		log.Printf("Error creating HTTP request: %v\n", err)
		return body, err
	}
	// Set Header objects.
	if c.Cookie != nil {
		// Session was previously successful, so create cookies
		req.AddCookie(c.Cookie)
		req.Header.Set("Content-Type", "application/json")
	} else {
		// No session tokens exist, so add headers for authentication
		req.Header.Add("Authorization", "Basic "+basicAuth(c.Host.UserName, c.Host.Password))
		c.Host.UserName = ""
		c.Host.Password = ""
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", c.Prefix)
	// Set http client
	client := &http.Client{
		Transport: httpTransport,
		Timeout:   45 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error creating http client: %+v\n", err)
		return body, err
	}
	if resp.StatusCode == 200 {
		for _, cookie := range resp.Cookies() {
			if cookie.Name == "ibapauth" {
				c.Cookie = cookie
			}
		}
	}
	req.Close = true
	req.Header.Set("Connection", "close")
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Response body read failed: %v\n", err)
		return body, err
	}
	return body, nil
}

// Returns web request with supplied interface applied.
func (c *Client) restRequestInterfaceResponse(verb string, url string, payload interface{}, response interface{}) error {
	res, err := c.restRequest(verb, url, payload)
	if err != nil || res == nil || len(res) == 0 {
		return err
	}
	return json.Unmarshal(res, &response)
}

// Get request with applicable interface.
func (c *Client) Get(uri string, response interface{}) error {
	return c.restRequestInterfaceResponse("GET", uri, nil, response)
}

// Post request with applicable interface.
func (c *Client) Post(uri string, payload interface{}, response interface{}) error {
	return c.restRequestInterfaceResponse("POST", uri, payload, response)
}

// Put request with applicable interface.
func (c *Client) Put(uri string, payload interface{}, response interface{}) error {
	return c.restRequestInterfaceResponse("PUT", uri, payload, response)
}

// Delete request with applicable interface.
func (c *Client) Delete(uri string, response interface{}) error {
	return c.restRequestInterfaceResponse("DELETE", uri, nil, response)
}
