package giphy

import "errors"

func (c *Client) Trending(args ...[]string) (Trending, error) {
	req, err := c.NewRequest("/trending")
	if err != nil {
		return Trending{}, err
	}

	var res Trending
	if _, err = c.Do(req, &res); err != nil {
		return res, err
	}

	if len(res.Data) == 0 {
		return res, errors.New("no trending images found")
	}

	return res, nil
}
