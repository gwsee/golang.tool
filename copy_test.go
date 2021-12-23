package tool

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCopyFields(t *testing.T) {
	var z0 ZZ
	var z1 ZZ
	z1.Remarks = "dddd"
	z1.DD.LogisticsNumber = "ddddddddddddddddd"
	z1.DdD.LogisticsNumber = "xxxxxxxxxx"
	var dds []DD
	dds = append(dds, DD{LogisticsNumber: "LLLLLLLLLLLLLL"})
	var d DD
	d.LogisticsNumber = "2222222222222"
	z1.DdD2 = &d
	z1.DDs = dds
	CopyFields(&z0, &z1)
	r, _ := json.Marshal(z0)
	fmt.Println(string(r))
}
