package main

import (
	"fmt"
	"github.com/demian9k/mygomodule2sub1"
	"mygomodule1" //
)

func main() {
	msg1 := mygomodule1.GoModule1Fn()
	msg2 := mygomodule2sub1.GoModule2sub1Fn()

	fmt.Println(msg1, msg2)

}
