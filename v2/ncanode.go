package ncanode

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type apiRequest struct {
	Version string      `json:"version"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type apiResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (r apiResponse) Error() string {
	return fmt.Sprintf("ncanode: request failed with status %d: %s", r.Status, r.Message)
}

// Time is a small wrapper of std time.
// Difference is time layout used by json decoding.
type Time struct{ time.Time }

// UnmarshalJSON implements custom unmarshaling of json decoder.
func (t *Time) UnmarshalJSON(data []byte) error {
	v := strings.Trim(string(data), "\"")
	if v == "null" {
		return nil
	}

	tt, err := time.Parse("2006-01-02 15:04:05", v)
	if err != nil {
		return err
	}

	*t = Time{tt}
	return nil
}

// MarshalJSON implements custom marshaling of json encoder.
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t)
}
