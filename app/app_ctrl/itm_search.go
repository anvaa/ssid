package app_ctrl

import (
	"app/app_db"
	"app/app_models"
	"fmt"
	"log"
)

func Itms_SearchMulti(locid, typid, manid, staid, fadte, tdate string) ([]app_models.Items, int) {
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
	if fadte != "" {
		conditions = append(conditions, "updated_at >= '"+fadte+"'")
	}
	if tdate != "" {
		conditions = append(conditions, "updated_at <= '"+tdate+"'")
	}
	sql := ""
	if len(conditions) > 0 {
		sql = conditions[0]
		for _, condition := range conditions[1:] {
			sql += " AND " + condition
		}
	}
	fmt.Println("SQL:", sql)
	// Search for the item
	err := app_db.AppDB.Where(sql).Find(&itms).Error
	if err != nil {
		log.Println(err)
		return nil, 0
	}

	return itms, len(itms)
}
