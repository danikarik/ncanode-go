package ncanode_test

import (
	"strings"
	"testing"

	"github.com/danikarik/ncanode-go/v2"
	"github.com/stretchr/testify/require"
)

func TestXMLVerify(t *testing.T) {
	client, err := ncanode.NewClient("http://127.0.0.1:14579")
	require.NoError(t, err)

	testCases := []struct {
		Name           string
		Path           string
		VerifyOCSP     bool
		VerifyCRL      bool
		ExpectedResult bool
	}{
		{
			Name:           "Personal/Active/Auth",
			Path:           "personal/active/AUTH_RSA256_b2c70a9ff7a5dc59de0ccc3c0ddde3437cc1f12d.xml",
			ExpectedResult: true,
		},
		{
			Name:           "Personal/Active/Sign",
			Path:           "personal/active/RSA256_11c707bd54cbfcccb3815e39c4eb57b1dc7dfea9.xml",
			ExpectedResult: true,
		},
		{
			Name:           "Personal/Revoked/Auth",
			Path:           "personal/revoked/AUTH_RSA256_60d5a0346dd52fb5f3b9148e6cfbcd6cf323d119.xml",
			VerifyOCSP:     true,
			VerifyCRL:      true,
			ExpectedResult: false,
		},
		{
			Name:           "Personal/Revoked/Sign",
			Path:           "personal/revoked/RSA256_346a7e57c2995259140b6fc375b6ff3bba7e852f.xml",
			VerifyOCSP:     true,
			VerifyCRL:      true,
			ExpectedResult: false,
		},
		{
			Name:           "Organization/Active/Auth",
			Path:           "organization/active/AUTH_RSA256_d682726b55b3e62600ea8fe1c74e75fc96f47790.xml",
			ExpectedResult: true,
		},
		{
			Name:           "Organization/Active/Sign",
			Path:           "organization/active/GOSTKNCA_13a3741b52d7bf860aab199b56fad1f4652e357b.xml",
			ExpectedResult: true,
		},
		{
			Name:           "Organization/Revoked/Auth",
			Path:           "organization/revoked/AUTH_RSA256_7d6d313ac5bf7367a2a69f28607a8deb80dd3ba9.xml",
			VerifyOCSP:     true,
			VerifyCRL:      true,
			ExpectedResult: false,
		},
		{
			Name:           "Organization/Revoked/Sign",
			Path:           "organization/revoked/GOSTKNCA_c372b7809440fcb681051bbc89ea2f089fd2fe16.xml",
			VerifyOCSP:     true,
			VerifyCRL:      true,
			ExpectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			xml, err := filecontent(tc.Path)
			require.NoError(t, err)

			if tc.VerifyOCSP || tc.VerifyCRL {
				xml = strings.Replace(xml, "NCANode", "NCANodeTest", -1)
			}

			resp, err := client.XMLVerify(xml, tc.VerifyOCSP, tc.VerifyCRL)
			require.NoError(t, err)
			require.Equal(t, tc.ExpectedResult, resp.Result.Valid)

			if tc.VerifyOCSP || tc.VerifyCRL {
				require.Equal(t, tc.ExpectedResult, resp.Result.Cert.Valid)
			}
		})
	}
}
