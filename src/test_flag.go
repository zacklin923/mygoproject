package main

import (
	"fmt"
	"flag"
)


func main(){
	var ip = flag.Int("flagname1",1234,"help message for flagname")
	//绑定到变量它们是值
	var flagvar int
	flag.IntVar(&flagvar, "flagname2", 1234, "help message for flagname" )
//	flag.Var(&flagVal, "name", "help message for flagname")	
	flag.Parse()

	fmt.Println("ip has value", *ip)
	fmt.Println("flagvar has value", flagvar)

}