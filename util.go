package main

import (
	"log"
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func CmdHandler(req string) (res *linebot.FlexMessage) {
	var info []byte
	req = strings.ToLower(req)
	req = strings.TrimSpace(req)

	switch req {
	case "help":
		info = UsageInfo
	case "about":
		info = AboutInfo
	case "skill":
		info = SkillInfo
	case "project":
		info = ProjectInfo
	case "experience":
		info = ExperinceInfo
	default:
		info = UsageInfo
	}

	container, err := linebot.UnmarshalFlexMessageJSON(info)
	if err != nil {
		log.Print(err)
	}

	return linebot.NewFlexMessage(req, container)
}
