package services

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetContratos(nup string) ([]byte, error) {

	const URL = "http://localhost:5556/getcontratos/{nup}?nup="
	resp, err := http.Get(URL + nup)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}
