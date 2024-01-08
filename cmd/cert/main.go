package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
)

const (
	caKeyPath      = "./certs/ca_key.pem"
	caCertPath     = "./certs/ca_cert.pem"
	serverKeyPath  = "./certs/server_key.pem"
	serverCertPath = "./certs/server_cert.pem"
	clientKeyPath  = "./certs/client_key.pem"
	clientCertPath = "./certs/client_cert.pem"
	keyBitSize     = 2048
	validityPeriod = 365 * 24 * time.Hour // 1 year
	org            = "YandexPracticum"
	orgUnit        = "Golang20"
	country        = "RU"
	state          = "KBR"
	locality       = "Nalchik"
	hostnameServer = "localhost"
	hostnameClient = "localhost"
)

func main() {
	createRootCert()

	privServer := createPrivateKey()
	createServerCert(privServer)

	privClient := createPrivateKey()
	createClientCert(privClient)
}

func createRootCert() {
	caPrivKey := createPrivateKey()

	caTemplate := &x509.Certificate{
		SerialNumber:          serialNumber(),
		Subject:               pkixName(org, orgUnit, country, state, locality, hostnameServer),
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(validityPeriod),
		BasicConstraintsValid: true,
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageKeyAgreement,
	}

	caCert, err := x509.CreateCertificate(rand.Reader, caTemplate, caTemplate, caPrivKey.Public(), caPrivKey)
	if err != nil {
		log.Fatal(err)
	}

	saveCertAndKey(caCertPath, pemCert(caCert))
	saveCertAndKey(caKeyPath, pemPrivateKey(caPrivKey))

	fmt.Println("Создан корневой сертификат")
}

func createServerCert(privKey *rsa.PrivateKey) {
	caCert, caKey := loadCertAndKey(caCertPath, caKeyPath)

	serverTemplate := &x509.Certificate{
		SerialNumber:          serialNumber(),
		Subject:               pkixName(org, orgUnit, country, state, locality, hostnameServer),
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(validityPeriod),
		BasicConstraintsValid: true,
		IsCA:                  false,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		DNSNames:              []string{hostnameServer},
	}

	serverCert, err := x509.CreateCertificate(rand.Reader, serverTemplate, caCert, privKey.Public(), caKey)
	if err != nil {
		log.Fatal(err)
	}

	saveCertAndKey(serverCertPath, pemCert(serverCert))
	saveCertAndKey(serverKeyPath, pemPrivateKey(privKey))

	fmt.Println("Создан сертификат сервера")
}

func createClientCert(privKey *rsa.PrivateKey) {
	caCert, caKey := loadCertAndKey(caCertPath, caKeyPath)

	clientTemplate := &x509.Certificate{
		SerialNumber:          serialNumber(),
		Subject:               pkixName(org, orgUnit, country, state, locality, hostnameClient),
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(validityPeriod),
		BasicConstraintsValid: true,
		IsCA:                  false,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		DNSNames:              []string{hostnameClient},
	}

	clientCert, err := x509.CreateCertificate(rand.Reader, clientTemplate, caCert, privKey.Public(), caKey)
	if err != nil {
		log.Fatal(err)
	}

	saveCertAndKey(clientCertPath, pemCert(clientCert))
	saveCertAndKey(clientKeyPath, pemPrivateKey(privKey))

	fmt.Println("Создан сертификат клиента")
}

func createPrivateKey() *rsa.PrivateKey {
	privKey, err := rsa.GenerateKey(rand.Reader, keyBitSize)
	if err != nil {
		log.Fatal(err)
	}

	return privKey
}

func saveCertAndKey(path string, data []byte) {
	err := os.WriteFile(path, data, 0600)
	if err != nil {
		log.Fatal(err)
	}
}

func loadCertAndKey(certPath, keyPath string) (*x509.Certificate, *rsa.PrivateKey) {
	certBytes, err := os.ReadFile(certPath)
	if err != nil {
		log.Fatal(err)
	}

	keyBytes, err := os.ReadFile(keyPath)
	if err != nil {
		log.Fatal(err)
	}

	certBlock, _ := pem.Decode(certBytes)
	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	keyBlock, _ := pem.Decode(keyBytes)
	privKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	return cert, privKey
}

func pemCert(cert []byte) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	})
}

func pemPrivateKey(privKey *rsa.PrivateKey) []byte {
	privBytes := x509.MarshalPKCS1PrivateKey(privKey)
	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privBytes,
	})
}

func serialNumber() *big.Int {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatal(err)
	}
	return serialNumber
}

func pkixName(org, orgUnit, country, state, locality, commonName string) pkix.Name {
	return pkix.Name{
		CommonName:         commonName,
		SerialNumber:       "1",
		Organization:       []string{org},
		Country:            []string{country},
		Province:           []string{state},
		Locality:           []string{locality},
		OrganizationalUnit: []string{orgUnit},
	}
}
