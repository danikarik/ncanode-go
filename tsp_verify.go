package ncanode

import "bytes"

type TSPVerifyRequest struct {
	CMS string `json:"cms"`
}

type TSPVerifyResponse struct {
	APIResponse
	Result struct {
		TSPHashAlgorithm HashAlgorithm `json:"tspHashAlgorithm"`
		SerialNumber     string        `json:"serialNumber"`
		GenTime          Time          `json:"genTime"`
		Hash             string        `json:"hash"`
		Policy           string        `json:"policy"`
	} `json:"result"`
}

func (c *Client) TSPVerify(cms string) (*TSPVerifyResponse, error) {
	if len(cms) == 0 {
		return nil, ErrInvalidRequestBody
	}

	body := APIRequest{
		Version: _v1,
		Method:  "TSP.verify",
		Params:  TSPVerifyRequest{CMS: cms},
	}

	mod := func(in []byte) ([]byte, error) {
		return bytes.Replace(in, []byte(`\\`), []byte(`\`), -1), nil
	}

	var reply TSPVerifyResponse
	if err := c.call(body, &reply, mod); err != nil {
		return nil, err
	}

	return &reply, nil
}
