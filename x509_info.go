package ncanode

type KeyUsage string

const (
	AUTH    KeyUsage = "AUTH"
	SIGN    KeyUsage = "SIGN"
	UNKNOWN KeyUsage = "UNKNOWN"
)

type Gender string

const (
	MALE   Gender = "MALE"
	FEMALE Gender = "FEMALE"
)

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

type KeyUser string

const (
	INDIVIDUAL       KeyUser = "INDIVIDUAL"
	ORGANIZATION     KeyUser = "ORGANIZATION"
	CEO              KeyUser = "CEO"
	CANSIGN          KeyUser = "CAN_SIGN"
	CANSIGNFINANCIAL KeyUser = "CAN_SIGN_FINANCIAL"
	HR               KeyUser = "HR"
	EMPLOYEE         KeyUser = "EMPLOYEE"
	NCAPRIVILEGES    KeyUser = "NCA_PRIVILEGES"
	NCAADMIN         KeyUser = "NCA_ADMIN"
	NCAMANAGER       KeyUser = "NCA_MANAGER"
	NCAOPERATOR      KeyUser = "NCA_OPERATOR"
)

type X509Response struct {
	APIResponse
	Result struct {
		Valid        bool           `json:"valid"`
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
	} `json:"result"`
}

type X509Request struct {
	Cert       string `json:"cert"`
	VerifyOCSP bool   `json:"verifyOcsp"`
	VerifyCRL  bool   `json:"verifyCrl"`
}

func (c *Client) X509Info(cert string, verifyOCSP, verifyCRL bool) (*X509Response, error) {
	if cert == "" {
		return nil, ErrInvalidRequestBody
	}

	body := APIRequest{
		Version: _v1,
		Method:  "X509.info",
		Params: X509Request{
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
