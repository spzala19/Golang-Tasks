package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func scanstr(message string) string {
	fmt.Printf(message)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
func main() {
	myMap := make(map[string]string)
	myMap["name"] = scanstr("Enter Your Name :")
	myMap["address"] = scanstr("Enter Your Address :")
	obj, err := json.Marshal(myMap)
	if err == nil {
		fmt.Printf("Json Object : " + string(obj) + "\n")
	} else {
		fmt.Print(err)
	}
	//var objj map[string]interface{}
	//json.Unmarshal([]byte(string(obj)), &objj)
	//fmt.Print(objj)
}
