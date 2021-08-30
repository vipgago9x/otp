package main

import (
	"fmt"

	"github.com/vipgago9x/otp/sdk"
)

func main() {
	status, id, err := sdk.OtpRequest("dasdsadsa", "dasjdjsa", "039884321")
	fmt.Println(status)
	fmt.Println(id)
	fmt.Println(err)
}
