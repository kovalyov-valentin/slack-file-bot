package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3755665871239-3772058390547-93pvWbreEd6ErbI0MS6JJMRG")
	os.Setenv("CHANNEL_ID", "C03N7KKSLNT")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"Golang_dlya_profi_rabota_s_setyu_mnogopotochnost_struktury_dannykh_i_mashinnoe_obuchenie_s_Go_2020_Tsukalos.pdf"}

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
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}
