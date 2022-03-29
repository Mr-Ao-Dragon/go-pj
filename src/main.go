package main

import (
	"fmt"
)

//import "github.com/beego/beego/v2/server/web"
func main() {
	sum := 0
loop:
	for i := 0; i <= 2; i++ {
		sum += i
	}
	fmt.Println(sum)
	goto loop
}
