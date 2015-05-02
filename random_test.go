package giphy

import (
	"net/http"
	"testing"
)

func TestRandom(t *testing.T) {
	requests := []*http.Request{}
	server, client := jsonServerAndClient(200, `{
		"data": {
			"type": "gif",
			"id": "feqkVgjJpYtjy",
			"url": "http://giphy.com/gifs/feqkVgjJpYtjy"
		},
		"meta": {
			"status": 200,
			"msg": "OK"
		}
	}`, &requests)
	defer server.Close()

	random, err := client.Random([]string{"bar", "baz"})
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	if got, want := random.Meta.Status, 200; got != want {
		t.Errorf(`gif.Meta.Status = %+v, want %+v`, got, want)
	}

	if got, want := random.Data.Type, "gif"; got != want {
		t.Errorf(`gif.Data.Type = %+v, want %+v`, got, want)
	}

	if got := len(requests); got != 1 {
		t.Fatalf(`unexpected number of requests %d`, got)
	}

	r := requests[0]

	if got := r.URL.Path; got != "/v1/gifs/random" {
		t.Errorf(`unexpected path %#v`, got)
	}

	params := r.URL.Query()

	if got := params.Get("api_key"); got != "dc6zaTOxFJmzC" {
		t.Errorf(`unexpected api_key %#v`, got)
	}

	if got := params.Get("rating"); got != "g" {
		t.Errorf(`unexpected rating %#v`, got)
	}

	if got := params.Get("tag"); got != "bar baz" {
		t.Errorf(`unexpected tag %#v`, got)
	}
}

func TestRandomRequest(t *testing.T) {
	var (
		expectedPath     = "/v1/gifs/random"
		expectedRawQuery = "api_key=dc6zaTOxFJmzC&rating=g&tag=bar+baz"
	)

	reqs := []http.Request{}

	server, client := testServerAndClient(
		func(w http.ResponseWriter, r *http.Request) {
			reqs = append(reqs, *r)
		},
	)
	defer server.Close()

	client.Random([]string{"bar", "baz"})

	if len(reqs) != 1 {
		t.Errorf(`unexpected number of requests`)
	}

	if got := reqs[0].URL.Path; got != expectedPath {
		t.Errorf(`reqs[0].URL.Path = %#v, want %#v`, got, expectedPath)
	}

	if got := reqs[0].URL.RawQuery; got != expectedRawQuery {
		t.Errorf(`reqs[0].URL.RawQuery = %+v, want %+v`, got, expectedRawQuery)
	}
}

func TestRandom_noImageFound(t *testing.T) {
	server, client := jsonServerAndClient(200, `{"data":[]}`)
	defer server.Close()

	if _, err := client.Random([]string{"bar"}); err != ErrNoImageFound {
		t.Errorf(`expected ErrNoImageFound, got %v`, err)
	}
}
