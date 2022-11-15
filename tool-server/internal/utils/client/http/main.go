package http

import (
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	httpClient *resty.Client
)

func Get(url string, params map[string]string) (res []byte, err error) {
	logrus.Tracef("Http Get: %s, params: %+v", url, params)

	resp, err := httpClient.R().
		SetQueryParams(params).
		Get(url)
	if err != nil {
		logrus.Errorf("Http Get Error: %+v", err)
	} else {
		res = resp.Body()
	}
	return res, err
}

func Post(url string, params map[string]string, body any) []byte {
	logrus.Tracef("Http Post: %s, params: %+v, body: %+v", url, params, body)
	res, err := httpClient.R().
		SetQueryParams(params).
		SetBody(body).
		Post(url)
	if err != nil {
		logrus.Errorf("Http Post Error, url: %s, params: %+v, body: %+v", url, params, body)
	}
	return res.Body()
}

func SetHeaders(headers map[string]string) {
	httpClient.SetHeaders(headers)
}

func SetHeader(header string, value string) {
	httpClient.SetHeader(header, value)
}

func SetCookie(name string, value string) {
	httpClient.SetCookie(&http.Cookie{
		Name:  name,
		Value: value,
	})
}

func SetCookies(cookies map[string]string) {
	for k, v := range cookies {
		httpClient.SetCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}
}

func init() {
	if httpClient == nil {
		httpClient = resty.New()
	}
}
