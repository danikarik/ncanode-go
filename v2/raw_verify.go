package ncanode

import "bytes"

type rawVerifyRequest struct {
	CMS        string `json:"cms"`
	VerifyOCSP bool   `json:"verifyOcsp"`
	VerifyCRL  bool   `json:"verifyCrl"`
}

// RawVerifyResponse describes json response from RawVerify.
type RawVerifyResponse struct {
	apiResponse
	Result struct {
		Valid bool `json:"valid"`
		Cert  Cert `json:"cert"`
	} `json:"result"`
}

// RawVerify validates cms signature.
//
// See https://ncanode.kz/docs.php?go=3f85fac8fa2729687ed307e791ce0fb17d704e26
func (c *Client) RawVerify(cms string, verifyOCSP, verifyCRL bool) (*RawVerifyResponse, error) {
	if cms == "" {
		return nil, ErrInvalidRequestBody
	}

	body := apiRequest{
		Version: c.version,
		Method:  "RAW.verify",
		Params: rawVerifyRequest{
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
