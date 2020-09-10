package ncanode

// NodeInfoResponse describes json response from NodeInfo.
type NodeInfoResponse struct {
	apiResponse
	Datetime Time   `json:"dateTime"`
	Timezone string `json:"timezone"`
	Name     string `json:"name"`
	Version  string `json:"version"`
}

// NodeInfo returns NCANode server stats.
//
// See https://ncanode.kz/docs.php?go=52aec28e247d690426ead226d7631b421290eea3
func (c *Client) NodeInfo() (*NodeInfoResponse, error) {
	body := apiRequest{
		Version: c.version,
		Method:  "node.info",
	}

	var reply NodeInfoResponse
	if err := c.call(body, &reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
