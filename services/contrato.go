package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func PostContrato(data url.Values) error {

	const URL = "http://localhost:5556/postcontrato/"
	res, err := http.PostForm(URL, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	return nil
}
