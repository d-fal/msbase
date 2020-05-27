package utils

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"text/template"

	"msbase/pkg/conf"
)

type SoapHandler struct {
	SoapAction string
	URI        string
	SoapBody   interface{}
	request    *http.Request
}

var (
	handler          *SoapHandler
	tmpl             *template.Template
	lastTemplatePath string
)

// Prepare prepares the soap object
func Prepare(uri string, soapAction string) *SoapHandler {
	h := &SoapHandler{
		SoapAction: fmt.Sprintf("http://tempuri.org/%s", soapAction),
		URI:        uri,
	}
	return h
}

// SetPayload sets the params to be sending
func (h *SoapHandler) SetPayload(payload interface{}) {
	h.SoapBody = payload
}

// PrepareHandler Soap template would be generated
func (h *SoapHandler) PrepareHandler(templateName string, templatePath string) error {
	var err error
	var tpl bytes.Buffer

	if tmpl == nil || templatePath != lastTemplatePath {
		tmpl, err = template.New(templateName).ParseFiles(templatePath)
		if err != nil {
			fmt.Println(err)
			return err
		}
		lastTemplatePath = templatePath
	}

	err = tmpl.Execute(&tpl, &h.SoapBody)
	if err != nil {
		log.Println(err)
		return err
	}
	buffer := &bytes.Buffer{}
	enc := xml.NewEncoder(buffer)
	err = enc.Encode(tpl.String())
	if err != nil {
		log.Println("Error: ", err.Error())
		return err
	}

	r, err := http.NewRequest(http.MethodPost, h.URI, bytes.NewBuffer([]byte(tpl.String())))
	r.Header.Set("Content-type", "text/xml")
	r.Header.Set("SOAPAction", h.SoapAction)

	if err != nil {
		log.Println("Error making a requets", err)
		return err
	}
	h.request = r

	return nil
}

// Call dials the soap target
func (h *SoapHandler) Call() (*http.Response, error) {

	client := &http.Client{Timeout: GetTimeout()}
	resp, err := client.Do(h.request)

	if err != nil {
		log.Println(err)
	}
	return resp, err
}

// FetchResults This method calls and fetches the response
func (h *SoapHandler) FetchResults() ([]byte, conf.ErrorBlock) {
	var (
		errorBlock conf.ErrorBlock
	)
	resp, err := h.Call()

	if err != nil {
		err2 := err.(net.Error)
		if err2.Timeout() {
			log.Println("connection timeout. fetch_services.queryRestService", err)
			errorBlock = conf.GetConfigObject().GetErrorList().ErrorConnectionTimeout
			// conf.UpdateCompetency(userInput.Type, userInput.Competency, false)
			return nil, errorBlock
		}
		// conf.UpdateCompetency((*userInput).Type, (*userInput).Competency, false)
		errorBlock = conf.GetConfigObject().GetErrorList().ErrorCannotFetchHTTPData
		log.Println("pkg.utils.soap.FetchResults call error", err)
		return []byte{}, errorBlock
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		errorBlock = conf.GetConfigObject().GetErrorList().FailedToParseServerResponse
		log.Println("pkg.utils.soap.soap.go line 120 ", err)
		return []byte{}, errorBlock
	}
	// fmt.Println("SOAP", aurora.Green(string(body)))
	defer resp.Body.Close()

	return body, errorBlock

}
