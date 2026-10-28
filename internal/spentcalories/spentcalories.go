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
	// TODO: реализовать функцию
	splt := strings.Split(data, ",")
	if len(splt) != 3 {
		return 0, "", 0, fmt.Errorf("Количество запятых не равно 2\n")
	}
	stps, err := strconv.Atoi(splt[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("Произошла ошибка преобразования: %v\n", err)
	}
	if stps < 1 {
		return 0, "", 0, fmt.Errorf("Кол-во шагов должно быть положительным\n")
	}
	drtn, err := time.ParseDuration(splt[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("Произошла ошибка преобразования: %v\n", err)
	}
	if drtn <= 0 {
		return 0, "", 0, fmt.Errorf("Длительность должна быть больше нуля\n")
	}
	return stps, splt[1], drtn, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	oneStep := height * stepLengthCoefficient
	dstnc := float64(steps) * oneStep / float64(mInKm)
	return dstnc
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration > 0 {
		distance := distance(steps, height)
		hours := duration.Hours()
		return distance / hours
	}
	return 0
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, training, duration, err := parseTraining(data)
	if err != nil {
		log.Print(err)
	}
	speed := meanSpeed(steps, height, duration)
	distance := distance(steps, height)
	switch training {
	case "Бег":
		cals, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "Ошибка при расчете калорий при беге", err
		}
		duration := duration.Hours()
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", training, duration, distance, speed, cals), nil
	case "Ходьба":
		cals, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "Ошибка при расчете калорий при ходьбе]", err
		}
		duration := duration.Hours()
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", training, duration, distance, speed, cals), nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Неверное значение\n")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Неверное значение\n")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	cals := weight * meanSpeed * durationInMinutes / minInH * walkingCaloriesCoefficient
	return cals, nil
}
