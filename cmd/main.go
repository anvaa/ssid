package main

import (
	
	"srv/global"
	"srv/server"
	"srv/srv_conf"
	"srv/srv_sec"
	"users"
	"srv_int"

	"app"
	"app/app_db"
	"app/app_conf"

	"os/signal"
	"syscall"
	"time"

	"log"
	"fmt"
	"os"
)

var WD string = getWD()
var Hostip []string

var CloseChan = make(chan os.Signal, 1)

func init() {

	setupCloseHandler()
	
	// Check/make srv.yaml 
	err := srv_int.ServerInit(WD)
	if err != nil {
		log.Fatal(err)
	}

	// Check/make userdb, usr.yaml & user http.FileSystem
	err = users.UserInit(srv_conf.AppDir)
	if err != nil {
		log.Fatal(err)
	}

	// Check/make appdb, app.yaml & app http.FileSystem
	err = app.AppInit(srv_conf.AppDir)
	if err != nil {
		log.Fatal(err)
	}

	app_conf.StartTime = time.Now().Unix()

}

func main() {

	r := server.InitWebServer()

	addr := fmt.Sprintf(":%d", srv_conf.GetInt("server_port"))
	go printsrvinfo(addr)

	r.RunTLS(addr, srv_sec.CertFile, srv_sec.KeyFile)

}

func getWD() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("WorkDir: %s", wd)
	return wd
}

func printsrvinfo(adr string) {

	Hostip, _ = global.GetIPv4Addresses()

	for _, ip := range Hostip {
		log.Printf("Server running at https://%s%s", ip, adr)
	}
}

func setupCloseHandler() {
	log.Println("SetupCloseHandler ...")
	// closeChan := make(chan os.Signal, 1) // Make the channel buffered with a capacity of 1
	signal.Notify(CloseChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-CloseChan
		log.Printf("\nClosing app ...\n")

		users.CloseUserDB()
		app_db.CloseAppDB()

		os.Exit(0)
	}()
}