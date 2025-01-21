package app_db

import (
	"srv/filefunc"
	"srv/srv_conf"
	"strings"

	"app/app_models"

	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
	barcode "github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"

	"image/png"
)

func Itm_AddUpd(newitm app_models.Items) error {
	
	var itm app_models.Items
	// Check if the item already exists
	res := AppDB.Where("itmid = ?", newitm.Itmid).
		Attrs(
			app_models.Items{
				Locid: newitm.Locid,
				Manid: newitm.Manid,
				Typid: newitm.Typid,
				Staid: 1212090603,

				Serial:      strings.Trim(newitm.Serial, " "),
				Description: strings.Trim(newitm.Description, " "),
				Price:       newitm.Price,
				UserId:      newitm.UserId,
				Itmid:       Itm_NewItmId(),
			},
		).
		FirstOrCreate(&itm)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		// update if exsits
		err := AppDB.Model(&itm).Updates(
			app_models.Items{
				Locid: newitm.Locid,
				Manid: newitm.Manid,
				Typid: newitm.Typid,
				Staid: newitm.Staid,

				Serial:      strings.Trim(newitm.Serial, " "),
				Description: strings.Trim(newitm.Description, " "),
				Price:       newitm.Price,
				UserId:      newitm.UserId,
			},
		).Error
		if err != nil {
			return err
		}

	}

	// update QR code
	go Itm_MakeQRCode(itm)

	if res.RowsAffected == 1 {
		// if new record add new status
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
	err := AppDB.Where("itmid = ?", itmid).Delete(&app_models.Items{}).Error
	if err != nil {
		return err
	}

	// delete status history
	err = AppDB.Where("itmid = ?", itmid).Delete(&app_models.Status_History{}).Error
	if err != nil {
		return err
	}

	// delete QR code
	imgpath := fmt.Sprintf("%s/%d.png", srv_conf.QRImgDir, itmid)
	if filefunc.IsExists(imgpath) {
		err := filefunc.DeleteFile(imgpath)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func Itm_UpdCurStatus(itmid any, staid any) error {
	// Update the current status of an item
	err := AppDB.Model(&app_models.Items{}).Where("itmid = ?", itmid).Update("staid", staid).Error
	if err != nil {
		return err
	}
	return nil
}

func Itm_NewItmId() int {
	// Generate a new item ID
	return int(uuid.New().ID())
}

func Itm_MakeQRCode(itm app_models.Items) error {
	// Generate QR code for an item

	imgpath := fmt.Sprintf("%s/%d.png", srv_conf.QRImgDir, itm.Itmid)
	if filefunc.IsExists(imgpath) {
		err := filefunc.DeleteFile(imgpath)
		if err != nil {
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
	err = qrcx.WriteFile(256, imgpath)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = Itm_MakeBARCode(itm.Serial, itm.Itmid)
	if err != nil {
		return err
	}

	return nil
}

func Itm_MakeBARCode(itm_serial string, itmid any) error {
	// Generate a barcode for an item
	imgpath := fmt.Sprintf("%s/%d.png", srv_conf.BarImgDir, itmid)
	if filefunc.IsExists(imgpath) {
		err := filefunc.DeleteFile(imgpath)
		if err != nil {
			return err
		}
	}

	// Write the barcode to a file
	file, _ := filefunc.CreateFile(imgpath)
	defer file.Close()
	
	writer := oned.NewCode128Writer()
	barCode, err := writer.Encode(itm_serial, barcode.BarcodeFormat_CODE_128, 200, 35, nil)
	if err != nil {
    	return err
	}
	
	png.Encode(file, barCode)

	return nil
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
	err := AppDB.Where("id = ?", id).First(&sta).Error
	if err != nil {
		return "nil"
	}

	return sta.Staname
}

func Sta_HistoryDelete(itmid any) error {
	// Delete the status history of an item
	err := AppDB.Where("itmid = ?", itmid).Delete(&app_models.Status_History{}).Error
	if err != nil {
		return err
	}
	return nil
}
