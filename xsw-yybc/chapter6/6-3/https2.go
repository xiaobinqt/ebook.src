package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("https2 success..."))
}

func YourListenAndServeTLS(addr, certFile, keyFile string,
	handler http.Handler) (err error) {
	config := &tls.Config{
		Rand:       rand.Reader,
		Time:       time.Now,
		NextProtos: []string{"http/1.1"},
	}

	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], err = YourLoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}

	conn, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	tlsListener := tls.NewListener(conn, config)
	return http.Serve(tlsListener, handler)
}

func YourLoadX509KeyPair(certFile, keyFile string) (cert tls.Certificate, err error) {
	certPEMBlock, err := ioutil.ReadFile(certFile)
	if err != nil {
		return
	}

	certDERBlock, restPEMBlock := pem.Decode(certPEMBlock)
	if certDERBlock == nil {
		err = fmt.Errorf("crypto/tls: failed to parse certificate pem data")
		return
	}

	certDERBlockChain, _ := pem.Decode(restPEMBlock)
	if certDERBlockChain == nil {
		cert.Certificate = [][]byte{
			certDERBlock.Bytes,
		}
	} else {
		cert.Certificate = [][]byte{
			certDERBlock.Bytes,
			certDERBlockChain.Bytes,
		}
	}

	keyPEMBlock, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return
	}

	keyDERBlock, _ := pem.Decode(keyPEMBlock)
	if keyDERBlock == nil {
		err = fmt.Errorf("crypto/tls: failed to parse key PEM data")
		return
	}

	key, err := x509.ParsePKCS1PrivateKey(keyDERBlock.Bytes)
	if err != nil {
		err = fmt.Errorf("crypto/tls: failed to parse key")
		return
	}

	cert.PrivateKey = key
	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		err = fmt.Errorf("x509.ParseCertificates(certDERBlock.Bytes) err")
		return
	}

	if x509Cert.PublicKeyAlgorithm != x509.RSA ||
		x509Cert.PublicKey.(*rsa.PublicKey).N.Cmp(key.PublicKey.N) != 0 {
		err = fmt.Errorf("crypto/tls: private key does not match public key")
		return
	}

	return
}

func main() {
	http.HandleFunc("/", rootHandler)
	YourListenAndServeTLS(":8081", "xsw.crt",
		"xsw.key", nil)
}
