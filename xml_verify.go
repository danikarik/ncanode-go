package ncanode

type xmlVerifyRequest struct {
	XML        string `json:"xml"`
	VerifyOCSP bool   `json:"verifyOcsp"`
	VerifyCRL  bool   `json:"verifyCrl"`
}

// XMLVerifyResponse describes json response from XMLVerify.
type XMLVerifyResponse struct {
	apiResponse
	Result struct {
		Valid bool `json:"valid"`
		Cert  Cert `json:"cert"`
	} `json:"result"`
}

// XMLVerify validates xml signature.
//
// See https://ncanode.kz/docs.php?go=50acb512216c279acfa7eeb6de6dc2592039bd83
func (c *Client) XMLVerify(xml string, verifyOCSP, verifyCRL bool) (*XMLVerifyResponse, error) {
	if xml == "" {
		return nil, ErrInvalidRequestBody
	}

	body := apiRequest{
		Version: _v1,
		Method:  "XML.verify",
		Params: xmlVerifyRequest{
			XML:        xml,
			VerifyOCSP: verifyOCSP,
			VerifyCRL:  verifyCRL,
		},
	}

	var reply XMLVerifyResponse
	if err := c.call(body, &reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
