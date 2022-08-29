package main

import (
	"fmt"
	"time"
)

func main1() {
	go mapTest()
}


// Book 定义一个结构体
type Book struct {
	Title string
	Author string
	Price string
	Sn string
	PublishTime time.Time
}
func structTest() {
	var WarAndPeace Book
	var MingDinasty Book
	WarAndPeace.Title = "WarAndPeace"
	WarAndPeace.PublishTime = time.Now()
	WarAndPeace.Sn = "12345678"
	WarAndPeace.Price = "$5"
	WarAndPeace.Author = "Hu gor"
	MingDinasty.PublishTime = time.Now()
	fmt.Println(WarAndPeace.Title)
	fmt.Println(MingDinasty.PublishTime)
}

func mapTest() {
	var countryCapitalMap map[string]string
	countryCapitalMap [ "France" ] = "pairs"
	countryCapitalMap [ "China" ] = "shanghai"
	countryCapitalMap [ "Taiwan" ] = "taipei"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是：", countryCapitalMap[country])
	}
}
