package discovery

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/h3mmy/ddns/ddns/internal/models"
)

// utility method for parsing an http response into a target iface
func parseJsonResponse(r *http.Response, target interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

// utility method for parsing an http response into a raw string
func parseStringResponse(r *http.Response) (*string, error) {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	s := string(b)
	return &s, nil
}

// Get HttpClient with specific TCP Version
func getHttpClientWithVersion(ver models.TCPVersion) *http.Client {
	zeroDialer := &net.Dialer{}
	httpClient := http.DefaultClient
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return zeroDialer.DialContext(ctx, string(ver), addr)
	}
	httpClient.Transport = tr
	return httpClient
}
