package main

import (
	"errors"
	"math/rand"
	"time"
)

type Storage interface {
	isBannerRegistered(bannerId, slotId int) (bool, error)
	addBannerPlacement(bannerId, slotId int) error
	removeBannerPlacement(bannerId, slotId int) error
	addClick(bannerId, slotId, socialGroupId int) error
	addDisplay(bannerId, slotId, socialGroupId int) error
	getStats(slotId, socialGroupId int) ([]Stat, error)
}

func addBannerPlacement(storage Storage, bannerId, slotId int) error {
	return storage.addBannerPlacement(bannerId, slotId)
}

func removeBannerPlacement(storage Storage, bannerId, slotId int) error {
	return storage.removeBannerPlacement(bannerId, slotId)
}

func addClick(storage Storage, bannerId, slotId, socialGroup int) error {
	registered, err := storage.isBannerRegistered(bannerId, slotId)
	if err != nil {
		return err
	}

	if !registered {
		return errors.New("banner placement wasn't found")
	}

	err = storage.addClick(bannerId, slotId, socialGroup)
	if err != nil {
		return err
	}

	return nil
}

func getBannerToDisplay(storage Storage, slotId, socialGroupId int) (int, error) {
	stats, err := storage.getStats(slotId, socialGroupId)
	if err != nil {
		return 0, err
	}

	maxIndex := 0
	for i, _ := range stats {
		if stats[i].prob > stats[maxIndex].prob {
			maxIndex = i
		}
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)

	if n > 70 {
		for {
			i := rand.Intn(len(stats))
			if i != maxIndex {
				maxIndex = i
				break
			}
		}
	}

	bannerId := stats[maxIndex].bannerId
	err = storage.addDisplay(bannerId, slotId, socialGroupId)
	if err != nil {
		return 0, err
	}
	return bannerId, nil
}

func main() {}