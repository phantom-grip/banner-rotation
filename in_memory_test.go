package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	bc := NewBannerRecomender(NewDBMemory())
	// assert equality
	slotId := 1
	groupId := 1

	bannerId1, bannerId2, bannerId3, bannerId4 := 1, 2, 3, 4
	bannerId1Shown, bannerId2Shown, bannerId3Shown, bannerId4Shown := false, false, false, false

	bc.addBannerPlacement(bannerId1, slotId)
	bc.addBannerPlacement(bannerId2, slotId)
	bc.addBannerPlacement(bannerId3, slotId)
	bc.addBannerPlacement(bannerId4, slotId)

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
			bc.addClick(bannerId1, slotId, groupId)
		}
	}

	assert.True(t, bannerId1Shown)
	assert.True(t, bannerId2Shown)
	assert.True(t, bannerId3Shown)
	assert.True(t, bannerId4Shown)

}
