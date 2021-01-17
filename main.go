package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type jsondata struct { //create struct for body's json data
	Data    struct {
		Code   string `json:"code"`
	} `json:"data"`
}

func main() {
	// initialize how many code will be generated
	fmt.Print("Enter the number of invite codes: ")
	var num1 string
	fmt.Scanln(&num1)
	num, _ := strconv.Atoi(num1)
	i := 0
	for i < num {
		// make post req with null data
		req, _ := http.Post("https://www.hackthebox.eu/api/invite/generate", "x-www-form-urlencoded", nil)
		// close the body for performance
		defer req.Body.Close()
		//  read body
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err) // error control
		}

		codedata := jsondata{}                                 // initialize empty struct
		json.Unmarshal([]byte(body), &codedata)                // unmarshal json to the struct
		codebase64 := codedata.Data.Code                       // get base64 encoded code
		code, _ := base64.StdEncoding.DecodeString(codebase64) // decode code
		fmt.Printf("Code : %s", string(code))                  // print code
		i++
	}
}
