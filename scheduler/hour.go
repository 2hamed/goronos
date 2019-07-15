package scheduler

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Hour represents an Hour in a 24H format
type Hour interface {
	Hour() int
	Minute() int
	IsAfter(t *time.Time) bool
	IsBefore(t *time.Time) bool

	String() string
}

// NewHour creates an Hour
func NewHour(h, m int) (Hour, error) {
	if h > 23 || h < 0 {
		return nil, errors.New("h can not be less than 0 or more than 23")
	}
	if m > 59 || m < 0 {
		return nil, errors.New("m can not be less than 0 or more than 59")
	}
	return hour{
		hour:   h,
		minute: m,
	}, nil
}

type hour struct {
	hour   int
	minute int
}

func (h hour) String() string {
	return fmt.Sprintf("%d:%02d", h.hour, h.minute)
}

func (h hour) IsAfter(t *time.Time) bool {
	return h.hour >= t.Hour() && h.minute >= t.Minute()
}

func (h hour) IsBefore(t *time.Time) bool {
	return h.hour <= t.Hour() && h.minute <= t.Minute()
}

func (h hour) Hour() int {
	return h.hour
}

func (h hour) Minute() int {
	return h.minute
}

// parseHour parses a string in the format of HH:mm into an `hour` struct
func parseHour(str string) (Hour, error) {
	hm := strings.Split(str, ":")
	if len(hm) == 0 {
		return nil, errors.New("invalid hour supplied")
	}
	h, err := strconv.Atoi(hm[0])

	if err != nil {
		return nil, errors.Wrap(err, "failed parsing the hour")
	}
	var m int
	if len(hm) > 1 {
		m, err = strconv.Atoi(hm[1])
		if err != nil {
			return nil, errors.Wrap(err, "failed parsing the minute")
		}
	}

	return NewHour(h, m)
}