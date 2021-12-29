package business

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sachinmahanin/commonpassword/config"
	"github.com/sachinmahanin/commonpassword/handler/business/model"

	webserver "github.com/zhongjie-cai/web-server"
)

func SearchPassword(session webserver.Session) (interface{}, error) {

	session.LogMethodLogic(
		webserver.LogLevelInfo,
		"business",
		"Strength", "WELCOME",
	)
	var passwordStrengthRequest model.PasswordRequest
	var bodyError = session.GetRequestBody(
		&passwordStrengthRequest,
	)
	if bodyError != nil {
		return nil, bodyError
	}
	//load the file
	passwordFile, err := os.Open(config.PassswordFilePath)

	if err != nil {
		return nil, err
	}
	defer passwordFile.Close()
	// read the file word by word using scanner
	scanner := bufio.NewScanner(passwordFile)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		// If password is present in the list then return
		if scanner.Text() == passwordStrengthRequest.Password {
			return "Your password is present in the common password list ", nil
		}
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return "Your password is not in the common password list ", nil
}
