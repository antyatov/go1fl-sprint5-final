package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dataSlice := strings.Split(datastring, ",")

	if len(dataSlice) != 2 {
		return fmt.Errorf("reсeived slice does not contain 2 items")
	}

	steps, err := strconv.Atoi(dataSlice[0])

	if err != nil {
		return fmt.Errorf("cannot parce steps count: %w", err)
	}

	duration, err := time.ParseDuration(dataSlice[1])

	if err != nil {
		return fmt.Errorf("cannot parce time duration: %w", err)
	}

	if steps <= 0 {
		return fmt.Errorf("steps (%d) must be greater than zero", steps)
	}

	if duration <= 0 {
		return fmt.Errorf("duration (%s) must be greater than zero", duration)
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)

	if err != nil {
		return "", fmt.Errorf("failed calc calories: %w", err)
	}

	message := "Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n"

	return fmt.Sprintf(message, ds.Steps, distance, calories), nil
}
