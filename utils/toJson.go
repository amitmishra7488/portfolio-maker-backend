package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

func ToJson(data ...any) {
	for _, v := range data {
		fmt.Println(string(jsonByte(v)))
	}
}

func jsonByte(v any) []byte {
	res, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Println(err)
		return nil
	}
	return res
}
