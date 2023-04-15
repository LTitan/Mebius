package http

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	resty "github.com/go-resty/resty/v2"
)

const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
	MIMEPROTOBUF          = "application/x-protobuf"
	MIMEMSGPACK           = "application/x-msgpack"
	MIMEMSGPACK2          = "application/msgpack"
)

// RestyClient default request content-type: application/json
var RestyClient *resty.Client

func init() {
	httpClient := &http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			MaxIdleConns:        2000,
			MaxIdleConnsPerHost: 5,
			MaxConnsPerHost:     100,
			DisableKeepAlives:   true,
			DialContext: (&net.Dialer{
				Timeout: 3 * time.Second,
			}).DialContext,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}
	RestyClient = resty.NewWithClient(httpClient)
	RestyClient.SetHeader("Content-Type", MIMEJSON)
}

// NewRestyClient .
func NewRestyClient() *resty.Client {
	cli := resty.New()
	return cli
}
