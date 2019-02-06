package pem

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccDataCertificate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Providers:  testProviders,
		Steps: []resource.TestStep{
			{
				Config: dataCertificateValid(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.pem_certificate.cert", "id", expectedSerial),
					resource.TestCheckResourceAttr("data.pem_certificate.cert", "public_key_pkix", expectedPublicKey),
				),
			},
		},
	})
}

func dataCertificateValid() string {
	return fmt.Sprintf(`
data "pem_certificate" "cert" {
  certificate = <<CERT
%s
CERT
}`, exampleCertificate)
}

var testProviders = map[string]terraform.ResourceProvider{
	"pem": Provider(),
}

const exampleCertificate = `-----BEGIN CERTIFICATE-----
MIIDazCCAlOgAwIBAgIUZaL6+OPb7ommRCC2AhOJb59G9WIwDQYJKoZIhvcNAQEL
BQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xOTAyMDYxNTE4MTBaFw0yOTAy
MDMxNTE4MTBaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw
HwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwggEiMA0GCSqGSIb3DQEB
AQUAA4IBDwAwggEKAoIBAQC7pU2sZhpfg1pO9U8cXsDvresOsUr9VCfiU105HzQZ
z/o7MP/djIiTozXV/ERK7kHNX783ic82VPAVE96ycNvOtQsn88/DD6idyTT+j6kf
YYvLpPcTh+FgWdlvAwedjNN+3QjQaEwsbdnGRwKRFvotO8A0IF65ifTkJsqUXycn
tpQ2Uxh3AaloVhfBB360yeeJz4zYVUilSYbaRh6cfAYPDEVj7M7jFj+q6ZyR29vW
V3FkYpmIecRRDW6VJX48HWjfbzyEDvgzLSaclnqQj9Gy7CYM0yJipU0B6hZeUzN0
YNq2+C036TGNVZckTMewHkXVqpOoydRkkTNMz3m+IQfXAgMBAAGjUzBRMB0GA1Ud
DgQWBBT5/UP7SmMaU7H/d1ucegMGHa/1zjAfBgNVHSMEGDAWgBT5/UP7SmMaU7H/
d1ucegMGHa/1zjAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQCz
+pF1N9Ez+T2jvoL2B4qO+W1vZlskzXGO3VcQDgu+HmrN85iSRMSmCHUON8irsY1J
unaVqJezZ0mCxpy1addclwpHaMc1YnPMqHtMOARZAB1HoVj0QTcoRLD7c2w4gWCY
PRtAN8qTBGF2HdvlJadA7inFvHIRDEiziSpJjN77krJnPZZ93iqr5YEHHe1VVvt5
6CrWKkPVy50Ci4YaPp4ctfz3f0mufKZJQKk3tIi62GW54xLjGEUFGQipFDm50zrM
r0zBTzkiE6zQScTvTUKo6pN62zKYsQrPulU/B6cH07AoQu+SjejdTCu20WRxH2Sr
vlGp9JrpxkQhDKB0uZX8
-----END CERTIFICATE-----`

const expectedSerial = "580242651339756395525272749920392440808831645026"
const expectedPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAu6VNrGYaX4NaTvVPHF7A
763rDrFK/VQn4lNdOR80Gc/6OzD/3YyIk6M11fxESu5BzV+/N4nPNlTwFRPesnDb
zrULJ/PPww+onck0/o+pH2GLy6T3E4fhYFnZbwMHnYzTft0I0GhMLG3ZxkcCkRb6
LTvANCBeuYn05CbKlF8nJ7aUNlMYdwGpaFYXwQd+tMnnic+M2FVIpUmG2kYenHwG
DwxFY+zO4xY/qumckdvb1ldxZGKZiHnEUQ1ulSV+PB1o3288hA74My0mnJZ6kI/R
suwmDNMiYqVNAeoWXlMzdGDatvgtN+kxjVWXJEzHsB5F1aqTqMnUZJEzTM95viEH
1wIDAQAB
-----END PUBLIC KEY-----
`
