package giphy

import (
	"net/http"
	"testing"
)

func TestTranslate(t *testing.T) {
	for _, tt := range []struct {
		args []string
		s    string
	}{
		{[]string{"foo", "bar"}, "foo bar"},
		{[]string{"bar", "baz"}, "bar baz"},
	} {
		requests := []*http.Request{}
		server, client := jsonServerAndClient(200, `{
			"data": {

			},
			"meta": {
				"status": 200,
				"msg": "OK"
			}
		}`, &requests)
		defer server.Close()

		res, err := client.Translate(tt.args)
		if err != nil {
			t.Errorf(`unexpected error %v`, err)
		}

		if got, want := res.Meta.Status, 200; got != want {
			t.Errorf(`search.Meta.Status = %+v, want %+v`, got, want)
		}

		if got := len(requests); got != 1 {
			t.Fatalf(`unexpected number of requests %d`, got)
		}

		r := requests[0]

		if got := r.URL.Path; got != "/v1/gifs/translate" {
			t.Errorf(`unexpected path %#v`, got)
		}

		params := r.URL.Query()

		if got := params.Get("api_key"); got != "dc6zaTOxFJmzC" {
			t.Errorf(`unexpected api_key %#v`, got)
		}

		if got := params.Get("rating"); got != "g" {
			t.Errorf(`unexpected rating %#v`, got)
		}

		if got := params.Get("s"); got != tt.s {
			t.Errorf(`unexpected s %#v`, got)
		}
	}
}
