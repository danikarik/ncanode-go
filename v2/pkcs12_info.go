package ncanode

type pkcs12Request struct {
	P12       string `json:"p12"`
	Password  string `json:"password"`
	CheckOCSP bool   `json:"checkOcsp"`
	CheckCRL  bool   `json:"checkCrl"`
	Alias     string `json:"alias,omitempty"`
}

// PKCS12Info returns P12 container info.
//
// See https://ncanode.kz/docs.php?go=fa530e09377c651e57ac892137b850e2134d741b
func (c *Client) PKCS12Info(p12, password string, checkOCSP, checkCRL bool, alias string) (*X509Response, error) {
	if p12 == "" || password == "" {
		return nil, ErrInvalidRequestBody
	}

	body := apiRequest{
		Version: c.version,
		Method:  "info.pkcs12",
		Params: pkcs12Request{
			P12:       p12,
			Password:  password,
			CheckOCSP: checkOCSP,
			CheckCRL:  checkCRL,
			Alias:     alias,
		},
	}

	var reply X509Response
	if err := c.call(body, &reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
