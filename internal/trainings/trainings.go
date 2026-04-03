package trainings

import (
	"errors"
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
	//Разобранная на элементы строка
	trainingSummary := strings.Split(datastring, ",")

	if len(trainingSummary) != 3 {
		return errors.New("ошибка: входящая строка не соответствует требованию о наличии 3 позиций")
	}

	// Количество шагов.
	stepsQuantity, err := strconv.Atoi(trainingSummary[0])
	if err != nil {
		return err
	}

	if stepsQuantity <= 0 {
		return errors.New("ошибка, значение количества шагов равно или менее нуля!")
	}

	t.Steps = stepsQuantity

	// Тип тренировки.
	t.TrainingType = trainingSummary[1]

	// Длительность тренировки.
	activityDuration, err := time.ParseDuration(trainingSummary[2])
	if err != nil {
		//t.Duration = 0
		return err
	}
	//ТУТ Я МОГ ДОПУСТИТЬ ОШИБКУ
	if activityDuration <= 0 {
		return errors.New("ошибка, значение длительности тренировки равно или менее нуля!")
	}
	t.Duration = activityDuration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// Расстояние, преодоленное в ходе тренировки.
	distance := spentenergy.Distance(t.Steps, t.Height)

	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch t.TrainingType {
	case "Бег":
		spentCalories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, avgSpeed, spentCalories)
		return result, err
	case "Ходьба":
		spentCalories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, avgSpeed, spentCalories)
		return result, err

	default:
		return "", errors.New("неизвестный тип тренировки")
	}

}
