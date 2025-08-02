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

	if steps <= 0 {
		return fmt.Errorf("steps (%d) must be greater than zero", steps)
	}

	if duration <= 0 {
		return fmt.Errorf("duration (%s) must be greater than zero", duration)
	}

	t.Steps = steps
	t.TrainingType = dataSlice[1]
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	message := "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n"

	switch strings.ToLower(t.TrainingType) {
	case "ходьба":
		calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("failed calc calories: %w", err)
		}

		return fmt.Sprintf(message, t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil

	case "бег":
		calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("failed calc calories: %w", err)
		}

		return fmt.Sprintf(message, t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil

	default:
		return "", fmt.Errorf("unknown activity kind\n(неизвестный тип тренировки):\n%s", t.TrainingType)
	}
}
