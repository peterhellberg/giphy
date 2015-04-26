package giphy

import "testing"

func TestRandom(t *testing.T) {
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
	}`)
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
}

func TestRandom_noImageFound(t *testing.T) {
	server, client := jsonServerAndClient(200, `{"data":[]}`)
	defer server.Close()

	if _, err := client.Random([]string{"bar"}); err != ErrNoImageFound {
		t.Errorf(`expected ErrNoImageFound, got %v`, err)
	}
}
