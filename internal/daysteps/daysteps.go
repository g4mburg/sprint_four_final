package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parseData := strings.Split(data, ",")
	if len(parseData) != 2 {
		return 0, 0, fmt.Errorf("there are not enough parameters")
	}
	stepCount, err := strconv.Atoi(parseData[0])
	if err != nil {
		return 0, 0, fmt.Errorf("error in number of steps")
	} else if stepCount == 0 {
		return 0, 0, fmt.Errorf("the number of steps is zero")
	}
	walkDuration, err := time.ParseDuration(parseData[1])
	if err != nil {
		return 0, 0, fmt.Errorf("error in the duration of the walk")
	} else if walkDuration == 0 {
		return 0, 0, fmt.Errorf("the walking time is zero")
	}
	return stepCount, walkDuration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	stepCount, walkDuration, err := parsePackage(data)
	if err != nil {
		fmt.Println("something went wrong", err)
		return ""
	} else if stepCount == 0 {
		fmt.Println("something went wrong", err)
		return ""
	}
	distance := float64(stepCount) * stepLength
	distance /= mInKm
	calories, err := spentcalories.WalkingSpentCalories(stepCount, weight, height, walkDuration)
	if err != nil {
		fmt.Println("something went wrong", err)
	}
	return (fmt.Sprintf("Количество шагов: %d.\nДистанция составила %f км.\nВы сожгли %f ккал.", stepCount, distance, calories))
}
