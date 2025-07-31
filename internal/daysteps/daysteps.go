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

	if steps <= 0 || duration <= 0 {
		return fmt.Errorf("steps count (%d) or wall duration (%d) cannot will be zero or negative value: %w", steps, duration, err)
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	message := "Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n"
	steps := ds.Steps
	weight := ds.Personal.Weight
	height := ds.Personal.Height
	duration := ds.Duration

	distance := spentenergy.Distance(steps, height)

	calories, err := spentenergy.WalkingSpentCalories(steps, weight, height, duration)

	if err != nil {
		return "", fmt.Errorf("failed calc calories: %w", err)
	}

	return fmt.Sprintf(message, steps, distance, calories), nil
}
