package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

// Login struct
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Credential struct
type Credential struct {
	Login Login `json:"login"`
}

// Session struct
type Session struct {
	Host      string
	Username  string
	Password  string
	Token     string
	Prefix    string
	SessionID string
}
type logout struct {
	Logout struct {
	} `json:"logout"`
}

// Logout kills session.
func (session *Session) Logout() (err error) {
	payload := logout{}
	url := session.Prefix + BaseURI + "logout"
	data, err := json.Marshal(payload)
	if err != nil {
		log.Panic(err)
	}
	body := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Panic(err)
	}
	req.AddCookie(&http.Cookie{Name: "NITRO_AUTH_TOKEN", Value: session.Token, Path: "/nitro/v1"})
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
	if err != nil {
		log.Panic(err)
	}
	return
}

// New establishes new REST session
func New(host string, userName string, password string) (s *Session, err error) {
	s = &Session{
		Host:     host,
		Username: userName,
		Password: password,
		Prefix:   "https://" + host,
	}
	err = s.init()
	if err != nil {
		log.Panic(err)
	}
	return
}
func (session *Session) init() (err error) {
	response := model.Response{}
	var credential Credential
	credential.Login.Username = session.Username
	credential.Login.Password = session.Password
	route := BaseURI + "login"
	err = session.Post(route, credential, &response)
	if err != nil {
		log.Panic(err)
	}
	if response.Errorcode != 0 {
		err = errors.New(response.Message)
	}
	if err != nil {
		log.Panic(err)
	}
	return
}
func (session *Session) request(verb string, uri string, payload interface{}) (r []byte, err error) {
	var body io.Reader
	url := session.Prefix + uri
	// Disable certificate verification
	httpTransport := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:       5,
		IdleConnTimeout:    2 * time.Second,
		DisableCompression: false,
	}
	// Inspect and load payload
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			log.Panic(err)
		}
		body = bytes.NewBuffer(data)
	}
	// Create http request
	request, err := http.NewRequest(verb, url, body)
	request.Close = true
	if err != nil {
		log.Panic(err)
	}
	// Set Header objects
	if session.Token != "" {
		// Session was previously successful, so create cookies
		request.AddCookie(&http.Cookie{Name: "NITRO_AUTH_TOKEN", Value: session.Token, Path: "/nitro/v1"})
		request.Header.Set("Content-Type", "application/json")
	} else {
		// No session tokens exist, so add headers for authentication
		request.Header.Add("Content-Type", "application/vnd.com.citrix.netscaler.login+json")
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Referer", session.Prefix)
	request.Header.Set("Connection", "close")
	// Set http client
	client := &http.Client{
		Transport: httpTransport,
		Timeout:   45 * time.Second,
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 201 {
		for _, cookie := range resp.Cookies() {
			if cookie.Name == "NITRO_AUTH_TOKEN" {
				session.Token = cookie.Value
			}
			if cookie.Name == "SESSID" {
				session.SessionID = cookie.Value
			}
		}
	}
	r, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	return
}

// Returns web request with supplied interface applied
func (session *Session) restRequestInterfaceResponse(verb string, url string, payload interface{}, response interface{}) error {
	res, err := session.request(verb, url, payload)
	if err != nil || res == nil || len(res) == 0 {
		return err
	}
	return json.Unmarshal(res, &response)
}

// Get request with applicable interface
func (session *Session) Get(uri string, response interface{}) error {
	return session.restRequestInterfaceResponse("GET", uri, nil, &response)
}

// Post request with applicable interface
func (session *Session) Post(uri string, payload interface{}, response interface{}) error {
	return session.restRequestInterfaceResponse("POST", uri, payload, &response)
}

// Put request with applicable interface
func (session *Session) Put(uri string, payload interface{}, response interface{}) error {
	return session.restRequestInterfaceResponse("PUT", uri, payload, &response)
}

// Delete request with applicable interface
func (session *Session) Delete(uri string, response interface{}) error {
	return session.restRequestInterfaceResponse("DELETE", uri, nil, &response)
}
