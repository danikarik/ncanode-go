package ncanode

type RawSignRequest struct {
	P12              string        `json:"p12"`
	Password         string        `json:"password"`
	Raw              string        `json:"raw"`
	CreateTsp        bool          `json:"createTsp,omitempty"`
	UseTsaPolicy     Policy        `json:"useTsaPolicy,omitempty"`
	TSPHashAlgorithm HashAlgorithm `json:"tspHashAlgorithm,omitempty"`
	TspInCms         bool          `json:"tspInCms,omitempty"`
}

type RawSignResponse struct {
	apiResponse
	Result struct {
		CMS string `json:"cms"`
		TSP string `json:"tsp,omitempty"`
	} `json:"result"`
}

func (c *Client) RawSign(p12, password, raw string, config *TSPConfig) (*RawSignResponse, error) {
	if p12 == "" || password == "" || raw == "" {
		return nil, ErrInvalidRequestBody
	}

	if config == nil {
		config = &TSPConfig{}
	}

	body := apiRequest{
		Version: _v1,
		Method:  "RAW.sign",
		Params: RawSignRequest{
			P12:              p12,
			Password:         password,
			Raw:              raw,
			CreateTsp:        config.Enabled,
			UseTsaPolicy:     config.Policy,
			TSPHashAlgorithm: config.HashAlgorithm,
			TspInCms:         config.InCMS,
		},
	}

	var reply RawSignResponse
	if err := c.call(body, &reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
