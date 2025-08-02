package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, str := range dataset {
		err := dp.Parse(str)
		if err != nil {
			log.Println("Skip.. Failed parse dataset:", err)
			continue
		}

		message, err := dp.ActionInfo()
		if err != nil {
			log.Println("failed get action info: %w", err)
		}

		fmt.Println(message)
	}
}
