package app_ctrl

import (
	"app/app_db"
	"app/app_models"
	"fmt"

	"errors"

	"github.com/leekchan/accounting"
)

// GetStats returns the stats for the application
func GetStats() (app_models.Stats, error) {
	var stats app_models.Stats

	// Get the number of items
	err := app_db.AppDB.Table("items").
		Where("deleted_at IS NULL").
		Count(&stats.Itm_count).Error
	if err != nil {
		return stats, err
	}

	if stats.Itm_count == 0 {
		return stats, errors.New("no items found")
	}

	// Get locid from items where deleted_at is null. sort asc
	err = app_db.AppDB.Table("items").
		Select("locid").
		Group("locid").
		Order("locid ASC").
		Find(&stats.Locs).Error
	if err != nil {
		return stats, err
	}

	for _, v := range stats.Locs {
		
		var lt app_models.Loc

		// add locname. Find and count types at location
		app_db.AppDB.Table("loc_names").Select("locname").Where("id = ?", v).Scan(&lt.Locname)

		// get distinct type ids at location
		app_db.AppDB.Table("items").Select("DISTINCT typid").Where("locid = ?", v).Find(&lt.Types)

		// get total price of items at location
		app_db.AppDB.Table("items").Select("SUM(price)").Where("locid = ? AND deleted_at is null", v).Scan(&lt.LocPrice)
		lt.LocTotPrice = accounting.FormatNumber(lt.LocPrice, 2, " ", ",")
		
		
		// get type names and count of items at location
		var totalNrOfItems int64 
		for _, t := range lt.Types {
			var tp app_models.Typ
			// add typename.
			app_db.AppDB.Table("typ_names").Select("typname").Where("id = ?", t).Scan(&tp.Typname)
			app_db.AppDB.Table("items").Where("locid = ? AND typid = ? AND deleted_at is null", v, t).Count(&tp.TypCount)
			lt.Loctype = append(lt.Loctype, tp)
			lt.ItemsCount = +1
			totalNrOfItems = totalNrOfItems + tp.TypCount
		}
		lt.ItemsTotalCount = totalNrOfItems
		stats.LocType = append(stats.LocType, lt)
		
		txt := fmt.Sprintf("Location: %s\nTot itms: %d\nitems: %d\n", 
			lt.Locname, lt.ItemsTotalCount, lt.Types)
		fmt.Println(txt)
	}


	fmt.Println(stats)
	// add total cost of all items
	app_db.AppDB.Table("items").Select("SUM(price)").Scan(&stats.Price)
	// format float64 into a currency string. Separate by thousands with space?
	stats.Total_Price = accounting.FormatNumber(stats.Price, 2, " ", ",")

	// get first and last dates of items
	app_db.AppDB.Table("items").Select("MIN(created_at)").Scan(&stats.FirstDate)
	app_db.AppDB.Table("items").Select("MAX(created_at)").Scan(&stats.LastDate)

	stats.FirstDate = stats.FirstDate[:10]
	stats.LastDate = stats.LastDate[:10]

	return stats, nil
}
