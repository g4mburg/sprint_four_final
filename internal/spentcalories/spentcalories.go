package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	parseData := strings.Split(data, ",")
	if len(parseData) != 3 {
		return 0, "", time.Duration(0), fmt.Errorf("there are not enough parameters")
	}
	stepCount, err := strconv.Atoi(parseData[0])
	if err != nil {
		return 0, "", time.Duration(0), err
	} else if stepCount <= 0 {
		return 0, "", time.Duration(0), fmt.Errorf("the number of steps is zero")
	}

	activityType := parseData[1]

	activityDuration, err := time.ParseDuration(parseData[2])
	if err != nil {
		return 0, "", time.Duration(0), err
	} else if activityDuration <= 0 {
		return 0, "", time.Duration(0), fmt.Errorf("the duration of the activity is zero")
	}

	return stepCount, activityType, activityDuration, nil
}

func distance(steps int, height float64) float64 {
	strideLength := height * stepLengthCoefficient
	distance := float64(steps) * strideLength
	distance /= mInKm
	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	distance := distance(steps, height)

	meanSpeed := distance / duration.Hours()

	return meanSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	stepCount, activityType, activityDuration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
	}

	trainingInfo := ""

	switch activityType {
	case "Ходьба":
		distance := distance(stepCount, height)
		meanSpeed := meanSpeed(stepCount, height, activityDuration)
		calories, err := WalkingSpentCalories(stepCount, weight, height, activityDuration)
		if err != nil {
			log.Println(err)
		}

		trainingInfo = fmt.Sprintf("Тип тренировки: %s\nДлительность: %s ч.\nДистанция: %f км.\nСкорость: %f км/ч\nСожгли калорий: %f ", activityType, activityDuration, distance, meanSpeed, calories)

		return trainingInfo, nil

	case "Бег":
		distance := distance(stepCount, height)
		meanSpeed := meanSpeed(stepCount, height, activityDuration)
		calories, err := RunningSpentCalories(stepCount, weight, height, activityDuration)
		if err != nil {
			log.Println(err)
		}

		trainingInfo = fmt.Sprintf("Тип тренировки: %s\nДлительность: %s ч.\nДистанция: %f км.\nСкорость: %f км/ч\nСожгли калорий: %f ", activityType, activityDuration, distance, meanSpeed, calories)

		return trainingInfo, nil

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("the number of steps is zero")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight is zero")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height is zero")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("the duration of the running is zero")
	}

	meanSpeed := meanSpeed(steps, height, duration)

	duration = time.Duration(duration.Minutes())

	calories := (weight * meanSpeed * float64(duration)) / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("the number of steps is zero")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight is zero")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height is zero")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("the duration of the running is zero")
	}

	meanSpeed := meanSpeed(steps, height, duration)

	duration = time.Duration(duration.Minutes())

	calories := ((weight * meanSpeed * float64(duration)) / minInH) * walkingCaloriesCoefficient

	return calories, nil
}
