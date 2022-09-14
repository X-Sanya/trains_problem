package train

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type Train struct {
	StartTime, EndTime time.Time
	Price              float32
	Number, From, To   int
}

func NewTrain(data []string) (Train, error) {
	if len(data) < 6 {
		return Train{}, errors.New("train: not enough data")
	}
	number, err := strconv.Atoi(strings.TrimSpace(data[0]))
	if err != nil {
		return Train{}, err
	}
	from, err := strconv.Atoi(strings.TrimSpace(data[1]))
	if err != nil {
		return Train{}, err
	}
	to, err := strconv.Atoi(strings.TrimSpace(data[2]))
	if err != nil {
		return Train{}, err
	}
	price, err := strconv.ParseFloat(strings.TrimSpace(data[3]), 32)
	if err != nil {
		return Train{}, err
	}
	startTime, err := time.Parse("15:04:05", strings.TrimSpace(data[4]))
	if err != nil {
		return Train{}, err
	}
	endTime, err := time.Parse("15:04:05", strings.TrimSpace(data[5]))
	if err != nil {
		return Train{}, err
	}
	return Train{Number: number, From: from, To: to, Price: float32(price), StartTime: startTime, EndTime: endTime}, nil
}

func (t *Train) String() string {
	return ("Train #" + strconv.Itoa(t.Number) +
		" from " + strconv.Itoa(t.From) + " to " + strconv.Itoa(t.To) +
		"; price: " + strconv.FormatFloat(float64(t.Price), 'f', 2, 32) +
		"; start at: " + t.StartTime.Format("15:04:05") +
		"; end at: " + t.EndTime.Format("15:04:05"))
}
