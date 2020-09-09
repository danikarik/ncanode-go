package ncanode

type TSPSignRequest struct {
	Raw              string        `json:"raw"`
	UseTsaPolicy     Policy        `json:"useTsaPolicy,omitempty"`
	TSPHashAlgorithm HashAlgorithm `json:"tspHashAlgorithm,omitempty"`
}

type TSPSignResponse struct {
	apiResponse
	Result struct {
		TSP string `json:"tsp"`
	} `json:"result"`
}

func (c *Client) TSPSign(raw string, policy Policy, alg HashAlgorithm) (*TSPSignResponse, error) {
	if raw == "" {
		return nil, ErrInvalidRequestBody
	}

	body := apiRequest{
		Version: _v1,
		Method:  "TSP.sign",
		Params: TSPSignRequest{
			Raw:              raw,
			UseTsaPolicy:     policy,
			TSPHashAlgorithm: alg,
		},
	}

	var reply TSPSignResponse
	if err := c.call(body, &reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
