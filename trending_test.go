package giphy

import (
	"net/http"
	"testing"
)

func TestTrending(t *testing.T) {
	for _, tt := range []struct {
		args []string
		s    string
	}{
		{[]string{"foo", "bar"}, "foo bar"},
		{[]string{"bar", "baz"}, "bar baz"},
	} {
		requests := []*http.Request{}
		server, client := jsonServerAndClient(200, `{
		  "data": [
		    {
		      "type": "gif",
		      "id": "eMu0803X2zkWY",
		      "url": "https://giphy.com/gifs/the-simpsons-eMu0803X2zkWY",
		      "bitly_gif_url": "http://gph.is/1H5kauM",
		      "bitly_url": "http://gph.is/1H5kauM",
		      "embed_url": "https://giphy.com/embed/eMu0803X2zkWY",
		      "username": "",
		      "source": "http://bluefastakan.tumblr.com/post/66074932372/every-day-on-the-internet",
		      "rating": "g",
		      "caption": "",
		      "content_url": "",
		      "import_datetime": "2013-11-05 06:56:00",
		      "trending_datetime": "2015-05-14 15:06:52",
		      "images": {
		        "original": {
		          "url": "https://media1.giphy.com/media/eMu0803X2zkWY/giphy.gif",
		          "width": "350",
		          "height": "268",
		          "size": "1313764",
		          "frames": "25",
		          "mp4": "https://media1.giphy.com/media/eMu0803X2zkWY/giphy.mp4",
		          "mp4_size": "91690",
		          "webp": "https://media1.giphy.com/media/eMu0803X2zkWY/giphy.webp",
		          "webp_size": "427080"
		        }
		      }
		    }
		  ],
		  "meta": {
		    "status": 200,
		    "msg": "OK"
		  },
		  "pagination": {
		    "count": 1,
		    "offset": 5
		  }
		}`, &requests)
		defer server.Close()

		res, err := client.Trending(tt.args)
		if err != nil {
			t.Errorf(`unexpected error %v`, err)
		}

		if got, want := res.Meta.Status, 200; got != want {
			t.Errorf(`res.Meta.Status = %+v, want %+v`, got, want)
		}

		if got, want := res.Pagination.Count, 1; got != want {
			t.Errorf(`res.Pagination.Count = %+v, want %+v`, got, want)
		}

		if got := len(requests); got != 1 {
			t.Fatalf(`unexpected number of requests %d`, got)
		}

		r := requests[0]

		if got := r.URL.Path; got != "/v1/gifs/trending" {
			t.Errorf(`unexpected path %#v`, got)
		}

		params := r.URL.Query()

		if got := params.Get("api_key"); got != "dc6zaTOxFJmzC" {
			t.Errorf(`unexpected api_key %#v`, got)
		}

		if got := params.Get("rating"); got != "g" {
			t.Errorf(`unexpected rating %#v`, got)
		}
	}
}
