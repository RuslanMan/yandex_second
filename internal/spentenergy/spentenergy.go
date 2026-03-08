package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	result, err := RunningSpentCalories(steps, weight, height, duration)
	return result * walkingCaloriesCoefficient, err
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("count steps under or aqual 0")
	}
	if weight <= 0 {
		return 0, errors.New("weight under or aqual 0")
	}
	if height <= 0 {
		return 0, errors.New("height under or aqual 0")
	}
	if duration <= 0 {
		return 0, errors.New("duration under or aqual 0")
	}
	avg_speed := MeanSpeed(steps, height, duration)
	return (float64(weight) * avg_speed * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	return float64(steps) * height * stepLengthCoefficient / mInKm
}
