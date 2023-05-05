package main

import (
	"fmt"

	"github.com/SajjanYadav/cryptit/decrypt"

	"github.com/SajjanYadav/cryptit/encrypt"
)

func main(){
	fmt.Println(encrypt.Nimbus("p^gg^k"))
	fmt.Println(decrypt.Nimbus("p^gg^k"))
}