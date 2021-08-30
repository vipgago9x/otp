package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vipgago9x/otp/sdk/models"
)

// Call api
func sendSMSRequest(token string, req_id string, phone_number string, template string) (*models.ResponseModel, error) {

	// Make body
	body := fmt.Sprintf(`{"request_id": "%s","phone_number": "%s","template": "%s"}`, req_id, phone_number, template)

	// Define http client
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://benthuba:8080/r/otp/send", bytes.NewBuffer([]byte(body)))

	if err != nil {
		return nil, err
	}

	// Set content-type and authorization header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	// Excute request
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	// Response body Raw
	resBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resBody models.ResponseModel
	// Convert Raw body to ResponseModel
	err = json.Unmarshal(resBodyBytes, &resBody)
	if err != nil {
		return nil, err
	}

	return &resBody, nil

}
