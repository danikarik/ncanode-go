package ncanode

// NodeInfoResponse describes json response from NodeInfo.
type NodeInfoResponse struct {
	apiResponse
	Result struct {
		Datetime Time   `json:"dateTime"`
		Timezone string `json:"timezone"`
		Name     string `json:"name"`
		Version  string `json:"version"`
	} `json:"result"`
}

// NodeInfo returns NCANode server stats.
//
// See https://ncanode.kz/docs.php?go=d8324d275a38b9c386071731e33afcaee4db7b50
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
