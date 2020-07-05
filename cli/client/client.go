package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yunussandikci/shareterm/common"
	"io/ioutil"
	"net/http"
)

type Client interface {
	CreatePaste(paste string) common.PasteCreateResponse
}

type client struct {
	BaseUrl    string
	HttpClient http.Client
}

func NewClient() Client {

	return client{
		HttpClient: http.Client{},
		BaseUrl: common.GetHost(),
	}
}

func (e client) CreatePaste(paste string) common.PasteCreateResponse {
	pasteCreateRequest := common.PasteCreateRequest{
		Content: paste,
	}
	pasteCreateBody, requestMarshalErr := json.Marshal(pasteCreateRequest)
	if requestMarshalErr != nil {
		panic(requestMarshalErr)
	}
	request, requestErr :=
		http.NewRequest("POST", fmt.Sprintf("%s/api", e.BaseUrl), bytes.NewBuffer(pasteCreateBody))
	if requestErr != nil {
		panic(requestErr)
	}
	request.Header.Set("Content-Type", "application/json")
	response, clientErr := e.HttpClient.Do(request)
	if clientErr != nil {
		panic(clientErr)
	}
	body, bodyReadErr := ioutil.ReadAll(response.Body)
	if bodyReadErr != nil {
		panic(bodyReadErr)
	}
	pasteCreateResponse := common.PasteCreateResponse{}
	responseUnmarshalError := json.Unmarshal(body, &pasteCreateResponse)
	if responseUnmarshalError != nil {
		panic(responseUnmarshalError)
	}
	return pasteCreateResponse
}
