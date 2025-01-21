package app_ctrl

import (
	"app/app_conf"
	"app/app_models"
	"srv/global"
	"strconv"
)

func inv_DoSearch(srcvals app_models.DoSearch, isList bool) ([]app_models.ItemsWeb, int) {
	var itms []app_models.Items
	var c_itms int

	// If search criteria is given, search for items.
	if hasSearchCriteria(srcvals) {
		itms, c_itms = Itms_SearchMulti(
			global.IntToString(srcvals.Locid),
			global.IntToString(srcvals.Typid),
			global.IntToString(srcvals.Manid),
			global.IntToString(srcvals.Staid),
		)
	} else {
		itms = nil
		c_itms = 0
	}

	// Prepare the items for the web view.
	itmsw := prepareItemsWeb(itms, isList)

	return itmsw, c_itms
}

func hasSearchCriteria(srcvals app_models.DoSearch) bool {
	return srcvals.Locid > 0 || srcvals.Typid > 0 || srcvals.Manid > 0 || srcvals.Staid > 0
}

func prepareItemsWeb(itms []app_models.Items, isList bool) []app_models.ItemsWeb {
	var itmsw []app_models.ItemsWeb
	for _, itm := range itms {
		desc := getDescription(itm.Description, isList)
		itmw := app_models.ItemsWeb{
			Itmid:       strconv.Itoa(itm.Itmid),
			Description: desc,
			Price:       strconv.FormatFloat(itm.Price, 'f', 2, 64),
			Serial:      itm.Serial,
			Updtime:     itm.UpdatedAt.Format("2006-01-02"),
			Loc:         Loc_GetLocName(itm.Locid),
			Typ:         Typ_GetTypName(itm.Typid),
			Man:         Man_GetManName(itm.Manid),
			Sta:         Sta_GetStatName(itm.Staid),
			Uid:         strconv.Itoa(itm.UserId),
		}
		itmsw = append(itmsw, itmw)
	}
	return itmsw
}

func getDescription(description string, isList bool) string {
	if isList {
		return description
	}
	return global.ShortenText(description, app_conf.TxtLength())
}
