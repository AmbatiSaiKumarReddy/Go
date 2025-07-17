package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const url="https://www.google.com"

func main(){
	fmt.Println("LCO web Request")

	response,err:=http.Get(url)

	if(err!=nil){
		panic(err)

	}

	fmt.Printf("Response is of type: %T\n",response)


	defer response.Body.Close()  //callers responsibility to close the connection



	databytes,err	:=ioutil.ReadAll(response.Body)


	if(err!=nil){
		panic(err)
	}

	content:=string(databytes)
	fmt.Println(content)


}