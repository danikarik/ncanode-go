package ncanode_test

import (
	"testing"

	"github.com/danikarik/ncanode-go/v2"
	"github.com/stretchr/testify/require"
)

func TestCMSSign(t *testing.T) {
	client, err := ncanode.NewClient("http://127.0.0.1:14579")
	require.NoError(t, err)

	testCases := []struct {
		Name  string
		Path  string
		Alias string
	}{
		{
			Name:  "Personal/Active/Auth",
			Path:  "personal/active/AUTH_RSA256_b2c70a9ff7a5dc59de0ccc3c0ddde3437cc1f12d.p12",
			Alias: "b2c70a9ff7a5dc59de0ccc3c0ddde3437cc1f12d",
		},
		{
			Name:  "Personal/Active/Sign",
			Path:  "personal/active/RSA256_11c707bd54cbfcccb3815e39c4eb57b1dc7dfea9.p12",
			Alias: "11c707bd54cbfcccb3815e39c4eb57b1dc7dfea9",
		},
		{
			Name:  "Personal/Revoked/Auth",
			Path:  "personal/revoked/AUTH_RSA256_60d5a0346dd52fb5f3b9148e6cfbcd6cf323d119.p12",
			Alias: "60d5a0346dd52fb5f3b9148e6cfbcd6cf323d119",
		},
		{
			Name:  "Personal/Revoked/Sign",
			Path:  "personal/revoked/RSA256_346a7e57c2995259140b6fc375b6ff3bba7e852f.p12",
			Alias: "346a7e57c2995259140b6fc375b6ff3bba7e852f",
		},
		{
			Name:  "Organization/Active/Auth",
			Path:  "organization/active/AUTH_RSA256_d682726b55b3e62600ea8fe1c74e75fc96f47790.p12",
			Alias: "d682726b55b3e62600ea8fe1c74e75fc96f47790",
		},
		{
			Name:  "Organization/Active/Sign",
			Path:  "organization/active/GOSTKNCA_13a3741b52d7bf860aab199b56fad1f4652e357b.p12",
			Alias: "13a3741b52d7bf860aab199b56fad1f4652e357b",
		},
		{
			Name:  "Organization/Revoked/Auth",
			Path:  "organization/revoked/AUTH_RSA256_7d6d313ac5bf7367a2a69f28607a8deb80dd3ba9.p12",
			Alias: "7d6d313ac5bf7367a2a69f28607a8deb80dd3ba9",
		},
		{
			Name:  "Organization/Revoked/Sign",
			Path:  "organization/revoked/GOSTKNCA_c372b7809440fcb681051bbc89ea2f089fd2fe16.p12",
			Alias: "c372b7809440fcb681051bbc89ea2f089fd2fe16",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			p12, err := base64content(tc.Path)
			require.NoError(t, err)

			resp, err := client.CMSSign("YXNkYXNk", ncanode.P12Request{
				P12:      p12,
				Password: _defaultPassword,
				Alias:    tc.Alias,
			})
			require.NoError(t, err)
			require.NotEmpty(t, resp.CMS)
		})
	}
}
