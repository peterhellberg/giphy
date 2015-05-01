package giphy

import (
	"net/http"
	"testing"
)

func TestSearch(t *testing.T) {
	requests := []*http.Request{}
	server, client := jsonServerAndClient(200, `{
		"meta": {
			"status": 200,
			"msg": "OK"
		}
	}`, &requests)
	defer server.Close()

	search, err := client.Search([]string{"foo"})
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	if got, want := search.Meta.Status, 200; got != want {
		t.Errorf(`search.Meta.Status = %+v, want %+v`, got, want)
	}

	if got := len(requests); got != 1 {
		t.Fatalf(`unexpected number of requests %d`, got)
	}

	r := requests[0]

	if got := r.URL.Path; got != "/v1/gifs/search" {
		t.Errorf(`unexpected path %#v`, got)
	}

	if got := r.URL.RawQuery; got != "api_key=dc6zaTOxFJmzC&limit=10&q=foo&rating=g" {
		t.Errorf(`unexpected query %#v`, got)
	}
}
