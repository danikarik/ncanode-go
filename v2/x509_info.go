package ncanode

// KeyUsage is an alias of digital key type.
type KeyUsage string

// List of values KeyUsage can take.
const (
	KeyUsageAuth    KeyUsage = "AUTH"
	KeyUsageSign    KeyUsage = "SIGN"
	KeyUsageUnknown KeyUsage = "UNKNOWN"
)

// Gender is an alias of person gender.
type Gender string

// List of values Gender can take.
const (
	GenderMale   Gender = "MALE"
	GenderFemale Gender = "FEMALE"
)

// Subject holds person or organization data.
type Subject struct {
	LastName     string `json:"lastName,omitempty"`
	Country      string `json:"country,omitempty"`
	CommonName   string `json:"commonName,omitempty"`
	Gender       Gender `json:"gender,omitempty"`
	Surname      string `json:"surname,omitempty"`
	Locality     string `json:"locality,omitempty"`
	DN           string `json:"dn,omitempty"`
	State        string `json:"state,omitempty"`
	BirthDate    string `json:"birthDate,omitempty"`
	IIN          string `json:"iin,omitempty"`
	BIN          string `json:"bin,omitempty"`
	Organization string `json:"organization,omitempty"`
	Email        string `json:"email,omitempty"`
}

// KeyUser is an alias of user type.
type KeyUser string

// List of values KeyUser can take.
const (
	KeyUserIndividual       KeyUser = "INDIVIDUAL"
	KeyUserOrganization     KeyUser = "ORGANIZATION"
	KeyUserCEO              KeyUser = "CEO"
	KeyUserCanSign          KeyUser = "CAN_SIGN"
	KeyUserCanSignFinancial KeyUser = "CAN_SIGN_FINANCIAL"
	KeyUserHR               KeyUser = "HR"
	KeyUserEmployee         KeyUser = "EMPLOYEE"
	KeyUserNCAPrivileges    KeyUser = "NCA_PRIVILEGES"
	KeyUserNCAAdmin         KeyUser = "NCA_ADMIN"
	KeyUserNCAManager       KeyUser = "NCA_MANAGER"
	KeyUserNCAOperator      KeyUser = "NCA_OPERATOR"
)

// Status is an alias of revocation status.
type Status string

// List of values Status can take.
const (
	StatusUnknown Status = "UNKNOWN"
	StatusActive  Status = "ACTIVE"
	StatusRevoked Status = "REVOKED"
)

// Revocation holds data of revoked certificate.
type Revocation struct {
	Reason    interface{} `json:"revokationReason"`
	Time      Time        `json:"revokationTime"`
	RevokedBy string      `json:"revokedBy,omitempty"`
	Status    Status      `json:"status"`
}

// Cert holds data of certificate.
type Cert struct {
	Valid        bool           `json:"valid"`
	Alias        string         `json:"alias"`
	NotAfter     Time           `json:"notAfter"`
	NotBefore    Time           `json:"notBefore"`
	Chain        []X509Response `json:"chain"`
	KeyUsage     KeyUsage       `json:"keyUsage"`
	SerialNumber string         `json:"serialNumber"`
	Subject      Subject        `json:"subject"`
	SignAlg      string         `json:"signAlg"`
	Sign         string         `json:"sign"`
	PublicKey    string         `json:"publicKey"`
	Issuer       Subject        `json:"issuer"`
	KeyUser      []KeyUser      `json:"keyUser"`
	OCSP         *Revocation    `json:"ocsp"`
	CRL          *Revocation    `json:"crl"`
}

// X509Response describes json response from X509Info.
type X509Response struct {
	apiResponse
	Cert Cert `json:"certificate"`
}

type x509Request struct {
	Cert       string `json:"cert"`
	VerifyOCSP bool   `json:"verifyOcsp"`
	VerifyCRL  bool   `json:"verifyCrl"`
}

// X509Info returns certifacate info.
//
// See https://ncanode.kz/docs.php?go=68c0077b854fcdb23c567751b1329be3a34447c0
func (c *Client) X509Info(cert string, verifyOCSP, verifyCRL bool) (*X509Response, error) {
	if cert == "" {
		return nil, ErrInvalidRequestBody
	}

	body := apiRequest{
		Version: c.version,
		Method:  "X509.info",
		Params: x509Request{
			Cert:       cert,
			VerifyOCSP: verifyOCSP,
			VerifyCRL:  verifyCRL,
		},
	}

	var reply X509Response
	if err := c.call(body, &reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
