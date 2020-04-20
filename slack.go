package main


import (
	"fmt"

	"github.com/psychedelicnekopunch/go-sample/infrastructure"
)


func main() {
	slack := infrastructure.NewSlack()
	channel, timestamp, text, err := slack.SendMessage("test")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Print(channel + "\n")
	fmt.Print(timestamp + "\n")
	fmt.Print(text + "\n")
}
