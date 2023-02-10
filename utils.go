package main

import (
	"log"
	"strconv"
)

func Str2int(text string) int{
	result, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal("Failed to parse string to int > ", err)
	}
	return result
}