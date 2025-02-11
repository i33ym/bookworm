package format

import (
	"encoding/json"
	"net/http"
)

func Respond(response http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		response.Header()[key] = value
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)

	_, err = response.Write(js)
	return err
}

func Read(request *http.Request, addr interface{}) error {
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(addr)
}
