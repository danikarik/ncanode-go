package ncanode

import (
	"encoding/json"
	"strings"
	"time"
)

// List of versions API can take.
const (
	_v1 string = "1.0"
	_v2 string = "2.0"
)

type APIRequest struct {
	Version string      `json:"version"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type APIResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	v := strings.Trim(string(data), "\"")

	tt, err := time.Parse("2006-01-02 15:04:05", v)
	if err != nil {
		return err
	}

	*t = Time{tt}
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t)
}
