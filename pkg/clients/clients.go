package clients

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func FastRequest(host, msg string) (string, error) {
	u := url.URL{Scheme: "http", Host: host, Path: "/fast"}
	q := u.Query()
	q.Add("msg", msg)
	u.RawQuery = q.Encode()
	req, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func SlowRequest(host, msg string) (string, error) {
	u := url.URL{Scheme: "http", Host: host, Path: "/slow"}
	q := u.Query()
	q.Add("msg", msg)
	u.RawQuery = q.Encode()
	req, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
