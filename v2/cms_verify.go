package ncanode

type cmsVerifyRequest struct {
	CMS       string `json:"cms"`
	CheckOCSP bool   `json:"checkOcsp"`
	CheckCRL  bool   `json:"checkCrl"`
}

type cmsVerifyResponse struct {
	apiResponse
	Result struct {
		TSP     []interface{} `json:"tsp"`
		Signers []struct {
			Chain []Cert `json:"chain"`
			Cert  Cert   `json:"cert"`
		} `json:"signers"`
	} `json:"result"`
}

// CMSVerifyResponse describes json response from CMSVerify.
type CMSVerifyResponse struct {
	apiResponse
	TSP     []interface{} `json:"tsp"`
	Signers []Cert        `json:"signers"`
}

// CMSVerify validates cms signature.
//
// See https://ncanode.kz/docs.php?go=3698f11c5f30bd4879983a0194e1bfabefa3a753
func (c *Client) CMSVerify(cms string, checkOCSP, checkCRL bool) (*CMSVerifyResponse, error) {
	if cms == "" {
		return nil, ErrInvalidRequestBody
	}

	body := apiRequest{
		Version: c.version,
		Method:  "cms.verify",
		Params: cmsVerifyRequest{
			CMS:       cms,
			CheckOCSP: checkOCSP,
			CheckCRL:  checkCRL,
		},
	}

	var reply cmsVerifyResponse
	if err := c.call(body, &reply); err != nil {
		return nil, err
	}

	resp := &CMSVerifyResponse{
		apiResponse: apiResponse{
			Status:  reply.Status,
			Message: reply.Message,
		},
		TSP: reply.Result.TSP,
	}

	for _, signer := range reply.Result.Signers {
		signer.Cert.Chain = signer.Chain
		resp.Signers = append(resp.Signers, signer.Cert)
	}

	return resp, nil
}
