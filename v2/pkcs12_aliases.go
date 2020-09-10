package ncanode

type pkcs12aliasesRequest struct {
	P12      string `json:"p12"`
	Password string `json:"password"`
}

// PKCS12AliasesResponse describes json response from PKCS12Aliases.
type PKCS12AliasesResponse struct {
	apiResponse
	Aliases []string `json:"aliases"`
}

// PKCS12Aliases returns P12 container aliases.
//
// See https://ncanode.kz/docs.php?go=f7a72200e668edcfa02023a5a44321b4e280d11b
func (c *Client) PKCS12Aliases(p12, password string) (*PKCS12AliasesResponse, error) {
	if p12 == "" || password == "" {
		return nil, ErrInvalidRequestBody
	}

	body := apiRequest{
		Version: c.version,
		Method:  "info.pkcs12aliases",
		Params: pkcs12Request{
			P12:      p12,
			Password: password,
		},
	}

	var reply PKCS12AliasesResponse
	if err := c.call(body, &reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
