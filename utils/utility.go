package utils

import (
	"errors"
	"strings"
	"time"
)

func TimeBetweenParse(mulai string, sampai string) (tMulai time.Time, tSampai time.Time, err error) {
	tMulai, err = time.Parse("2006-1-2", strings.Trim(mulai, " "))
	if err != nil {
		return
	}
	tSampai, err = time.Parse("2006-1-2", strings.Trim(sampai, " "))
	if err != nil {
		return
	}

	if tMulai.Before(tSampai) {
		err = errors.New("waktu mulai lebih lama dari waktu sampai")
	}
	return
}
