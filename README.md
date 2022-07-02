# CertView

Simple X509 Viewer

## Usage
```bash
certview <flags> [certificate Path or Host]
One or more certificates are needed

Allowed flags
  -allowInsecureTLS
        Allow insecure TLS connections
  -no-color
        Disable color output
  -output string
        Ouput format (txt,json,yaml) (default "txt")
  -remote
        Given Arguments are urls an not files

```

### Examples

**View Remote SSL**
```bash
./certview -remote amazon.com:443
Certificate Information
         Common Name: *.peg.a2z.com
         Subject Alt Names: amazon.co.uk
                        uedata.amazon.co.uk
                        www.amazon.co.uk
                        origin-www.amazon.co.uk
                        *.peg.a2z.com
                        amazon.com
                        amzn.com
                        uedata.amazon.com
                        us.amazon.com
                        www.amazon.com
                        www.amzn.com
                        corporate.amazon.com
                        buybox.amazon.com
                        iphone.amazon.com
                        yp.amazon.com
                        home.amazon.com
                        origin-www.amazon.com
                        origin2-www.amazon.com
                        buckeye-retail-website.amazon.com
                        huddles.amazon.com
                        amazon.de
                        www.amazon.de
                        origin-www.amazon.de
                        amazon.co.jp
                        amazon.jp
                        www.amazon.jp
                        www.amazon.co.jp
                        origin-www.amazon.co.jp
                        *.aa.peg.a2z.com
                        *.ab.peg.a2z.com
                        *.ac.peg.a2z.com
                        origin-www.amazon.com.au
                        www.amazon.com.au
                        *.bz.peg.a2z.com
                        amazon.com.au
                        origin2-www.amazon.co.jp
         Not Before: 2021-10-06T00:00:00Z
         Not After: 2022-09-19T23:59:59Z
         Serial Number: E4239AB85E2E6A27C52C6DE9B9078D9
         Key Usages: Digital Signature,Key Encipherment
         Extended Key Usages: TLS Web Server Authentication,TLS Web Client Authentication
         Issuer: CN=DigiCert Global CA G2,O=DigiCert Inc,C=US
         sha1: 08040755C8B6852A5DB945A2B380571111DEFD2D
         sha-256: 5BF3D7E0E6927F773D5106C822C53F6F52C199F7EB1B3B8154B41F2924391C75
```

**View Remote, output as json**
```bash
./certview -remote -output json google.com:443
```
```json
[{"commonName":"*.peg.a2z.com","altNames":["amazon.co.uk","uedata.amazon.co.uk","www.amazon.co.uk","origin-www.amazon.co.uk","*.peg.a2z.com","amazon.com","amzn.com","uedata.amazon.com","us.amazon.com","www.amazon.com","www.amzn.com","corporate.amazon.com","buybox.amazon.com","iphone.amazon.com","yp.amazon.com","home.amazon.com","origin-www.amazon.com","origin2-www.amazon.com","buckeye-retail-website.amazon.com","huddles.amazon.com","amazon.de","www.amazon.de","origin-www.amazon.de","amazon.co.jp","amazon.jp","www.amazon.jp","www.amazon.co.jp","origin-www.amazon.co.jp","*.aa.peg.a2z.com","*.ab.peg.a2z.com","*.ac.peg.a2z.com","origin-www.amazon.com.au","www.amazon.com.au","*.bz.peg.a2z.com","amazon.com.au","origin2-www.amazon.co.jp"],"notBefore":"2021-10-06T00:00:00Z","notAfter":"2022-09-19T23:59:59Z","keyUsages":["Digital Signature","Key Encipherment"],"extKeyUsages":["TLS Web Server Authentication","TLS Web Client Authentication"],"serialNumber":"E4239AB85E2E6A27C52C6DE9B9078D9","issuer":"CN=DigiCert Global CA G2,O=DigiCert Inc,C=US","fingerprints":[{"hash":"sha1","fingerprint":"08040755C8B6852A5DB945A2B380571111DEFD2D"},{"hash":"sha-256","fingerprint":"5BF3D7E0E6927F773D5106C822C53F6F52C199F7EB1B3B8154B41F2924391C75"}]}]
```

**View local, output as Yaml**
```bash
./createview -output yaml twitter.pem
```
```yaml
- commonName: twitter.com
  altNames:
    - twitter.com
    - www.twitter.com
  notBefore: 2021-12-13T00:00:00Z
  notAfter: 2022-12-12T23:59:59Z
  keyUsages:
    - Digital Signature
    - Key Encipherment
  extKeyUsages:
    - TLS Web Server Authentication
    - TLS Web Client Authentication
  serialNumber: 630B4D3E1A04A3428146B4DBD1502B2
  issuer: CN=DigiCert TLS RSA SHA256 2020 CA1,O=DigiCert Inc,C=US
  fingerprints:
    - hash: sha1
      fingerprint: E3BA714291A065F576752BF0E0E18A1AE363A607
    - hash: sha-256
      fingerprint: D5D54531D10113F2ADBDF1B862440130920C4A5D4D6D924099585802FCAFB7CB
```
