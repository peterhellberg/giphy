package giphy

import "testing"

func TestGIF(t *testing.T) {
	server, client := testServerAndClient(200, []byte(`{
		"data": {
			"type": "gif",
			"id": "feqkVgjJpYtjy",
			"url": "http://giphy.com/gifs/feqkVgjJpYtjy"
		},
		"meta": {
			"status": 200,
			"msg": "OK"
		}
	}`))
	defer server.Close()

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
