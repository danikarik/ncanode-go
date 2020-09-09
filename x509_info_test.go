package ncanode_test

import (
	"testing"

	"github.com/danikarik/ncanode-go"
	"github.com/stretchr/testify/require"
)

func TestX509Info(t *testing.T) {
	client, err := ncanode.NewClient("http://127.0.0.1:14579")
	require.NoError(t, err)

	testCases := []struct {
		Name       string
		Cert       string
		VerifyOCSP bool
		VerifyCRL  bool
	}{
		{
			Name:       "Default",
			Cert:       "MIIGZTCCBE2gAwIBAgIUFX1cSXpU/SdXs4r74PS8YFuVbAowDQYJKoZIhvcNAQELBQAwUjELMAkGA1UEBhMCS1oxQzBBBgNVBAMMOtKw0JvQotCi0KvSmiDQmtCj05jQm9CQ0J3QlNCr0KDQo9Co0Ksg0J7QoNCi0JDQm9Cr0pogKFJTQSkwHhcNMTgwODIyMTIxMTM2WhcNMTkwODIyMTIxMTM2WjCBpzEeMBwGA1UEAwwV0KLQldCh0KLQntCSINCi0JXQodCiMRUwEwYDVQQEDAzQotCV0KHQotCe0JIxGDAWBgNVBAUTD0lJTjEyMzQ1Njc4OTAxMTELMAkGA1UEBhMCS1oxFTATBgNVBAcMDNCQ0JvQnNCQ0KLQqzEVMBMGA1UECAwM0JDQm9Cc0JDQotCrMRkwFwYDVQQqDBDQotCV0KHQotCe0JLQmNCnMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtKWLOJf9qCqA6EO/SVtiMuPZ8q3Sg2RjO0dWXqKQRP7BWhIyMucMv+WmpRs8RuJ987Hm3B/JszSdiPrmtA9BpIERKphRwp3n4QR6pfLUBEp+5QNetNsv+dbiPcefWCzgJZCqEZVbPvSkiFH20y13YQ2FhEBUp4lLOqydBD2CsDVoTusvLanEgR+AdziJPq2+iXwhttpNPShKRTXGbGkxUa4P7YMUCUqWstR7svLaJqxKDMhaR7MpEt56a2pfntm5oFxKNFoBQjRXKbiBNIKciMRAeznjezv9ZA98WzWPIMuWzi38fPW5X7IVqa7ZbAFWvZIHWJmrl57uKGBNd9EUewIDAQABo4IB2zCCAdcwDgYDVR0PAQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMCBggqgw4DAwQBATAPBgNVHSMECDAGgARbanQRMB0GA1UdDgQWBBRrNhuGTGeWAbZS/jh/YfzZMDwDJzBeBgNVHSAEVzBVMFMGByqDDgMDAgQwSDAhBggrBgEFBQcCARYVaHR0cDovL3BraS5nb3Yua3ovY3BzMCMGCCsGAQUFBwICMBcMFWh0dHA6Ly9wa2kuZ292Lmt6L2NwczBWBgNVHR8ETzBNMEugSaBHhiFodHRwOi8vY3JsLnBraS5nb3Yua3ovbmNhX3JzYS5jcmyGImh0dHA6Ly9jcmwxLnBraS5nb3Yua3ovbmNhX3JzYS5jcmwwWgYDVR0uBFMwUTBPoE2gS4YjaHR0cDovL2NybC5wa2kuZ292Lmt6L25jYV9kX3JzYS5jcmyGJGh0dHA6Ly9jcmwxLnBraS5nb3Yua3ovbmNhX2RfcnNhLmNybDBiBggrBgEFBQcBAQRWMFQwLgYIKwYBBQUHMAKGImh0dHA6Ly9wa2kuZ292Lmt6L2NlcnQvbmNhX3JzYS5jZXIwIgYIKwYBBQUHMAGGFmh0dHA6Ly9vY3NwLnBraS5nb3Yua3owDQYJKoZIhvcNAQELBQADggIBACy0Lxj0D/q3SwUz0X9BICyKPw/U6sXmedqUcrghzZuT9ojnUp9w7g4ndZOKTRRxQyLiUYb9neJ3SGVuF/XYcs7Ovrp5RGNNHuVUR8bQz9cbWd/O2qRUY6qlg4ZSjYsjFYaQm8o+uO56PuqWG125O7XNUdAUHNBc2hUrrngGKU0FKxlBygxLpvTf4I9q3QA0PJ6MnHrUKlor4sRGar4hMJCbrxeMG4pv3Jx/r9fsKy7f+yZeQo3T4XAIXmUTXF8UC3HtIroxAP6yEoEhG76oS3qvYc1K/krI48ju5VYxmzEabNqRhiiEBpocIwCqFLLo9x3CKuUkuA7pwEib4YcCNxCTucCtd9x8dGgZRNffJV4de/Aja/VP84q8rxmcyogbUQzvPb+2/zKRh6cxYxnRsuL4wWUV+fxp/usy0mJMboQF7IcRFe1fXosU0RWYmKHITOCDbs0NKxTn7TSxEKMYdJN6YngCmKlmwR/+AfxhN1QMSQpU/m8Glwl+f5wZIL5MQJVhrrWIteh0tnb+OuDQHz4g2vmD2xq5jUQDFIrjXOdy4zToqM6tirt3nGDsblWgcgsPac50FLT1+um7W26UsmtZ9/wXvkxYC9kL5gUX53VD/bcKki8fogjrNoYEZiORRqmwvZ5EVe4w3Hfb7YCnc3NzhhIg6hqmzumXNCgLCt2q",
			VerifyOCSP: false,
			VerifyCRL:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			resp, err := client.X509Info(tc.Cert, tc.VerifyOCSP, tc.VerifyCRL)
			require.NoError(t, err)
			require.False(t, resp.Result.Valid)
			require.Equal(t, ncanode.KeyUsageAuth, resp.Result.KeyUsage)
			require.Equal(t, "122684438670642568061334282296011886211357830154", resp.Result.SerialNumber)

			t.Log(resp.Result.OCSP)
			t.Log(resp.Result.CRL)
		})
	}
}
