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
		if err_parse != nil {
			log.Println(err_parse)
			fmt.Println(err_parse)
			continue
		}
		s, err_action = dp.ActionInfo()
		if err_action != nil {
			log.Println(err_action)
			fmt.Println(err_action)
			continue
		}
		fmt.Println(s)
	}
}
