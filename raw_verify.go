package ncanode

import "bytes"

type RawVerifyRequest struct {
	CMS        string `json:"cms"`
	VerifyOCSP bool   `json:"verifyOcsp"`
	VerifyCRL  bool   `json:"verifyCrl"`
}

type RawVerifyResponse struct {
	APIResponse
	Result struct {
		Valid bool `json:"valid"`
		Cert  Cert `json:"cert"`
	} `json:"result"`
}

func (c *Client) RawVerify(cms string, verifyOCSP, verifyCRL bool) (*RawVerifyResponse, error) {
	if cms == "" {
		return nil, ErrInvalidRequestBody
	}

	body := APIRequest{
		Version: _v1,
		Method:  "RAW.verify",
		Params: RawVerifyRequest{
			CMS:        cms,
			VerifyOCSP: verifyOCSP,
			VerifyCRL:  verifyCRL,
		},
	}

	mod := func(in []byte) ([]byte, error) {
		return bytes.Replace(in, []byte(`\\`), []byte(`\`), -1), nil
	}

	var reply RawVerifyResponse
	if err := c.call(body, &reply, mod); err != nil {
		return nil, err
	}

	return &reply, nil
}
