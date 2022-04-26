package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("BOT_TOKEN", "xoxb-3412834228678-3438581438581-EymzI2DmsGD2FogzU1yx4Yw3")
	os.Setenv("CHANNEL_ID", "C03BWS3M18F")

	api := slack.New(os.Getenv("BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"Test_PDF.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("File name:%s and URL:%s\n", file.Name, file.URL)
	}
}
