package giphy

import "testing"

func TestGIF(t *testing.T) {
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

	if _, err := client.GIF("invalid?/id"); err == nil {
		t.Errorf(`expected error for GIF id: invalid?/id`)
	}

	gif, err := client.GIF("foo")
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	if got, want := gif.Meta.Status, 200; got != want {
		t.Errorf(`gif.Meta.Status = %+v, want %+v`, got, want)
	}

	if got, want := gif.Data.Type, "gif"; got != want {
		t.Errorf(`gif.Data.Type = %+v, want %+v`, got, want)
	}
}

func TestGIF_noImageFound(t *testing.T) {
	server, client := jsonServerAndClient(200, `{"data":[]}`)
	defer server.Close()

	if _, err := client.GIF("foo"); err != ErrNoImageFound {
		t.Errorf(`expected ErrNoImageFound, got %v`, err)
	}
}
