package daysteps

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

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	info := strings.Split(datastring, ",")
	if len(info) != 2 {
		return errors.New("count element not aqual 2")
	}
	num, err := strconv.Atoi(info[0])
	if err != nil {
		return err
	}
	if num <= 0 {
		return errors.New("count steps under or aqual 0")
	}
	ds.Steps = num
	t, err := time.ParseDuration(info[1])
	if err != nil {
		return err
	}
	if t <= 0 {
		return errors.New("duration under or aqual 0")
	}
	ds.Duration = t
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	result, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		log.Println(err)
		return "", err
	}
	message := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, result)
	return message, nil
}
