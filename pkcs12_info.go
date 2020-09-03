package ncanode

type PKCS12Request struct {
	P12        string `json:"p12"`
	Password   string `json:"password"`
	VerifyOCSP bool   `json:"verifyOcsp"`
	VerifyCRL  bool   `json:"verifyCrl"`
}

func (c *Client) PKCS12Info(p12, password string, verifyOCSP, verifyCRL bool) (*X509Response, error) {
	if p12 == "" || password == "" {
		return nil, ErrInvalidRequestBody
	}

	body := APIRequest{
		Version: _v1,
		Method:  "PKCS12.info",
		Params: PKCS12Request{
			P12:        p12,
			Password:   password,
			VerifyOCSP: verifyOCSP,
			VerifyCRL:  verifyCRL,
		},
	}

	var reply X509Response
	if err := c.call(body, &reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
