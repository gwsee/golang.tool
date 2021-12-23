package common_tools

import (
	"fmt"
	"log"
	"testing"
)

func TestIsSames(t *testing.T) {
	var a ZZ
	a.Remarks = "dd"
	a.LogisticsNumber = "num"
	var b ZZ
	b.Remarks = "d2d"
	b.LogisticsNumber = "num"
	f, m, e := IsSames(&a, &b, "json", 3)
	if e != nil {
		log.Println(e.Error())
		return
	}
	fmt.Println(f)
	for k, v := range m {
		fmt.Println(k, v)
	}
}
