package main

import (
	data "EncodeDecodeJWT/data"
	decode "EncodeDecodeJWT/decode"
	encode "EncodeDecodeJWT/encode"
	"fmt"
)

func main() {
	///? Encode
	token, err := encode.GenerateJWTToken(data.UserId)
	if err != nil {
		fmt.Println("err ", err)
	}
	fmt.Println(token)

	///? Decode
	id, issuer, err := decode.DecriptJWTTokenToId(token)
	if err != nil {
		fmt.Println("err ", err)
	}
	fmt.Printf("ID: %v, Issuer: %v", id, issuer)

}
