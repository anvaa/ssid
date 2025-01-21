package app_db

import (
	"fmt"
	"image/png"
	"log"
	"strings"

	"app/app_models"
	"srv/filefunc"
	"srv/srv_conf"

	"github.com/google/uuid"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
	"github.com/skip2/go-qrcode"
)

func Itm_AddUpd(newitm app_models.Items) error {
	var itm app_models.Items

	// Check if the item already exists
	res := AppDB.Where("itmid = ?", newitm.Itmid).
		Attrs(app_models.Items{
			Locid:       newitm.Locid,
			Manid:       newitm.Manid,
			Typid:       newitm.Typid,
			Staid:       1212090603,
			Serial:      strings.TrimSpace(newitm.Serial),
			Description: strings.TrimSpace(newitm.Description),
			Price:       newitm.Price,
			UserId:      newitm.UserId,
			Itmid:       Itm_NewItmId(),
		}).
		FirstOrCreate(&itm)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		// Update if exists
		err := AppDB.Model(&itm).Updates(app_models.Items{
			Locid:       newitm.Locid,
			Manid:       newitm.Manid,
			Typid:       newitm.Typid,
			Staid:       newitm.Staid,
			Serial:      strings.TrimSpace(newitm.Serial),
			Description: strings.TrimSpace(newitm.Description),
			Price:       newitm.Price,
			UserId:      newitm.UserId,
		}).Error
		if err != nil {
			return err
		}
	}

	// Update QR code
	go Itm_MakeQRCode(itm)

	if res.RowsAffected == 1 {
		// If new record, add new status
		sta := app_models.Status_History{
			Itmid:   itm.Itmid,
			UserId:  itm.UserId,
			Staid:   itm.Staid,
			Comment: "Item added to inventory",
		}
		err := AppDB.Create(&sta).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func Itm_Delete(itmid any) error {
	// Delete an item
	if err := AppDB.Where("itmid = ?", itmid).Delete(&app_models.Items{}).Error; err != nil {
		return err
	}

	// Delete status history
	if err := AppDB.Where("itmid = ?", itmid).Delete(&app_models.Status_History{}).Error; err != nil {
		return err
	}

	// Delete QR code
	imgpath := fmt.Sprintf("%s/%d.png", srv_conf.QRImgDir, itmid)
	if filefunc.IsExists(imgpath) {
		if err := filefunc.DeleteFile(imgpath); err != nil {
			log.Println(err)
		}
	}

	return nil
}

func Itm_UpdCurStatus(itmid, staid any) error {
	// Update the current status of an item
	return AppDB.Model(&app_models.Items{}).Where("itmid = ?", itmid).Update("staid", staid).Error
}

func Itm_NewItmId() int {
	// Generate a new item ID
	return int(uuid.New().ID())
}

func Itm_MakeQRCode(itm app_models.Items) error {
	// Generate QR code for an item
	imgpath := fmt.Sprintf("%s/%d.png", srv_conf.QRImgDir, itm.Itmid)
	if filefunc.IsExists(imgpath) {
		if err := filefunc.DeleteFile(imgpath); err != nil {
			fmt.Println(err)
		}
	}

	// Create the QR code data
	qrtxt := fmt.Sprintf("%s\nSN %s\nID %d", Typ_GetTypName(itm.Typid), itm.Serial, itm.Itmid)
	qrcx, err := qrcode.New(qrtxt, qrcode.Medium)
	if err != nil {
		return err
	}
	qrcx.DisableBorder = true

	// Write the QR code to a file
	if err := qrcx.WriteFile(256, imgpath); err != nil {
		fmt.Println(err.Error())
	}

	return Itm_MakeBARCode(itm.Serial, itm.Itmid)
}

func Itm_MakeBARCode(itm_serial string, itmid any) error {
	// Generate a barcode for an item
	imgpath := fmt.Sprintf("%s/%d.png", srv_conf.BarImgDir, itmid)
	if filefunc.IsExists(imgpath) {
		if err := filefunc.DeleteFile(imgpath); err != nil {
			return err
		}
	}

	// Write the barcode to a file
	file, _ := filefunc.CreateFile(imgpath)
	defer file.Close()

	writer := oned.NewCode128Writer()
	barCode, err := writer.Encode(itm_serial, gozxing.BarcodeFormat_CODE_128, 200, 35, nil)
	if err != nil {
		return err
	}

	return png.Encode(file, barCode)
}

func Typ_GetTypName(id any) string {
	var typ app_models.TypNames
	AppDB.Where("id = ?", id).First(&typ)

	if typ.Id == 0 {
		return "nil"
	}

	return typ.Typname
}

func Loc_GetLocName(id any) string {
	var loc app_models.LocNames
	AppDB.Where("id = ?", id).First(&loc)

	if loc.Id == 0 {
		return "nil"
	}

	return loc.Locname
}

func Man_GetManName(id any) string {
	var man app_models.ManNames
	AppDB.Where("id = ?", id).First(&man)

	if man.Id == 0 {
		return "nil"
	}

	return man.Manname
}

func Sta_GetStaName(id any) string {
	var sta app_models.StaNames
	if err := AppDB.Where("id = ?", id).First(&sta).Error; err != nil {
		return "nil"
	}

	return sta.Staname
}

func Sta_HistoryDelete(itmid any) error {
	// Delete the status history of an item
	return AppDB.Where("itmid = ?", itmid).Delete(&app_models.Status_History{}).Error
}
