package pem

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	fieldCertificate = "certificate"
	fieldPublicKey   = "public_key_pkix"
)

func newCertificateDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCertificateDataSource,
		Schema: map[string]*schema.Schema{
			fieldCertificate: {
				Type:         schema.TypeString,
				Description:  "PEM encoded certificate",
				Required:     true,
				ValidateFunc: validateCertificateDataSource,
			},
			fieldPublicKey: {
				Description: "PEM encoded PKIX public key for the certificate",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func parseCertificate(value interface{}) (id string, publicKey []byte, err error) {
	certBytes, ok := value.(string)

	if !ok {
		return "", nil, fmt.Errorf("field %s was not set or not a string", fieldCertificate)
	}

	certBlock, _ := pem.Decode([]byte(certBytes))

	if certBlock == nil || certBlock.Type != "CERTIFICATE" {
		return "", nil, fmt.Errorf("field %s was not a valid PEM-encoded certificate", fieldCertificate)
	}

	cert, err := x509.ParseCertificate(certBlock.Bytes)

	if err != nil {
		return "", nil, fmt.Errorf("unable to parse certificate. %v", err)
	}

	var pubKeyBlockBytes []byte
	switch cert.PublicKeyAlgorithm.String() {
	case "RSA":
		pk, ok := cert.PublicKey.(*rsa.PublicKey)
		if !ok {
			return "", nil, fmt.Errorf("public key algorithm is RSA, but could not convert")
		}

		pubKeyBlockBytes, err = x509.MarshalPKIXPublicKey(pk)
		if err != nil {
			return "", nil, fmt.Errorf("error marshalling public key")
		}

	default:
		return "", nil, fmt.Errorf("public key algorith, %s, is not supported", cert.PublicKeyAlgorithm.String())
	}

	pubKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBlockBytes,
	}

	pubKeyEncoded := pem.EncodeToMemory(pubKeyBlock)

	return cert.SerialNumber.String(), pubKeyEncoded, nil
}

func readCertificateDataSource(r *schema.ResourceData, _ interface{}) error {
	id, pubKey, err := parseCertificate(r.Get(fieldCertificate))

	if err != nil {
		return err
	}

	r.SetId(id)
	err = r.Set(fieldPublicKey, string(pubKey))

	if err != nil {
		return fmt.Errorf("error setting field, %s. %v", fieldPublicKey, err)
	}

	return nil
}

func validateCertificateDataSource(value interface{}, _ string) ([]string, []error) {
	if _, _, err := parseCertificate(value); err != nil {
		return nil, []error{err}
	}
	return nil, nil
}
