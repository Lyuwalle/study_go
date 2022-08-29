package main

import "fmt"

func hello()  {
	var stockcode = 100
	var enddate = "2021-08-20"
	var url = "Code=%d&endDate=%s"
	var target_url=fmt.Sprintf(url,stockcode,enddate)
	fmt.Println(target_url)
	fmt.Println("こんちは")
}