package giphy

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	expectedAPIKey        = "dc6zaTOxFJmzC"
	expectedLimit         = 10
	expectedRating        = "g"
	expectedBaseURLString = "https://api.giphy.com"
	expectedBasePath      = "/v1"
	expectedUserAgent     = "giphy.go"
)

func TestDefaultClient(t *testing.T) {
	if got, want := DefaultClient.APIKey, expectedAPIKey; got != want {
		t.Errorf("DefaultClient.APIKey = %s, want %s", got, want)
	}

	if got, want := DefaultClient.Limit, expectedLimit; got != want {
		t.Errorf("DefaultClient.Limit = %d, want %d", got, want)
	}

	if got, want := DefaultClient.Rating, expectedRating; got != want {
		t.Errorf("DefaultClient.Rating = %s, want %s", got, want)
	}

	if got, want := DefaultClient.BaseURL.String(), expectedBaseURLString; got != want {
		t.Errorf("DefaultClient.BaseURL.String() = %s, want %s", got, want)
	}

	if got, want := DefaultClient.BasePath, expectedBasePath; got != want {
		t.Errorf("DefaultClient.BasePath = %s, want %s", got, want)
	}

	if got, want := DefaultClient.UserAgent, expectedUserAgent; got != want {
		t.Errorf("DefaultClient.UserAgent = %s, want %s", got, want)
	}
}

func TestNewRequest(t *testing.T) {
	for _, tt := range []struct {
		s string
		u string
	}{
		{"/foo", "https://api.giphy.com/v1/foo?api_key=dc6zaTOxFJmzC&rating=g"},
		{"/bar", "https://api.giphy.com/v1/bar?api_key=dc6zaTOxFJmzC&rating=g"},
	} {
		req, err := DefaultClient.NewRequest(tt.s)
		if err != nil {
			t.Errorf(`unexpected error %v`, err)
		}

		if got, want := req.Method, "GET"; got != want {
			t.Errorf(`req.Method = %v, want %v`, got, want)
		}

		if got, want := req.Header.Get("User-Agent"), expectedUserAgent; got != want {
			t.Errorf(`unexpected User-Agent: %s, want %s`, got, want)
		}

		if got, want := req.URL.String(), tt.u; got != want {
			t.Errorf(`req.URL.String() = %v, want %v`, got, want)
		}
	}
}

func TestDo(t *testing.T) {
	server, client := jsonServerAndClient(200, `{"foo": 123}`)
	defer server.Close()

	req, err := client.NewRequest("/")
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	v := map[string]int{}

	client.Do(req, &v)

	if v["foo"] != 123 {
		t.Errorf(`unexpected response %+v`, v)
	}
}

func jsonServerAndClient(code int, body string, requests ...*[]*http.Request) (*httptest.Server, *Client) {
	return testServerAndClient(func(w http.ResponseWriter, r *http.Request) {
		if len(requests) > 0 {
			*requests[0] = append(*requests[0], r)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write([]byte(body))
	})
}

func testServerAndClient(f func(http.ResponseWriter, *http.Request)) (*httptest.Server, *Client) {
	server := httptest.NewServer(http.HandlerFunc(f))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		}}

	client := NewClient(&http.Client{Transport: transport})
	client.BaseURL, _ = url.Parse(server.URL)

	return server, client
}
