package main

import "popcount"


func main(){
	for i := 0; i < 5; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}
