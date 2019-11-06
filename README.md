
# Simple Ip Address Generator from subnet

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)


```go
package main

import (
	"fmt"
	ipgen "github.com/wahyuhadi/go-ipgen"

)


func main(){

	ip := ipgen.IpAddressGen("10.30.0.0/16")

	for i:=0; i<len(ip); i++ {
		fmt.Println(ip[i])
	}

	fmt.Println(len(ip))

}
```