package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func validateSpentParams(steps int, weight, height float64, duration time.Duration) error {
	if steps <= 0 {
		return fmt.Errorf("invalid steps value :%d", steps)
	}

	if weight <= 0 || height <= 0 {
		return fmt.Errorf("invalid weight(%f) or height (%f) value", weight, height)
	}

	if duration <= 0 {
		return fmt.Errorf("invalid diration time: %s", duration)
	}
	return nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	err := validateSpentParams(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	walkingSpeed := MeanSpeed(steps, height, duration)
	calories := weight * walkingSpeed * duration.Minutes()

	return (calories / minInH) * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	err := validateSpentParams(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	runningSpeed := MeanSpeed(steps, height, duration)
	calories := weight * runningSpeed * duration.Minutes()

	return calories / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	distanceOfMeters := (height * stepLengthCoefficient) * float64(steps)

	return distanceOfMeters / mInKm
}
