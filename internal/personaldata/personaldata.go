package personaldata

import (
	"fmt"
)

type Personal struct {
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	message := "Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n"
	fmt.Printf(message, p.Name, p.Weight, p.Height)
}
