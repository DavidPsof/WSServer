package cmd

import (
	"WSServer/server"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Scan() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		handleInputText(scanner.Text())
	}
}

func handleInputText(text string) {
	// command parts
	cp := strings.Split(text, " ")

	switch cp[0] {
	case "sendTest":
		fmt.Println(server.SendMessage("test message", "test"))

	case "send":
		if len(cp) > 1 && strings.HasPrefix(cp[1], "--hub") {
			fmt.Println(server.SendMessage("Random", strings.TrimPrefix(cp[1], "--hub=")))
		}

	case "sendc":
		if len(cp) > 1 && strings.HasPrefix(cp[1], "--id") {
			fmt.Println(server.SendMessageToUser(strings.TrimPrefix(cp[1], "--id="), "test message"))
		}

	case "help":
		fmt.Println("send --hub - (где параметр --hub -номер hub) осуществляется broadcast рассылка произвольного сообщения всеклиентам указанного hub")

	default:
		fmt.Println("unknown command, you can use \"help\" to find out information about the supported commands")
	}
}
