package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://fcm.googleapis.com/fcm/send"

	payload := strings.NewReader("{\n \"to\" : \"/topics/Arqui1\",\n  \"notification\" : {\n \"body\" : \"El sensor de humo detecto algo\",\n \"title\" : \"Sensor de humo activado\"\n \"content_available\" : \"true\",\n \"priority\" : \"high\",\n }\n}\n")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", "key=AIzaSyCmvxnrqEFD5nwkH_n4RB-ItWLVFsYwCfI")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}