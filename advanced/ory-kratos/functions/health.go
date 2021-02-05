package functions

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	statusAliveEndpoint = "http://127.0.0.1:4455/kratos/public/health/alive"
)

// GetStatusAlive
func GetStatusAlive() bool {
	tr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	c := http.Client{
		Transport: tr,
		Timeout:   time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, statusAliveEndpoint, nil)
	if err != nil {
		log.Fatal(err)
		return false
	}

	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
		return false
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return false
	}

	if len(body) == 0 {
		return false
	}

	return true
}
