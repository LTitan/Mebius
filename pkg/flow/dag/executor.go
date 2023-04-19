package dag

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/LTitan/Mebius/pkg/utils/http"
	"github.com/go-resty/resty/v2"
)

var (
	blankRegexp = regexp.MustCompile(`\s`)
)

type Executor interface {
	Execute(string) (string, error)
}

type BashExecutor struct {
}

func (be *BashExecutor) Execute(command string) (string, error) {
	return be.executeWithStdout(command)
}

func (be *BashExecutor) executeWithStdout(command string) (string, error) {
	var stdout bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	err := cmd.Run()
	return strings.Trim(stdout.String(), "\n"), err
}

type HTTPExecutor struct {
	method  string
	url     string
	timeout int
	headers map[string]string
}

func NewHTTPExecutor(raw string, timeout int, headers map[string]string) (*HTTPExecutor, error) {
	sep := blankRegexp.Split(raw, -1)
	if len(sep) != 2 {
		return nil, errors.New("http type format error")
	}
	return &HTTPExecutor{
		method:  strings.ToUpper(sep[0]),
		url:     sep[1],
		timeout: 1,
		headers: headers,
	}, nil
}

func (he *HTTPExecutor) Execute(body string) (string, error) {
	return he.request([]byte(body))
}

func (he *HTTPExecutor) request(body []byte) (string, error) {
	cli := http.RestyClient.SetTimeout(time.Duration(he.timeout) * time.Second)
	request := cli.R().SetHeaders(he.headers).SetBody(body)
	var (
		resp *resty.Response
		err  error
	)
	switch he.method {
	case "GET":
		resp, err = request.Get(he.url)
	case "POST":
		resp, err = request.Post(he.url)
	case "PUT":
		resp, err = request.Put(he.url)
	case "DELETE":
		resp, err = request.Delete(he.url)
	}
	if err != nil {
		return "", err
	}
	if resp.StatusCode()/100 != 2 {
		return "", fmt.Errorf("response code is %d", resp.StatusCode())
	}
	return string(resp.Body()), nil
}
