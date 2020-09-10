package ncanode

type pkcs12Request struct {
	P12        string `json:"p12"`
	Password   string `json:"password"`
	VerifyOCSP bool   `json:"verifyOcsp"`
	VerifyCRL  bool   `json:"verifyCrl"`
}

// PKCS12Info returns P12 container info.
//
// See https://ncanode.kz/docs.php?go=9797704344e3175efb3260fbc5e58dac6c2fed3d
func (c *Client) PKCS12Info(p12, password string, verifyOCSP, verifyCRL bool) (*X509Response, error) {
	if p12 == "" || password == "" {
		return nil, ErrInvalidRequestBody
	}

	body := apiRequest{
		Version: c.version,
		Method:  "PKCS12.info",
		Params: pkcs12Request{
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
