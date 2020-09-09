package ncanode

type NodeInfoResponse struct {
	apiResponse
	Result struct {
		Datetime Time   `json:"dateTime"`
		Timezone string `json:"timezone"`
		Name     string `json:"name"`
		Version  string `json:"version"`
	} `json:"result"`
}

func (c *Client) NodeInfo() (*NodeInfoResponse, error) {
	body := apiRequest{
		Version: _v1,
		Method:  "NODE.info",
	}

	var reply NodeInfoResponse
	if err := c.call(body, &reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
