package giphy

import "testing"

func TestSearch(t *testing.T) {
	server, client := jsonServerAndClient(200, `{
		"meta": {
			"status": 200,
			"msg": "OK"
		}
	}`)
	defer server.Close()

	search, err := client.Search([]string{"foo"})
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	if got, want := search.Meta.Status, 200; got != want {
		t.Errorf(`search.Meta.Status = %+v, want %+v`, got, want)
	}
}
