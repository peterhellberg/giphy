package giphy

import (
	"encoding/json"
	"errors"
	"strings"
)

func (c *Client) Translate(args []string) (Translate, error) {
	argsStr := strings.Join(args, " ")

	req, err := c.NewRequest("/translate?s=" + argsStr)
	if err != nil {
		return Translate{}, err
	}

	var translate Translate
	if _, err = c.Do(req, &translate); err != nil {
		return Translate{}, err
	}

	// Check if the first character in Data is a [
	if translate.RawData[0] == '[' {
		return Translate{}, errors.New("no image found")
	}

	err = json.Unmarshal(translate.RawData, &translate.Data)
	if err != nil {
		return Translate{}, errors.New("could not unmarshal JSON data")
	}

	return translate, nil
}
