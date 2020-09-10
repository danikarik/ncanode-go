package ncanode

// P12Request holds p12 container info.
type P12Request struct {
	P12      string `json:"p12"`
	Password string `json:"password"`
	Alias    string `json:"alias,omitempty"`
}

type cmsSignRequest struct {
	Data string       `json:"data"`
	P12s []P12Request `json:"p12array"`
}

// CMSSignResponse describes json response from CMSSign.
type CMSSignResponse struct {
	apiResponse
	CMS string `json:"cms"`
}

// CMSSign creates cms signature.
//
// See https://ncanode.kz/docs.php?go=e3ebdc77c47d8fec8ce37ef76253a712790d5c89
func (c *Client) CMSSign(data string, p12s ...P12Request) (*CMSSignResponse, error) {
	if data == "" || len(p12s) == 0 {
		return nil, ErrInvalidRequestBody
	}

	body := apiRequest{
		Version: c.version,
		Method:  "cms.sign",
		Params: cmsSignRequest{
			Data: data,
			P12s: p12s,
		},
	}

	var reply CMSSignResponse
	if err := c.call(body, &reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
