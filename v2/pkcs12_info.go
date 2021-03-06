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
	Valid        bool        `json:"valid"`
	Alias        string      `json:"alias"`
	NotAfter     Time        `json:"notAfter"`
	NotBefore    Time        `json:"notBefore"`
	Chain        []Cert      `json:"chain"`
	KeyUsage     KeyUsage    `json:"keyUsage"`
	SerialNumber string      `json:"serialNumber"`
	Subject      Subject     `json:"subject"`
	SignAlg      string      `json:"signAlg"`
	Sign         string      `json:"sign"`
	PublicKey    string      `json:"publicKey"`
	Issuer       Subject     `json:"issuer"`
	KeyUser      []KeyUser   `json:"keyUser"`
	OCSP         *Revocation `json:"ocsp"`
	CRL          *Revocation `json:"crl"`
}

type pkcs12Request struct {
	P12       string `json:"p12"`
	Password  string `json:"password"`
	CheckOCSP bool   `json:"checkOcsp"`
	CheckCRL  bool   `json:"checkCrl"`
	Alias     string `json:"alias,omitempty"`
}

// X509Response describes json response from PKCS12Info.
type X509Response struct {
	apiResponse
	Cert Cert `json:"certificate"`
}

// PKCS12Info returns P12 container info.
//
// See https://ncanode.kz/docs.php?go=fa530e09377c651e57ac892137b850e2134d741b
func (c *Client) PKCS12Info(p12, password string, checkOCSP, checkCRL bool, alias string) (*X509Response, error) {
	if p12 == "" || password == "" {
		return nil, ErrInvalidRequestBody
	}

	body := apiRequest{
		Version: c.version,
		Method:  "info.pkcs12",
		Params: pkcs12Request{
			P12:       p12,
			Password:  password,
			CheckOCSP: checkOCSP,
			CheckCRL:  checkCRL,
			Alias:     alias,
		},
	}

	var reply X509Response
	if err := c.call(body, &reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
