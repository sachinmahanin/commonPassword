package main

import (
	"io"
	"net/http"
	"os"

	"github.com/sachinmahanin/commonpassword/config"
	customization "github.com/sachinmahanin/commonpassword/customization"
)

// This is a sample of how to setup application for running the server
func main() {
	var configError = configSetupApplication()
	if configError != nil {
		panic(
			configError,
		)
	}
	var port, convErr = strconvAtoi(config.AppPort)
	if convErr != nil {
		panic(
			fmtErrorf(
				"Invalid port number provided: %v",
				config.AppPort,
			),
		)
	}
	//Download file
	err := DownloadFileFunc(config.PassswordFilePath, config.CommonPasswordListURL)
	if err != nil {
		panic(
			fmtErrorf(
				"Not able to download the file from url: %v",
				config.CommonPasswordListURL,
			),
		)
	}
	var application = webserverNewApplication(
		config.AppName,
		port,
		config.AppVersion,
		&customization.Customization{},
	)
	defer application.Stop()
	application.Start()
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
