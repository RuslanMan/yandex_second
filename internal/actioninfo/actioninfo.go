package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	var err_parse, err_action error
	var s string
	for _, value := range dataset {
		err_parse = dp.Parse(value)
		log.Println(err_parse)
		s, err_action = dp.ActionInfo()
		log.Println(err_action)
		fmt.Println(s)
	}
}
