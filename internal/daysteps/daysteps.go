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
	// TODO: реализовать функцию
	splt := strings.Split(data, ",")
	if len(splt) != 2 {
		return 0, 0, fmt.Errorf("В строке отсутствуют запятые или их больше 1\n")
	}
	stps, err := strconv.Atoi(splt[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Произошла ошибка преобразования: %v\n", err)
	}
	if stps < 1 {
		return 0, 0, fmt.Errorf("Кол-во шагов должно быть положительным\n")
	}
	drtn, err := time.ParseDuration(splt[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Произошла ошибка преобразования: %v\n", err)
	}
	if drtn <= 0 {
		return 0, 0, fmt.Errorf("Продолжительность должна быть больше нуля\n")
	}
	return stps, drtn, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	stps, drtn, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if stps < 1 {
		return ""
	}
	dstns := float64(stps) * stepLength / mInKm
	ccals, err := spentcalories.WalkingSpentCalories(stps, weight, height, drtn)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", stps, dstns, ccals)
}
