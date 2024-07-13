package service

import (
	"errors"
	"github.com/Bartosz-D3V/recruitment-task-go/config"
	"math"
)

type SearchSvc interface {
	BinarySearch(int) (int, error)
}

type searchSvc struct {
	config config.AppConfig
}

func New(config config.AppConfig) SearchSvc {
	return &searchSvc{
		config: config,
	}
}

func (svc *searchSvc) BinarySearch(number int) (int, error) {
	log := svc.config.Logger

	log.Infof("Searching for number %d", number)
	numbers := svc.config.Numbers
	l := 0
	r := len(*numbers) - 1
	var res *int = nil
	for l <= r {
		mid := l + (r-l)/2
		midNum := (*numbers)[mid]
		log.Debugf("Finding number %d. Left pointer: %d, Right pointer: %d, Middle pointer: %d", number, l, r, mid)
		if midNum == number {
			return mid, nil
		}

		if midNum > number {
			r = mid - 1
			if withinTenPercent(float64(midNum), float64(number)) {
				res = &mid
			}
		}
		if midNum < number {
			l = mid + 1
			if withinTenPercent(float64(midNum), float64(number)) {
				res = &mid
			}
		}
	}

	if res != nil {
		return *res, nil
	}

	log.Infof("Number %d not found", number)
	return -1, errors.New("number not found")
}

func withinTenPercent(a, b float64) bool {
	difference := math.Abs(a - b)
	tolerance := 0.1 * math.Abs(a)
	return difference <= tolerance
}
