package main

import (
	"fmt"
	"strconv"
)

func convert(){
	PI := "3.1415926"
	f, _ := strconv.ParseInt(PI,10, 64);

	fmt.Println(f);

}