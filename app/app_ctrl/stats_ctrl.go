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

	if err := getItemCount(&stats); err != nil {
		return stats, err
	}

	if stats.Itm_count == 0 {
		return stats, errors.New("no items found")
	}

	if err := getLocs(&stats); err != nil {
		return stats, err
	}

	for _, locID := range stats.Locs {
		if err := processLocation(locID, &stats); err != nil {
			return stats, err
		}
	}

	if err := getTotalPrice(&stats); err != nil {
		return stats, err
	}

	if err := getFirstAndLastDates(&stats); err != nil {
		return stats, err
	}

	return stats, nil
}

func getItemCount(stats *app_models.Stats) error {
	return app_db.AppDB.Table("items").
		Where("deleted_at IS NULL").
		Count(&stats.Itm_count).Error
}

func getLocs(stats *app_models.Stats) error {
	return app_db.AppDB.Table("items").
		Select("locid").
		Group("locid").
		Order("locid ASC").
		Find(&stats.Locs).Error
}

func processLocation(locID int64, stats *app_models.Stats) error {
	var loc app_models.Loc

	if err := getLocName(locID, &loc); err != nil {
		return err
	}

	if err := getDistinctTypes(locID, &loc); err != nil {
		return err
	}

	if err := getTotalPriceAtLocation(locID, &loc); err != nil {
		return err
	}

	if err := getTypeNamesAndCounts(locID, &loc); err != nil {
		return err
	}

	stats.LocType = append(stats.LocType, loc)

	fmt.Printf("Location: %s\nTot itms: %d\nitems: %d\n", loc.Locname, loc.ItemsTotalCount, loc.Types)
	return nil
}

func getLocName(locID int64, loc *app_models.Loc) error {
	return app_db.AppDB.Table("loc_names").Select("locname").Where("id = ?", locID).Scan(&loc.Locname).Error
}

func getDistinctTypes(locID int64, loc *app_models.Loc) error {
	return app_db.AppDB.Table("items").Select("DISTINCT typid").Where("locid = ?", locID).Find(&loc.Types).Error
}

func getTotalPriceAtLocation(locID int64, loc *app_models.Loc) error {
	return app_db.AppDB.Table("items").Select("SUM(price)").Where("locid = ? AND deleted_at is null", locID).Scan(&loc.LocTotPrice).Error
}

func getTypeNamesAndCounts(locID int64, loc *app_models.Loc) error {
	var totalNrOfItems int64
	for _, typeID := range loc.Types {
		var typ app_models.Typ
		if err := getTypeName(typeID, &typ); err != nil {
			return err
		}
		if err := getTypeCount(locID, typeID, &typ); err != nil {
			return err
		}
		loc.Loctype = append(loc.Loctype, typ)
		loc.ItemsCount++
		totalNrOfItems += typ.TypCount
	}
	loc.ItemsTotalCount = totalNrOfItems
	return nil
}

func getTypeName(typeID int64, typ *app_models.Typ) error {
	return app_db.AppDB.Table("typ_names").Select("typname").Where("id = ?", typeID).Scan(&typ.Typname).Error
}

func getTypeCount(locID, typeID int64, typ *app_models.Typ) error {
	return app_db.AppDB.Table("items").Where("locid = ? AND typid = ? AND deleted_at is null", locID, typeID).Count(&typ.TypCount).Error
}

func getTotalPrice(stats *app_models.Stats) error {
	if err := app_db.AppDB.Table("items").Select("SUM(price)").Scan(&stats.Price).Error; err != nil {
		return err
	}
	stats.Total_Price = accounting.FormatNumber(stats.Price, 2, " ", ",")
	return nil
}

func getFirstAndLastDates(stats *app_models.Stats) error {
	if err := app_db.AppDB.Table("items").Select("MIN(created_at)").Scan(&stats.FirstDate).Error; err != nil {
		return err
	}
	if err := app_db.AppDB.Table("items").Select("MAX(created_at)").Scan(&stats.LastDate).Error; err != nil {
		return err
	}
	stats.FirstDate = stats.FirstDate[:10]
	stats.LastDate = stats.LastDate[:10]
	return nil
}
