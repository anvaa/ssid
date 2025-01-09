package srv_int

import (
	"srv/filefunc"
	"srv/srv_conf"
	"srv/server"

	"log"
)

func ServerInit(app_dir string) error {

	// write the srv_conf file
	err := srv_conf.WriteConfigFile(app_dir)
	if err != nil {
		return err
	}

	// Check for .crt/.key files
	server.CheckTLS(app_dir, srv_conf.GetInt("tls_keysize"))

	// Check for static folder
	CheckFolder()

	return nil

}

func CheckFolder() error {
	
	_stat_dir := srv_conf.StaticDir
	if filefunc.IsExists(_stat_dir) {
		log.Println("Deleting", _stat_dir)
		err := filefunc.DeleteFolder_FR(_stat_dir)
		if err != nil {
			return err
		}
	}

	err := filefunc.CreateFolder(_stat_dir)
	if err != nil {
		return err
	}

	assets_dir := srv_conf.AssetsDir
	if !filefunc.IsExists(assets_dir) {
		
		filefunc.CreateFolder(assets_dir)
		filefunc.CreateFolder(srv_conf.ReportsDir)
		filefunc.CreateFolder(srv_conf.QRImgDir)
		log.Println("Created assed dir:", assets_dir)
	}

	return nil
}