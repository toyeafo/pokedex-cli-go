package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Results []struct {
// 	Name string `json:"name"`
// 	URL  string `json:"url"`
// }

// func fillCache(result Results) []byte {
// 	// data := Results{}
// 	jsonbytes, err := json.Marshal(result)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return jsonbytes
// }

// func unfillCache(result []byte) Results {
// 	data := Results{}
// 	err := json.Unmarshal(result, &data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return data
// }
