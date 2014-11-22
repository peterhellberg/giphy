package giphy

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func (c *Client) GIF(id string) (GIF, error) {
	if strings.ContainsAny(id, "/&?") {
		return GIF{}, fmt.Errorf("Invalid giphy id: `%v`", id)
	}

	req, err := c.NewRequest("/" + id)
	if err != nil {
		return GIF{}, err
	}

	var gif GIF
	if _, err = c.Do(req, &gif); err != nil {
		return GIF{}, err
	}

	if gif.RawData == nil || gif.RawData[0] == '[' {
		return GIF{}, errors.New("no image found")
	}

	// Check if the first character in Data is a {
	if gif.RawData[0] == '{' {
		var d Data

		err = json.Unmarshal(gif.RawData, &d)
		if err != nil {
			return GIF{}, errors.New("could not unmarshal JSON data")
		}

		gif.Data = d

		return gif, nil
	}

	return GIF{}, errors.New("unknown error")
}
