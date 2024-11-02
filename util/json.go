package util

import (
	"encoding/json"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	"io"
	"os"
)

func BuildJson(s interface{}, filename string) {

	file, err := os.Create(fmt.Sprintf("json/%s.json", filename))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(&s); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}

func ReadJson(s interface{}, filename string) interface{} {
	file, err := os.Open(fmt.Sprintf("json/%s.json", filename))

	if err != nil {
		fmt.Println("build JSON Error : ", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&s)
	if err != nil && err != io.EOF {
		fmt.Println("error decoding JSON: ", err)
	}
	return s
}

func PrintJson(s interface{}, title ...string) {
	if len(title) != 0 {
		fmt.Println(gola.Tf(gola.Bold, title[0], gola.LightGreen))
	}
	indented, _ := json.MarshalIndent(s, "", " ")
	fmt.Println(string(indented))
}
