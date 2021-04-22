package service

import (
	"errors"
	"io"
	"log"
	"net/http"
)

type RestAPI struct {
	URI  string
	Body io.ReadCloser
}

func (api *RestAPI) Get() (err error) {

	client := http.Client{
		Timeout: http.DefaultClient.Timeout,
	}

	getRequest, err := http.NewRequest("GET", api.URI, nil)

	if err != nil {
		log.Println("Error in New Request", err)
		return
	}

	getResponse, err := client.Do(getRequest)

	if err != nil {
		log.Println("Error in Sending Request", "Error", err)
		return
	}

	if getResponse.StatusCode != http.StatusOK &&
		getResponse.StatusCode != http.StatusConflict {
		log.Println("HTTP response Error", "URI", api.URI, "Status", getResponse.StatusCode)
		err = errors.New("HTTP reponse Error")
		return
	}

	api.Body = getResponse.Body

	return
}
