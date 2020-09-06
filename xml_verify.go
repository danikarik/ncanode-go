package ncanode

type XMLVerifyRequest struct {
	XML        string `json:"xml"`
	VerifyOCSP bool   `json:"verifyOcsp"`
	VerifyCRL  bool   `json:"verifyCrl"`
}

type XMLVerifyResponse struct {
	APIResponse
	Result struct {
		Valid bool `json:"valid"`
		Cert  Cert `json:"cert"`
	} `json:"result"`
}

func (c *Client) XMLVerify(xml string, verifyOCSP, verifyCRL bool) (*XMLVerifyResponse, error) {
	if xml == "" {
		return nil, ErrInvalidRequestBody
	}

	body := APIRequest{
		Version: _v1,
		Method:  "XML.verify",
		Params: XMLVerifyRequest{
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
