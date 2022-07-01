package v1

import (
	"encoding/json"
	"net/http"
)

func SendJSONResponse(writer http.ResponseWriter, data interface{}) error {
	jsonResponse, err := json.Marshal(&data)

	if err != nil {
		return err
	}

	writer.Header().Set("Content-Type", "application/json")

	_, err = writer.Write(jsonResponse)
	if err != nil {
		return err
	}
	return nil
}
