package main

type Stat struct {
	bannerId int
	prob     int
}

type ClickDisplay struct {
	bannerId      int
	slotId        int
	socialGroupId int
	clicks        int
	displays      int
}

type BannerPlacement struct {
	slotId   int
	bannerId int
}

type DBMemory struct {
	bannerPlacements []BannerPlacement
	clicksDisplays   []ClickDisplay
}

func NewDBMemory() *DBMemory {
	return &DBMemory{bannerPlacements: []BannerPlacement{}, clicksDisplays: []ClickDisplay{}}
}

func (db *DBMemory) addBannerPlacement(bannerId, slotId int) error {
	db.bannerPlacements = append(db.bannerPlacements, BannerPlacement{
		slotId:   slotId,
		bannerId: bannerId,
	})

	return nil
}

func (db *DBMemory) isBannerRegistered(bannerId, slotId int) (bool, error) {
	for _, bp := range db.bannerPlacements {
		if bp.bannerId == bannerId && bp.slotId == slotId {
			return true, nil
		}
	}

	return false, nil
}

func (db *DBMemory) addClick(bannerId, slotId, socialGroupId int) error {
	for i, cd := range db.clicksDisplays {
		if cd.bannerId == bannerId && cd.slotId == slotId && cd.socialGroupId == socialGroupId {
			db.clicksDisplays[i].clicks += 1
			return nil
		}
	}

	db.clicksDisplays = append(db.clicksDisplays, ClickDisplay{
		bannerId:      bannerId,
		slotId:        slotId,
		socialGroupId: socialGroupId,
		clicks:        1,
		displays:      0,
	})

	return nil
}

func (db *DBMemory) addDisplay(bannerId, slotId, socialGroupId int) error {
	for i, cd := range db.clicksDisplays {
		if cd.bannerId == bannerId && cd.slotId == slotId && cd.socialGroupId == socialGroupId {
			db.clicksDisplays[i].displays += 1
			return nil
		}
	}

	db.clicksDisplays = append(db.clicksDisplays, ClickDisplay{
		bannerId:      bannerId,
		slotId:        slotId,
		socialGroupId: socialGroupId,
		clicks:        0,
		displays:      1,
	})

	return nil
}

func (db *DBMemory) removeBannerPlacement(bannerId, slotId int) error {
	newBannerPlacements := make([]BannerPlacement, 0, len(db.bannerPlacements)-1)
	newClicksDisplays := make([]ClickDisplay, 0)

	for _, bp := range db.bannerPlacements {
		if bp.bannerId == bannerId && bp.slotId == slotId {
			continue
		}
		newBannerPlacements = append(newBannerPlacements, bp)
	}

	for _, cd := range db.clicksDisplays {
		if cd.bannerId == bannerId && cd.slotId == slotId {
			continue
		}
		newClicksDisplays = append(newClicksDisplays, cd)
	}

	db.bannerPlacements = newBannerPlacements
	db.clicksDisplays = newClicksDisplays

	return nil
}

func (db *DBMemory) getBannersForSlot(slotId int) ([]int, error) {
	// TODO: should be an id
	var banners []int

	for _, bp := range db.bannerPlacements {
		if bp.slotId == slotId {
			banners = append(banners, bp.bannerId)
		}
	}

	return banners, nil
}

func (db *DBMemory) getStats(slotId, socialGroupId int) ([]Stat, error) {
	stats := make([]Stat, 0)

	for _, cd := range db.clicksDisplays {
		if cd.slotId == slotId && cd.socialGroupId == socialGroupId {
			prob := 0

			if cd.displays != 0 {
				prob = cd.clicks / cd.displays
			}

			stats = append(stats, Stat{
				bannerId: cd.bannerId,
				prob:     prob,
			})
		}
	}

	return stats, nil
}
