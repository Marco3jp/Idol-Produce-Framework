package main

import "fmt"

func errorCheck(err *error){
	if err != nil {
		fmt.Println("error:", err)
	}
}