package ncanode

type tspSignRequest struct {
	Raw              string        `json:"raw"`
	UseTsaPolicy     Policy        `json:"useTsaPolicy,omitempty"`
	TSPHashAlgorithm HashAlgorithm `json:"tspHashAlgorithm,omitempty"`
}

// TSPSignResponse describes json response from TSPSign.
type TSPSignResponse struct {
	apiResponse
	Result struct {
		TSP string `json:"tsp"`
	} `json:"result"`
}

// TSPSign signs any input string using TSP.
//
// See https://ncanode.kz/docs.php?go=366ea24993a9887051c5f647f8dba8fa5e236d58
func (c *Client) TSPSign(raw string, policy Policy, alg HashAlgorithm) (*TSPSignResponse, error) {
	if raw == "" {
		return nil, ErrInvalidRequestBody
	}

	body := apiRequest{
		Version: c.version,
		Method:  "TSP.sign",
		Params: tspSignRequest{
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
