package daysteps

import (
	"errors"
	"fmt"
	"log"
	"spentcalories"
	"strconv"
	"strings"
	"time"
)

const (
	stepLength                 = 0.65 // Длина одного шага в метрах
	mInKm                      = 1000 // Количество метров в одном километре
	minInH                     = 60   // минут в часе
	stepLengthCoefficient      = 0.45 // коэффициент длины шага от роста
	walkingCaloriesCoefficient = 0.5  // поправочный коэффициент для ходьбы
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("invalid data format")
	}

	stepsStr := parts[0]
	durStr := parts[1]

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps: %w", err)
	}
	if steps <= 0 {
		return 0, 0, errors.New("invalid steps: must be > 0")
	}

	dur, err := time.ParseDuration(durStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid duration: %w", err)
	}
	if dur <= 0 {
		return 0, 0, errors.New("invalid duration: must be > 0")
	}

	return steps, dur, nil
}

func distance(steps int, height float64) float64 {
	stepLen := height * stepLengthCoefficient
	return float64(steps) * stepLen / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height) // км
	return dist / duration.Hours()  // км/ч
}

func walkingCalories(steps int, weight, height float64, dur time.Duration) (float64, error) {
	cals, err := spentcalories.WalkingSpentCalories(steps, weight, height, dur)
	if err != nil {
		return 0, fmt.Errorf("walkingCalories error: %w", err)
	}
	return cals, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, dur, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	distanceKm := (float64(steps) * stepLength) / mInKm

	cals, err := walkingCalories(steps, weight, height, dur)
	if err != nil {
		log.Println(err)
		return ""
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distanceKm, cals,
	)
}
