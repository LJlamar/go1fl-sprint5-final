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

// WalkingSpentCalories принимает количество шагов, вес в кг и рост в м пользователя, продолжительность ходьбы, возвращая потраченные калории либо ошибку.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if weight <= 0 {
		return 0, errors.New("ошибка, вес указан как отрицательный либо нулевой!")
	}
	if height <= 0 {
		return 0, errors.New("ошибка, рост указан как отрицательный либо нулевой!")
	}
	if duration <= 0 {
		return 0, errors.New("ошибка, длительность менее либо равна нулю!")
	}
	if steps <= 0 {
		return 0, errors.New("ошибка, количество шагов менее либо равно нулю")
	}
	// Средняя скорость
	avgSpeed := MeanSpeed(steps, height, duration)
	spentCaloriesWalking := (weight * avgSpeed * duration.Minutes()) / minInH * walkingCaloriesCoefficient
	return spentCaloriesWalking, nil
}

// RunningSpentCalories принимает количество шагов, вес в кг и рост в м пользователя, продолжительность бега, возвращая потраченные калории либо ошибку.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if weight <= 0 {
		return 0, errors.New("ошибка, вес указан как отрицательный либо нулевой!")
	}
	if height <= 0 {
		return 0, errors.New("ошибка, рост указан как отрицательный либо нулевой!")
	}
	if duration <= 0 {
		return 0, errors.New("ошибка, длительность менее либо равна нулю!")
	}
	if steps <= 0 {
		return 0, errors.New("ошибка, количество шагов менее либо равно нулю")
	}
	//Средняя скорость
	avgSpeed := MeanSpeed(steps, height, duration)
	spentCaloriesRunning := (weight * avgSpeed * duration.Minutes()) / minInH

	return spentCaloriesRunning, nil

}

// MeanSpeed принимает количество шагов, рост пользователя, продолжительность активности и возвращает среднюю скорость через distance().
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	// Средняя скорость
	avgSpeed := Distance(steps, height) / duration.Hours()
	return avgSpeed

}

// Distance принимает количество шагов и рост пользователя в метрах, просчитывает и возвращает дистанцию в км.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distance := float64(steps) * stepLength / mInKm
	return distance
}
