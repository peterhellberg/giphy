package giphy

import (
	"net/http"
	"testing"
)

func TestSearch(t *testing.T) {
	for _, tt := range []struct {
		q string
	}{
		{"foo"},
		{"bar"},
	} {
		requests := []*http.Request{}
		server, client := jsonServerAndClient(200, `{
			"meta": {
				"status": 200,
				"msg": "OK"
			}
		}`, &requests)
		defer server.Close()

		search, err := client.Search([]string{tt.q})
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

		params := r.URL.Query()

		if got := params.Get("api_key"); got != "dc6zaTOxFJmzC" {
			t.Errorf(`unexpected api_key %#v`, got)
		}

		if got := params.Get("limit"); got != "10" {
			t.Errorf(`unexpected limit %#v`, got)
		}

		if got := params.Get("rating"); got != "g" {
			t.Errorf(`unexpected rating %#v`, got)
		}

		if got := params.Get("q"); got != tt.q {
			t.Errorf(`unexpected q %#v`, got)
		}
	}
}
