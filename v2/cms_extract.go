package ncanode

type cmsExtractRequest struct {
	CMS string `json:"cms"`
}

// CMSExtractResponse describes json response from CMSExtract.
type CMSExtractResponse struct {
	apiResponse
	OriginalData string `json:"originalData"`
}

// CMSExtract returns original data from cms signature.
func (c *Client) CMSExtract(cms string) (*CMSExtractResponse, error) {
	if cms == "" {
		return nil, ErrInvalidRequestBody
	}

	body := apiRequest{
		Version: c.version,
		Method:  "cms.extract",
		Params:  cmsExtractRequest{CMS: cms},
	}

	var reply CMSExtractResponse
	if err := c.call(body, &reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
