package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	dataSlice := strings.Split(datastring, ",")

	if len(dataSlice) != 3 {
		return fmt.Errorf("reсeived slice does not contain 3 items")
	}

	steps, err := strconv.Atoi(dataSlice[0])

	if err != nil {
		return fmt.Errorf("cannot parce steps count: %w", err)
	}

	duration, err := time.ParseDuration(dataSlice[2])

	if err != nil {
		return fmt.Errorf("cannot parce time duration: %w", err)
	}

	if steps <= 0 || duration <= 0 {
		return fmt.Errorf("steps count (%d) or wall duration (%d) cannot will be zero or negative value: %w", steps, duration, err)
	}

	t.Steps = steps
	t.TrainingType = dataSlice[1]
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {

	steps := t.Steps
	weight := t.Personal.Weight
	height := t.Personal.Height
	duration := t.Duration
	activity := t.TrainingType

	message := "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n"

	distance := spentenergy.Distance(steps, height)
	speed := spentenergy.MeanSpeed(steps, height, duration)

	switch strings.ToLower(activity) {
	case "ходьба":
		calories, err := spentenergy.WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", fmt.Errorf("failed calc calories: %w", err)
		}

		return fmt.Sprintf(message, activity, duration.Hours(), distance, speed, calories), nil

	case "бег":
		calories, err := spentenergy.RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", fmt.Errorf("failed calc calories: %w", err)
		}

		return fmt.Sprintf(message, activity, duration.Hours(), distance, speed, calories), nil

	default:
		return "", fmt.Errorf("unknown activity kind\n(неизвестный тип тренировки):\n%s", activity)
	}
}
