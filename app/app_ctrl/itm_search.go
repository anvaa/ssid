package app_ctrl

import (
	"log"
	"app/app_db"
	"app/app_models"
)

func Itms_SearchMulti(locid, typid, manid, staid string) ([]app_models.Items, int) {
	// Search for an item
	var itms []app_models.Items
	
	var sql string
	if locid != "0" {
		sql = "locid = " + locid
	}
	if typid != "0" {
		if sql != "" {
			sql += " AND "
		}
		sql += "typid = " + typid
	}
	if manid != "0" {
		if sql != "" {
			sql += " AND "
		}
		sql += "manid = " + manid
	}
	if staid != "0" {
		if sql != "" {
			sql += " AND "
		}
		sql += "staid = " + staid
	}

	// Search for the item
	err := app_db.AppDB.Where(sql).Find(&itms).Error
	if err != nil {
		log.Println(err)
		return nil, 0
	}

	return itms, len(itms)
}