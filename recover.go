package common_tools

import (
	"fmt"
	"log"
)

func cover() (err error) {
	if errs := recover(); errs != nil {
		log.Println(errs)
		err = fmt.Errorf("%v", errs)
	}
	return
}
