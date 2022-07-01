package types

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

var KeyUsageNameMap = map[int]string{
	0: "Digital Signature",
	1: "Non Repudiation",
	2: "Key Encipherment",
	3: "Data Encipherment",
	4: "Key Agreement",
	5: "Certificate Sign",
	6: "CRL Sign",
	7: "Encipher Only",
	8: "Decipher Only",
}

var ExtKeyUsagesNameMap = map[x509.ExtKeyUsage]string{
	x509.ExtKeyUsageAny:                            "Any",
	x509.ExtKeyUsageServerAuth:                     "TLS Web Server Authentication",
	x509.ExtKeyUsageClientAuth:                     "TLS Web Client Authentication",
	x509.ExtKeyUsageCodeSigning:                    "Code Signing",
	x509.ExtKeyUsageEmailProtection:                "E-mail Protection",
	x509.ExtKeyUsageIPSECEndSystem:                 "IPSec End System",
	x509.ExtKeyUsageIPSECTunnel:                    "IPSec Tunnel",
	x509.ExtKeyUsageIPSECUser:                      "IPSec User",
	x509.ExtKeyUsageTimeStamping:                   "Time Stamping",
	x509.ExtKeyUsageOCSPSigning:                    "OCSP Signing",
	x509.ExtKeyUsageMicrosoftServerGatedCrypto:     "Microsoft Server Gated Crypto",
	x509.ExtKeyUsageNetscapeServerGatedCrypto:      "Netscape Server Gated Crypto",
	x509.ExtKeyUsageMicrosoftCommercialCodeSigning: "Microsoft Commercial Code Signing",
	x509.ExtKeyUsageMicrosoftKernelCodeSigning:     "Microsoft Kernel Code Signing",
}
var KeyUsageArray = []x509.KeyUsage{
	x509.KeyUsageDigitalSignature,
	x509.KeyUsageContentCommitment,
	x509.KeyUsageKeyEncipherment,
	x509.KeyUsageDataEncipherment,
	x509.KeyUsageKeyAgreement,
	x509.KeyUsageCertSign,
	x509.KeyUsageCRLSign,
	x509.KeyUsageEncipherOnly,
	x509.KeyUsageDecipherOnly,
}

type CertFingerPrints struct {
	Type string `json:"hash" yaml:"hash"`
	FingerPrint string `json:"fingerprint" yaml:"fingerprint"`
}

type Cert struct {
	CommonName   string    `json:"commonName" yaml:"commonName"`
	AltNames     []string  `json:"altNames" yaml:"altNames"`
	NotBefore    time.Time `json:"notBefore" yaml:"notBefore"`
	NotAfter     time.Time `json:"notAfter" yaml:"notAfter"`
	KeyUsages    []string  `json:"keyUsages" yaml:"keyUsages"`
	ExtKeyUsages []string  `json:"extKeyUsages" yaml:"extKeyUsages"`
	SerialNumber string    `json:"serialNumber" yaml:"serialNumber"`
	Issuer       string    `json:"issuer" yaml:"issuer"`
	FingerPrints []CertFingerPrints `json:"fingerprints" yaml:"fingerprints"` 
}

func BuildCert(certificate *x509.Certificate) *Cert {
	buildCertificate := &Cert{}
	if certificate != nil {
		buildCertificate.CommonName = certificate.Subject.CommonName
		buildCertificate.AltNames = certificate.DNSNames
		buildCertificate.NotBefore = certificate.NotBefore
		buildCertificate.NotAfter = certificate.NotAfter
		buildCertificate.SerialNumber = fmt.Sprintf("%X", certificate.SerialNumber)
		buildCertificate.KeyUsages = buildKeyUsages(certificate.KeyUsage)
		buildCertificate.ExtKeyUsages = buildExtKeyUsages(certificate.ExtKeyUsage)
		buildCertificate.Issuer = certificate.Issuer.String()
		buildCertificate.FingerPrints = []CertFingerPrints{
			{
				Type: "sha1",
				FingerPrint: fmt.Sprintf("%X", sha1.Sum(certificate.Raw)),
			},
			{
				Type: "sha-256",
				FingerPrint: fmt.Sprintf("%X", sha256.Sum256(certificate.Raw)),
			},
		}
	}
	
	return buildCertificate
}

func buildExtKeyUsages(extKeyUsage []x509.ExtKeyUsage) []string {
	extKeyUsages := make([]string, len(extKeyUsage))
	for i := 0; i < len(extKeyUsage); i++ {
		extKeyUsages[i] = ExtKeyUsagesNameMap[extKeyUsage[i]]
	}
	return extKeyUsages
}

func buildKeyUsages(keyUsage x509.KeyUsage) []string {
	var keyUsages []string
	for i, ku := range KeyUsageArray {
		if keyUsage&ku != 0 {
			keyUsages = append(keyUsages, KeyUsageNameMap[i])
		}
	}
	return keyUsages
}

func (cert *Cert) ToTxt(withColors bool) string {
	var sb strings.Builder
	sb.WriteString("Certificate Information \n")
	sb.WriteString(formatString("\t Common Name: %s \n", cert.CommonName))
	sb.WriteString(formatString("\t Subject Alt Names: %s \n", strings.Join(cert.AltNames, "\n\t\t\t")))
	sb.WriteString(formatString("\t Not Before: %s \n ", checkBeforeDate(cert.NotBefore, withColors)))
	sb.WriteString(formatString("\t Not After: %s \n ", checkAfterDate(cert.NotAfter, withColors)))
	sb.WriteString(formatString("\t Serial Number: %s \n", cert.SerialNumber))
	sb.WriteString(formatString("\t Key Usages: %s \n", strings.Join(cert.KeyUsages, ",")))
	sb.WriteString(formatString("\t Extended Key Usages: %s \n", strings.Join(cert.ExtKeyUsages, ",")))
	sb.WriteString(formatString("\t Issuer: %s \n", cert.Issuer))
	for _, fingerPrint := range(cert.FingerPrints) {
		sb.WriteString(formatString("\t %s: %s \n", fingerPrint.Type, fingerPrint.FingerPrint))
	}
	return sb.String()

}

func checkBeforeDate(certBeforeDate time.Time, withColors bool) string {
	today := time.Now().UTC()
	isValid := today.After(certBeforeDate.UTC())
	dateString := certBeforeDate.UTC().Format(time.RFC3339)
	if isValid && withColors {
		return color.GreenString(dateString)
	}
	if !isValid && withColors {
		return color.RedString(dateString)
	}
	return dateString

}

func checkAfterDate(certAfterDate time.Time, withColors bool) string {
	today := time.Now().UTC()
	fmt.Println(today.Format(time.RFC3339))
	isValid := today.Before(certAfterDate.UTC())
	dateString := certAfterDate.UTC().Format(time.RFC3339)
	if isValid && withColors {
		return color.GreenString(dateString)
	}
	if !isValid && withColors {
		return color.RedString(dateString)
	}
	return dateString

}

func formatString(message string, a ...any) string {
	return fmt.Sprintf(message, a...)
}
