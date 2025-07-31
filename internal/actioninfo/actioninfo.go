package actioninfo

import "log"

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

		dp.ActionInfo()
	}
}
