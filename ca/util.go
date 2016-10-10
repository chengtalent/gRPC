package ca

import "crypto/x509"

func VerifySignature(c *x509.Certificate, caCert *x509.Certificate) error {
	return c.CheckSignatureFrom(caCert)
}
