package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var (
	sha = flag.Int("s", 256, "使用するSHA")
)

func String2SHA256(s string) [32]uint8 { return sha256.Sum256([]byte(s)) }

func String2SHA384(s string) [48]uint8 { return sha512.Sum384([]byte(s)) }

func String2SHA512(s string) [64]uint8 { return sha512.Sum512([]byte(s)) }

func main() {
	flag.Parse()
	var s string
	fmt.Scan(&s)
	switch opt := *sha; opt {
	case 256:
		fmt.Printf("%x\n", String2SHA256(s))
	case 384:
		fmt.Printf("%x\n", String2SHA384(s))
	case 512:
		fmt.Printf("%x\n", String2SHA512(s))
	}
}
