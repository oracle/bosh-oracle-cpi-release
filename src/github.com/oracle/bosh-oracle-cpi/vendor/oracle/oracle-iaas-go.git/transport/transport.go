package transport

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/f2prateek/train"
	//	"log"
)

// Config API authentication and target configuration
type Config struct {
	Tenant      string `ini:"tenancy" mapstructure:"tenant"`
	User        string `ini:"user" mapstructure:"user"`
	Fingerprint string `ini:"fingerprint" mapstructure:"fingerprint"`
	Host        string `ini:"host" mapstructure:"host"`
	// File location of the pem private  key
	KeyFile string `ini:"key" mapstructure:"key"`
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error {
	return nil
}

func parsePrivateKey(content []byte) (*rsa.PrivateKey, error) {
	keyBlock, _ := pem.Decode(content)
	der := keyBlock.Bytes

	if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}
	if key, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		switch key := key.(type) {
		case *rsa.PrivateKey:
			return key, nil
		default:
			return nil, errors.New("Failed ot parse key")
		}
	}
	return nil, errors.New("failed to parse private key")
}

// CreateAuthenticatedHTTPTarget Creates an HTTP endpoint that adds bare metal signature authentication to each outgoing request.
//
func CreateAuthenticatedHTTPTarget(target http.RoundTripper, config Config) (http.RoundTripper, error) {

	keycontent, err := ioutil.ReadFile(config.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to load private key from %s : %s", config.KeyFile, err.Error())
	}

	key, err := parsePrivateKey(keycontent)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse private key from %s : %s", config.KeyFile, err.Error())
	}

	signerFn := func(chain train.Chain) (resp *http.Response, err error) {
		req := chain.Request()
		var buf *bytes.Buffer

		if req.Body != nil {
			buf = new(bytes.Buffer)
			buf.ReadFrom(req.Body)
			req.Body = nopCloser{buf}
		}
		if req.Header.Get("date") == "" {
			req.Header.Set("date", time.Now().UTC().Format(http.TimeFormat))
		}
		if req.Header.Get("content-type") == "" {
			req.Header.Set("content-type", "application/json")
		}
		if req.Header.Get("accept") == "" {
			req.Header.Set("accept", "application/json")
		}

		if req.Header.Get("host") == "" {
			req.Header.Set("host", req.URL.Host)
		}
		var signheaders []string
		if (req.Method == "PUT" || req.Method == "POST") && buf != nil {
			signheaders = []string{"(request-target)", "host", "date", "content-length", "content-type", "x-content-sha256"}
			if req.Header.Get("content-length") == "" {
				req.Header.Set("content-length", fmt.Sprintf("%d", buf.Len()))
			}
			hasher := sha256.New()
			hasher.Write(buf.Bytes())

			req.Header.Set("x-content-sha256", base64.StdEncoding.EncodeToString(hasher.Sum(nil)))
		} else {
			signheaders = []string{"date", "host", "(request-target)"}
		}

		var signbuffer bytes.Buffer
		for idx, header := range signheaders {
			signbuffer.WriteString(header)
			signbuffer.WriteString(": ")
			if header == "(request-target)" {

				signbuffer.WriteString(strings.ToLower(req.Method))
				signbuffer.WriteString(" ")
				signbuffer.WriteString(req.URL.RequestURI())
			} else {
				signbuffer.WriteString(req.Header.Get(header))
			}
			if idx < len(signheaders)-1 {
				signbuffer.WriteString("\n")
			}
		}

		h := sha256.New()
		h.Write(signbuffer.Bytes())
		digest := h.Sum(nil)
		signature, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, digest)
		if err != nil {
			return
		}
		req.Header.Add("Authorization",
			fmt.Sprintf("Signature headers=\"%s\",keyId=\"%s/%s/%s\",algorithm=\"rsa-sha256\",signature=\"%s\",version=\"1\"",
				strings.Join(signheaders, " "), config.Tenant, config.User, config.Fingerprint, base64.StdEncoding.EncodeToString(signature)))

		dreq, _ := httputil.DumpRequestOut(req, true)
		logrus.Debug(string(dreq))
		thersp, err := chain.Proceed(req)
		if err != nil {
			logrus.Debug(err)
			return nil, err
		}
		dresp, _ := httputil.DumpResponse(thersp, true)
		logrus.Debug(string(dresp))
		return thersp, err
	}
	return train.TransportWith(target, train.InterceptorFunc(signerFn)), nil
}
