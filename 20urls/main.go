package main

import (
	"fmt"
	"net/url"
)

const myurl string="https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghp_1234567890"

func main(){

	fmt.Println("Welcome to Handling Urls in Golang")
	fmt.Println(myurl)


	result,_:=url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams:=result.Query()

	fmt.Printf("The type of query params are: %T\n",qparams)

	fmt.Println(qparams["coursename"])

	partsOfUrl:=&url.URL{
		Scheme:"https",
		Host:"lco.dev",
		Path:"/tutcss",
		RawQuery:"user=hitesh",
	}

	newUrl:=result.ResolveReference(partsOfUrl)
	fmt.Println(newUrl.String())

	anotherUrl:=partsOfUrl.String()
	fmt.Println(anotherUrl)

	
	
	





}