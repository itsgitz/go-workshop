package functions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/itsgitz/go-workshop/ory-kratos/models"
)

const (
	identityEndpoint = "http://127.0.0.1:4455/kratos/common/identities"
)

// GetIdentities function for get or fetch all ory kratos identities
func GetIdentities() {
	c := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, identityEndpoint, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(body))
	fmt.Println()

	kratos := []models.WebClientIdentity{}

	err = json.Unmarshal(body, &kratos)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(resp.Header)
	fmt.Println(len(kratos))

	if kratos == nil {

	}

	//fmt.Println("Get ID: ", kratos[0].ID)
}

// CreateIdentity for create new ORY Kratos Identity
func CreateIdentity(args []string) {
	// define the url/endpoint
	email := args[2]
	username := args[3]

	fmt.Println("Creating identity ...")

	c := &http.Client{
		Timeout: time.Second * 10,
	}

	// request body here as json data
	kratos := &models.WebClientIdentity{
		ID: "hello",
		RecoveryAddresses: []*models.WebClientRecoveryAddress{{
			Value: email,
			Via:   "email",
		}},
		Traits: models.WebClientTraits{
			Email:    email,
			Username: username,
		},
	}

	requestBody, err := json.Marshal(kratos)
	if err != nil {
		log.Println("encoding error", err)
		return
	}

	fmt.Println("sent request body:")
	fmt.Println(string(requestBody))

	req, err := http.NewRequest(http.MethodPost, identityEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("request error:", err)
		return
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Println("response error:", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("body error:", err)
		return
	}

	parseResponse := &models.WebClientIdentity{}
	err = json.Unmarshal(body, parseResponse)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("parseResponse:", parseResponse)
	if parseResponse.Error.Code != 0 {
		fmt.Println("code:", parseResponse.Error.Code)
		fmt.Println("message:", parseResponse.Error.Message)
		return
	}

	fmt.Println(string(body))
}
