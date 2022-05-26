package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"time"
)

var AllKeys [][]byte

func main() {
	All()

	num8, _ := GenerateRandomBytes(8)
	num16, _ := GenerateRandomBytes(16)
	num32, _ := GenerateRandomBytes(32)
	num64, _ := GenerateRandomBytes(64)
	num128, _ := GenerateRandomBytes(128)
	num256, _ := GenerateRandomBytes(256)
	num512, _ := GenerateRandomBytes(512)
	num1024, _ := GenerateRandomBytes(1024)
	num2048, _ := GenerateRandomBytes(2048)
	num4096, _ := GenerateRandomBytes(4096)

	fmt.Println(num8)
	fmt.Println(num16)
	fmt.Println(num32)
	fmt.Println(num64)
	fmt.Println(num128)
	fmt.Println(num256)
	fmt.Println(num512)
	fmt.Println(num1024)
	fmt.Println(num2048)
	fmt.Println(num4096)

	fmt.Println("Keys -- ", AllKeys)

	BootForce(num32)

}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	AllKeys = append(AllKeys, b)

	return b, nil
}

func BootForce(key []byte) {
	start := time.Now()
	duration := time.Since(start)
	for i, n := range AllKeys {
		if key[i] == n[i] {
			fmt.Println("It is your key", n)
			log.Print(duration)
		}
	}

}

func All() {
	Vint8 := math.Pow(2, 8)
	Vint16 := math.Pow(2, 16)
	Vint32 := math.Pow(2, 32)
	Vrint32 := int64(Vint32)
	Vint64 := math.Pow(2, 64)
	Vrint64 := uint64(Vint64)
	Vint128 := big.NewInt(9223372036854775807) // int64
	Vint128.Mul(Vint128, big.NewInt(9223372036854775807))
	Vint256 := Vint128
	Vint256.Mul(Vint256, Vint128)
	Vint512 := Vint256
	Vint512.Mul(Vint512, Vint256)
	Vint1024 := Vint512
	Vint1024.Mul(Vint1024, Vint512)
	Vint2048 := Vint1024
	Vint2048.Mul(Vint2048, Vint1024)
	Vint4096 := Vint2048
	Vint4096.Mul(Vint4096, Vint2048)

	fmt.Println(Vint8)
	fmt.Println(Vint16 - 1)
	fmt.Println(Vrint32 - 1)
	fmt.Println(Vrint64 - 1)
	fmt.Println(Vint128)
	fmt.Println(Vint256)
	fmt.Println(Vint512)
	fmt.Println(Vint1024)
	fmt.Println(Vint2048)
	fmt.Println(Vint4096)

}
