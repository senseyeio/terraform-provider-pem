# PEM Terraform Provider

A terraform provider, which decodes a PEM encoded x509 certificate and makes its 
public key available

## Data Sources

### `pem_certifictate`

Inputs:

* `certificate` : The PEM encoded certificate to decode

Outputs:

* `public_key_pkix` : The PEM encoded PKIX public key for the certificate.
