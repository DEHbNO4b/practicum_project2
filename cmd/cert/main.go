package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

var (
	caCertPath = "./certs/caCert.pem"
	caKeyPath  = "./certs/caKey.pem"
)

// Шаблон сертификата
var certTemplate = &x509.Certificate{
	// указываем уникальный номер сертификата
	SerialNumber: big.NewInt(1658),
	// заполняем базовую информацию о владельце сертификата
	Subject: pkix.Name{
		Country:      []string{"RU"},
		Locality:     []string{"Nalchik"},
		Organization: []string{"Yandex.Praktikum"},
	},
	// разрешаем использование сертификата для 127.0.0.1 и ::1
	IPAddresses: []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
	// сертификат верен, начиная со времени создания
	NotBefore: time.Now(),
	// время жизни сертификата — 10 лет
	NotAfter:     time.Now().AddDate(10, 0, 0),
	SubjectKeyId: []byte{1, 2, 3, 4, 6},
	// устанавливаем использование ключа для цифровой подписи,
	// а также клиентской и серверной авторизации
	ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
	KeyUsage:    x509.KeyUsageDigitalSignature,
}

func main() {

	caCert, caKey, err := createCert(certTemplate, certTemplate)
	if err != nil {
		panic(err)
	}

	err = writeFile(caCertPath, caCert.Bytes())
	if err != nil {
		panic(err)
	}
	err = writeFile(caKeyPath, caKey.Bytes())
	if err != nil {
		panic(err)
	}

	ca, err := tls.LoadX509KeyPair(caCertPath, caKeyPath)

	srvCert, srvKey, err := createCert(certTemplate, ca)
	if err != nil {
		panic(err)
	}
	kfile, err := os.OpenFile("./certs/key.pem", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer kfile.Close()

	_, err = kfile.Write(privateKeyPEM.Bytes())
	if err != nil {
		panic(err)
	}
}

func createCert(template, parent *x509.Certificate) (bytes.Buffer, bytes.Buffer, error) {
	// используется rand.Reader в качестве источника случайных данных
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return bytes.Buffer{}, bytes.Buffer{}, err
	}

	// создаём сертификат x.509
	certBytes, err := x509.CreateCertificate(rand.Reader, template, parent, &privateKey.PublicKey, privateKey)
	if err != nil {
		return bytes.Buffer{}, bytes.Buffer{}, err
	}

	// кодируем сертификат и ключ в формате PEM, который
	// используется для хранения и обмена криптографическими ключами
	var certPEM bytes.Buffer
	pem.Encode(&certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})

	var privateKeyPEM bytes.Buffer
	pem.Encode(&privateKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	return certPEM, privateKeyPEM, nil
}

func writeFile(path string, data []byte) error {
	// file, err := os.OpenFile("./certs/ca.pem", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 664)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 664)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
