package app_ctrl

import (
	"log"
	"app/app_db"
	"app/app_models"
)
func Itms_SearchMulti(locid, typid, manid, staid string) ([]app_models.Items, int) {
	// Search for an item
	var itms []app_models.Items
	
	conditions := []string{}
	if locid != "0" {
		conditions = append(conditions, "locid = "+locid)
	}
	if typid != "0" {
		conditions = append(conditions, "typid = "+typid)
	}
	if manid != "0" {
		conditions = append(conditions, "manid = "+manid)
	}
	if staid != "0" {
		conditions = append(conditions, "staid = "+staid)
	}

	sql := ""
	if len(conditions) > 0 {
		sql = conditions[0]
		for _, condition := range conditions[1:] {
			sql += " AND " + condition
		}
	}

	// Search for the item
	err := app_db.AppDB.Where(sql).Find(&itms).Error
	if err != nil {
		log.Println(err)
		return nil, 0
	}

	return itms, len(itms)
}
