package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/cortiz/certview/internal/types"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
func printError(message string) {
	fmt.Printf("%s \n", color.RedString(message))
}
func printHelp() {
	fmt.Println("certview <flags> [certificate Path or Host]")
	fmt.Println()
	printError("One or more certificates are needed")
}
func readCertFile(certFilePath string) *x509.Certificate {
	if fileExists(certFilePath) {
		r, err := ioutil.ReadFile(certFilePath)

		if err != nil {
			printError(err.Error())
			os.Exit(2)
		}
		block, _ := pem.Decode(r)
		if block == nil {
			printError("Unable to read " + certFilePath)
			os.Exit(3)
		}

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			printError(err.Error())
		}
		return cert
	}
	printError("File " + certFilePath + " not found")
	return nil
}

func readRemoteCertFile(remoteHost string, allowInsecureTLS bool) *x509.Certificate {

	conf := &tls.Config{
		InsecureSkipVerify: allowInsecureTLS,
	}
	conn, err := tls.Dial("tcp", remoteHost, conf)
	if err != nil {
		printError("Error in Dial " + err.Error())
		return nil
	}

	defer conn.Close()
	return conn.ConnectionState().PeerCertificates[0]
}

func main() {
	var flagNoColor = flag.Bool("no-color", false, "Disable color output")
	var useUrl = flag.Bool("remote", false, "Given Arguments are urls an not files")
	var allowInsecureTLS = flag.Bool("allowInsecureTLS", false, "Allow insecure TLS connections")
	var outputFormat = flag.String("output", "txt", "Ouput format (txt,json,yaml)")
	flag.Parse()
	if *flagNoColor {
		color.NoColor = true // disables colorized output
	}

	if len(os.Args) <= 1 {
		printHelp()
		fmt.Println()
		fmt.Println("Allowed flags")
		flag.PrintDefaults()
		return
	}

	var certificates []types.Cert

	for _, a := range flag.Args() {
		var cert *x509.Certificate
		if *useUrl {
			cert = readRemoteCertFile(a, *allowInsecureTLS)
		} else {
			cert = readCertFile(a)
		}
		if cert != nil {
			certificate := types.BuildCert(cert)
			certificates = append(certificates, *certificate)
		}
	}
	switch strings.ToLower(*outputFormat) {
	case "json":
		outputJson(certificates)
	case "txt":
		for _, cert := range certificates {
			fmt.Println()
			fmt.Println(cert.ToTxt(!*flagNoColor))
			fmt.Println()
		}
	case "yaml":
		outputYaml(certificates)
	default:
		fmt.Println("Unsupported OutFormant")
	}
}

func outputYaml(certificates []types.Cert) {
	byte, err := yaml.Marshal(certificates)
	if err != nil {
		fmt.Printf("Unable to output yaml %s", err)
	}
	fmt.Println(string(byte))
}

func outputJson(certificates []types.Cert) {
	byte, err := json.Marshal(certificates)
	if err != nil {
		fmt.Printf("Unable to output json %s", err)
	}
	fmt.Println(string(byte))

}
