package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {

	//Разобранная на элементы строка
	walking := strings.Split(datastring, ",")

	if len(walking) != 2 {
		ds.Steps = 0
		return errors.New("ошибка: входящая строка не соответствует требованию о наличии 2 позиций")
	}

	// Количество шагов, преобразование.
	stepsQuantity, err := strconv.Atoi(walking[0])
	if err != nil {
		ds.Steps = 0
		return err
	}

	if stepsQuantity <= 0 {
		ds.Steps = 0
		return errors.New("ошибка, значение количества шагов равно или менее нуля!")
	}

	ds.Steps = stepsQuantity

	// Длительность активности (прогулки).
	walkDuration, err := time.ParseDuration(walking[1])
	if err != nil {
		ds.Duration = 0
		return err
	}

	if walkDuration <= 0 {
		ds.Duration = 0
		return errors.New("ошибка, значение длительности прогулки равно или менее нуля!")
	}
	ds.Duration = walkDuration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// Пройденное расстояние.
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	// Сожжённые калории.
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, spentCalories)
	return result, err
}
