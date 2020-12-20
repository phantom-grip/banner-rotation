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

type BannerRecomender struct {
	storage Storage
}

func NewBannerRecomender(storage Storage) *BannerRecomender {
	return &BannerRecomender{storage: storage}
}

func (bc BannerRecomender) addBannerPlacement(bannerId, slotId int) error {
	return bc.storage.addBannerPlacement(bannerId, slotId)
}

func (bc BannerRecomender) removeBannerPlacement(bannerId, slotId int) error {
	return bc.storage.removeBannerPlacement(bannerId, slotId)
}

func (bc BannerRecomender) addClick(bannerId, slotId, socialGroup int) error {
	registered, err := bc.storage.isBannerRegistered(bannerId, slotId)
	if err != nil {
		return err
	}

	if !registered {
		return errors.New("banner placement wasn't found")
	}

	err = bc.storage.addClick(bannerId, slotId, socialGroup)
	if err != nil {
		return err
	}

	return nil
}

func (bc BannerRecomender) getBannerToDisplay(slotId, socialGroupId int) (int, error) {
	stats, err := bc.storage.getStats(slotId, socialGroupId)
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
	err = bc.storage.addDisplay(bannerId, slotId, socialGroupId)
	if err != nil {
		return 0, err
	}
	return bannerId, nil
}

func main() {}
