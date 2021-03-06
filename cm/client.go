package cm

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// CreateGetRequest creates an CM Server GET request
func (c CMServer) CreateGetRequest(urlSuffix string) *http.Request {
	uri := c.GetCMUri(urlSuffix)
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Upgrade-Insecure-Reqests", "1")
	request.SetBasicAuth(c.Username, c.Password)
	return request
}

// CreatePostRequest creates an CM Server POST request with body
func (c CMServer) CreatePostRequest(body bytes.Buffer, urlSuffix string) *http.Request {
	uri := c.GetCMUri(urlSuffix)
	request, err := http.NewRequest("POST", uri, &body)
	if err != nil {
		panic(err)
	}
	request.SetBasicAuth(c.Username, c.Password)
	return request
}

// CreatePutRequest creates an CM Server PUT request with body
func (c CMServer) CreatePutRequest(body bytes.Buffer, urlSuffix string) *http.Request {
	uri := c.GetCMUri(urlSuffix)
	request, err := http.NewRequest("PUT", uri, &body)
	if err != nil {
		panic(err)
	}
	request.SetBasicAuth(c.Username, c.Password)
	return request
}

// GetCMUri creates the CM uri with /api/vx/ suffix (+ /api/vx/clusters/<cluster> suffix is useCluster is enabled)
func (c CMServer) GetCMUri(uriSuffix string) string {
	return fmt.Sprintf("%s://%s:%v/api/v%v/%s", c.Protocol, c.Hostname, c.Port, c.APIVersion, uriSuffix)
}

// GetHttpClient create HTTP client instance for CM Server
func GetHttpClient() *http.Client {
	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			MaxIdleConns:          100,
			IdleConnTimeout:       30 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		},
	}
	return httpClient
}

// ProcessCMItems get "items" from CM server response
func ProcessCMItems(request *http.Request) Items {
	bodyBytes := ProcessRequest(request)
	return ProcessCMItemsFromBytes(bodyBytes)
}

// ProcessCMItemsFromSSHResponse get "items" from CM gateway SSH response
func ProcessCMItemsFromSSHResponse(sshResponse RemoteResponse) Items {
	return ProcessCMItemsFromBytes([]byte(sshResponse.StdOut))
}

// ProcessCMItemsFromBytes get "items" from bytes
func ProcessCMItemsFromBytes(bodyBytes []byte) Items {
	var cmItems Items
	err := json.Unmarshal(bodyBytes, &cmItems)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return cmItems
}

// ProcessAsMap get map format response
func ProcessAsMap(request *http.Request) map[string]interface{} {
	return ProcessAsMapFromBytes(ProcessRequest(request))
}

// ProcessAsMapFromSSHResponse get JSON map from CM gateway SSH response
func ProcessAsMapFromSSHResponse(sshResponse RemoteResponse) map[string]interface{} {
	return ProcessAsMapFromBytes([]byte(sshResponse.StdOut))
}

// ProcessAsMapFromBytes get JSON map from bytes
func ProcessAsMapFromBytes(bodyBytes []byte) map[string]interface{} {
	var responseMap map[string]interface{}
	err := json.Unmarshal(bodyBytes, &responseMap)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return responseMap
}

// ProcessRequest get a simple response from a REST call
func ProcessRequest(request *http.Request) []byte {
	client := GetHttpClient()
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		errorResponseMessage := fmt.Sprintf("Response status code: %v", response.StatusCode)
		fmt.Println(errorResponseMessage)
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(bodyBytes))
		os.Exit(1)
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return bodyBytes
}

// CreateGatewayCurlGetCommand generate a curl GET command for gateway usage
func (c CMServer) CreateGatewayCurlGetCommand(uri string) string {
	return createGatewayCurlCommand(c, uri, "GET", "")
}

// CreateGatewayCurlPostCommand generate a curl POST command for gateway usage
func (c CMServer) CreateGatewayCurlPostCommand(uri string, body string) string {
	return createGatewayCurlCommand(c, uri, "POST", body)
}

// CreateGatewayCurlPutCommand generate a curl PUT command for gateway usage
func (c CMServer) CreateGatewayCurlPutCommand(uri string, body string) string {
	return createGatewayCurlCommand(c, uri, "PUT", body)
}

func createGatewayCurlCommand(c CMServer, uri string, method string, body string) string {
	bodyCommand := ""
	if len(body) > 0 {
		bodyCommand = fmt.Sprintf(" -H \"Content-Type: application/json\" --data '%v'", body)
	}
	return fmt.Sprintf("curl -s -X %v -k -u %v:%v%v 'http://localhost:7180/api/v%v/%v'", method, c.Username, c.Password, bodyCommand, c.APIVersion, uri)
}
