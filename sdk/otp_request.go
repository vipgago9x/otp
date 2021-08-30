package sdk

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/vipgago9x/otp/sdk/models"
)

// Otp service
func OtpRequest(token string, req_id string, phone_number string) (status int, id string, _err error) {

	var file []byte
	var err error
	var fileExist bool = true

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// Make folder containing file if not exist
	if _, err := os.Stat(dirname + `\otp-sdk`); os.IsNotExist(err) {
		err := os.Mkdir(dirname+`\otp-sdk`, 0664)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Try to open current date file
	file, err = ioutil.ReadFile(fmt.Sprintf(`%s\otp-sdk\%s.txt`, dirname, time.Now().Format("2006-01-02")))

	if err != nil {
		fileExist = false
		fmt.Println(err)
	}

	var fileData = []string{}
	if fileExist {
		fileData = append(fileData, strings.Split(string(file), ",")...)
		if len(fileData) > 500 {
			return -1, "", errors.New("The number of requests per day is more than 500.")
		}
		count := 0
		for i := 0; i < len(fileData); i++ {
			if fileData[i] == phone_number {
				count++
			}
		}
		if count >= 5 {
			return -1, "", errors.New("The number of requests per day for this phone number is more than 5.")
		}
	}

	// Random OTP
	rand.Seed(time.Now().UTC().UnixNano())
	optNum := rand.Intn(999999-100000) + 100000

	// Make OTP request template
	template := fmt.Sprintf(`Ma xac thuc OTP la: %d`, optNum)

	var res *models.ResponseModel

	// Retry 3 times
	for i := 0; i < 3; i++ {
		res, err = sendSMSRequest(token, req_id, phone_number, template)
		if err != nil {
			if i < 2 {
				continue
			} else {
				return -1, "", err
			}
		}
		break
	}

	// Check if error
	if res.Error.Code != 0 {
		return -1, "", errors.New(res.Error.Message)
	}

	fileData = append(fileData, phone_number)

	// Add this phone number to today file
	err = ioutil.WriteFile(fmt.Sprintf(`%s\otp-sdk\%s.txt`, dirname, time.Now().Format("2006-01-02")), []byte(strings.Join(fileData, ",")), 0644)

	if err != nil {
		fmt.Println(err)
	}
	return 1, res.Data.Id, nil
}
