package ncanode

type Policy string

const (
	TSAGostPolicy   Policy = "TSA_GOST_POLICY"
	TSAGostGtPolicy Policy = "TSA_GOSTGT_POLICY"
)

type HashAlgorithm string

const (
	MD5         HashAlgorithm = "MD5"
	SHA1        HashAlgorithm = "SHA1"
	SHA224      HashAlgorithm = "SHA224"
	SHA256      HashAlgorithm = "SHA256"
	SHA384      HashAlgorithm = "SHA384"
	SHA512      HashAlgorithm = "SHA512"
	RIPEMD128   HashAlgorithm = "RIPEMD128"
	RIPEMD160   HashAlgorithm = "RIPEMD160"
	RIPEMD256   HashAlgorithm = "RIPEMD256"
	GOST34311GT HashAlgorithm = "GOST34311GT"
	GOST34311   HashAlgorithm = "GOST34311"
)

type TSPConfig struct {
	Enabled       bool
	Policy        Policy
	HashAlgorithm HashAlgorithm
	InCMS         bool
}

type xmlSignRequest struct {
	P12              string        `json:"p12"`
	Password         string        `json:"password"`
	XML              string        `json:"xml"`
	CreateTsp        bool          `json:"createTsp,omitempty"`
	UseTsaPolicy     Policy        `json:"useTsaPolicy,omitempty"`
	TSPHashAlgorithm HashAlgorithm `json:"tspHashAlgorithm,omitempty"`
}

type XMLSignResponse struct {
	apiResponse
	Result struct {
		XML string `json:"xml"`
		TSP string `json:"tsp"`
	} `json:"result"`
}

func (c *Client) XMLSign(p12, password, xml string, config *TSPConfig) (*XMLSignResponse, error) {
	if p12 == "" || password == "" || xml == "" {
		return nil, ErrInvalidRequestBody
	}

	if config == nil {
		config = &TSPConfig{}
	}

	body := apiRequest{
		Version: _v1,
		Method:  "XML.sign",
		Params: xmlSignRequest{
			P12:              p12,
			Password:         password,
			XML:              xml,
			CreateTsp:        config.Enabled,
			UseTsaPolicy:     config.Policy,
			TSPHashAlgorithm: config.HashAlgorithm,
		},
	}

	var reply XMLSignResponse
	if err := c.call(body, &reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
