package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllBannersAreShown(t *testing.T) {
	bc := NewBannerRecomender(NewDBMemory())
	// assert equality
	slotId := 1
	groupId := 1

	bannerId1, bannerId2, bannerId3, bannerId4 := 1, 2, 3, 4
	bannerId1Shown, bannerId2Shown, bannerId3Shown, bannerId4Shown := false, false, false, false

	if err := bc.addBannerPlacement(bannerId1, slotId); err != nil {
		log.Fatal(err)
	}
	if err := bc.addBannerPlacement(bannerId2, slotId); err != nil {
		log.Fatal(err)
	}
	if err := bc.addBannerPlacement(bannerId3, slotId); err != nil {
		log.Fatal(err)
	}
	if err := bc.addBannerPlacement(bannerId4, slotId); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 100; i++ {
		bannerId, err := bc.getBannerToDisplay(slotId, groupId)
		if err != nil {
			log.Println(err)
		}

		switch bannerId {
		case bannerId1:
			bannerId1Shown = true
		case bannerId2:
			bannerId2Shown = true
		case bannerId3:
			bannerId3Shown = true
		case bannerId4:
			bannerId4Shown = true
		}

		if bannerId == bannerId1 {
			if err = bc.addClick(bannerId1, slotId, groupId); err != nil {
				log.Fatal(err)
			}
		}
	}

	assert.True(t, bannerId1Shown)
	assert.True(t, bannerId2Shown)
	assert.True(t, bannerId3Shown)
	assert.True(t, bannerId4Shown)
}

func TestMostPopularBanner(t *testing.T) {
	bc := NewBannerRecomender(NewDBMemory())
	// assert equality
	slotId := 1
	groupId := 1

	bannerId1, bannerId2, bannerId3, bannerId4 := 1, 2, 3, 4

	if err := bc.addBannerPlacement(bannerId1, slotId); err != nil {
		log.Fatal(err)
	}
	if err := bc.addBannerPlacement(bannerId2, slotId); err != nil {
		log.Fatal(err)
	}
	if err := bc.addBannerPlacement(bannerId3, slotId); err != nil {
		log.Fatal(err)
	}
	if err := bc.addBannerPlacement(bannerId4, slotId); err != nil {
		log.Fatal(err)
	}

	bannerId1DisplayCount := 0
	totalDisplays := 100000

	for i := 0; i < totalDisplays; i++ {
		bannerId, err := bc.getBannerToDisplay(slotId, groupId)
		if err != nil {
			log.Println(err)
		}

		if bannerId == bannerId1 {
			if err = bc.addClick(bannerId1, slotId, groupId); err != nil {
				log.Fatal(err)
			}
			bannerId1DisplayCount += 1
			//log.Fatal("inc")
		}
	}

	log.Println("displays = ", bannerId1DisplayCount)
	log.Println("total displays = ", totalDisplays)
	log.Println("koef = ", float64(bannerId1DisplayCount)/ float64(totalDisplays))
	koef := float64(bannerId1DisplayCount)/ float64(totalDisplays)
	assert.True(t, koef < 0.72)
	assert.True(t, koef > 0.70)
}
