package model

import (
	"fmt"
	"math/rand"
	"regexp"
)

func GeneratePassword(n int) string {

	const charactersAvailable string = "abcdefghijklmnopqrstuvwxyz#$%&_ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var pass string

	regex_az, _ := regexp.Compile("[a-z]")
	regex_AZ, _ := regexp.Compile("[A-Z]")
	regex_09, _ := regexp.Compile("[0-9]")
	regex_special, err := regexp.Compile("[#$%&_]")
	if err != nil {
		fmt.Println("Regex Compilation error: ", err)
	}

	for !regex_az.MatchString(pass) || !regex_AZ.MatchString(pass) || !regex_09.MatchString(pass) || !regex_special.MatchString(pass) {
		pass = ""
		for i := 0; i < n; i++ {
			randInt := rand.Intn(len(charactersAvailable))
			pass += string(charactersAvailable[randInt])
		}
	}

	return pass
}
