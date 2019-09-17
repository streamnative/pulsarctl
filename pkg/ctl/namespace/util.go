package namespace

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

func validateSizeString(s string) (int64, error) {
	end := s[len(s)-1:]
	value := s[:len(s)-1]
	switch end {
	case "k":
		fallthrough
	case "K":
		v, err := strconv.ParseInt(value, 10, 64)
		return v * 1024, err
	case "m":
		fallthrough
	case "M":
		v, err := strconv.ParseInt(value, 10, 64)
		return v * 1024 * 1024, err
	case "g":
		fallthrough
	case "G":
		v, err := strconv.ParseInt(value, 10, 64)
		return v * 1024 * 1024 * 1024, err
	case "t":
		fallthrough
	case "T":
		v, err := strconv.ParseInt(value, 10, 64)
		return v * 1024 * 1024 * 1024 * 1024, err
	default:
		return strconv.ParseInt(s, 10, 64)
	}
}

func parseRelativeTimeInSeconds(relativeTime string) (time.Duration, error) {
	if relativeTime == "" {
		return -1, errors.New("Time can not be empty.")
	}

	unitTime := relativeTime[len(relativeTime)-1:]
	t := relativeTime[:len(relativeTime)-1]
	timeValue, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return -1, errors.Errorf("Invalid time '%s'", t)
	}

	switch strings.ToLower(unitTime) {
	case "s":
		return time.Duration(timeValue) * time.Second, nil
	case "m":
		return time.Duration(timeValue) * time.Minute, nil
	case "h":
		return time.Duration(timeValue) * time.Hour, nil
	case "d":
		return time.Duration(timeValue) * time.Hour * 24, nil
	case "w":
		return time.Duration(timeValue) * time.Hour * 24 * 7, nil
	case "y":
		return time.Duration(timeValue) * time.Hour * 24 * 7 * 365, nil
	default:
		return -1, errors.Errorf("Invalid time unit '%s'", unitTime)
	}
}
