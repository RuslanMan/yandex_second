package trainings

import (
	"errors"
	"fmt"
	"log"
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
	info := strings.Split(datastring, ",")
	if len(info) != 3 {
		return errors.New("count element not aqual 3")
	}
	num, err := strconv.Atoi(info[0])
	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}
	if num <= 0 {
		return errors.New("count steps under or aqual 0")
	}
	t.Steps = num
	t.TrainingType = info[1]
	ti, err := time.ParseDuration(info[2])
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}
	if ti <= 0 {
		return errors.New("duration under or aqual 0")
	}
	t.Duration = ti
	return nil
}

func (t Training) ActionInfo() (string, error) {

	distanceInKm := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	message := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distanceInKm, speed, calories)
	return message, nil
}
