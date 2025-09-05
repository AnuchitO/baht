package main

import (
	"fmt"

	"github.com/anuchito/bahttext"
)

func main() {
	fmt.Println(bahttext.Words(1234.56))                  // หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์
	fmt.Println(bahttext.Words(-100))                     // ลบหนึ่งร้อยบาทถ้วน
	fmt.Println(bahttext.WordsFromString("1,000,000.01")) // หนึ่งล้านบาทหนึ่งสตางค์

	fmt.Println(bahttext.Words(1_000_000.01))     // หนึ่งล้านบาทหนึ่งสตางค์
	fmt.Println(bahttext.Words(1_000_000_000.01)) // หนึ่งล้านบาทหนึ่งสตางค์
}
