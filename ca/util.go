package ca

import (
	"crypto/x509"
	"encoding/pem"
)

func VerifySignature(cert []byte, caCert []byte) error {
	c := BuildCertificateFromBytes(cert)
	caC := BuildCertificateFromBytes(caCert)

	return c.CheckSignatureFrom(caC)
}

func BuildCertificateFromBytes(cooked []byte) *x509.Certificate {
	block, _ := pem.Decode(cooked)
	cert, err := x509.ParseCertificate(block.Bytes)

	if err != nil {
		caLogger.Panic(err)
	}
	return cert
}

//func CreateKeyPairs(ca *CA, name string) (*ecdsa.PrivateKey, error) {
//	// read or create signing key pair
//	priv, err := ca.readCAPrivateKey(name)
//	if err != nil {
//		priv = ca.createCAKeyPair(name)
//	}
//
//	return priv, nil
//}