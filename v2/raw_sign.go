package ncanode

type rawSignRequest struct {
	P12              string        `json:"p12"`
	Password         string        `json:"password"`
	Raw              string        `json:"raw"`
	CreateTsp        bool          `json:"createTsp,omitempty"`
	UseTsaPolicy     Policy        `json:"useTsaPolicy,omitempty"`
	TSPHashAlgorithm HashAlgorithm `json:"tspHashAlgorithm,omitempty"`
	TspInCms         bool          `json:"tspInCms,omitempty"`
}

// RawSignResponse describes json response from RawSign.
type RawSignResponse struct {
	apiResponse
	Result struct {
		CMS string `json:"cms"`
		TSP string `json:"tsp,omitempty"`
	} `json:"result"`
}

// RawSign signs any input string and saves into cms.
//
// See https://ncanode.kz/docs.php?go=b52dfc5eddafafb5d7c8cccb06c5f0e011a27f3d
func (c *Client) RawSign(p12, password, raw string, config *TSPConfig) (*RawSignResponse, error) {
	if p12 == "" || password == "" || raw == "" {
		return nil, ErrInvalidRequestBody
	}

	if config == nil {
		config = &TSPConfig{}
	}

	body := apiRequest{
		Version: c.version,
		Method:  "RAW.sign",
		Params: rawSignRequest{
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
