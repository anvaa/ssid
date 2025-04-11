package main

import (
	"srv/global"
	"srv/server"
	"srv/srv_conf"
	"srv/srv_sec"
	"srv_int"
	users "usr"

	"app"
	"app/app_conf"
	"app/app_db"

	"os/signal"
	"syscall"
	"time"

	"fmt"
	"log"
	"os"
)

var app_dir string = getWD()
var Hostip []string
var is_app bool = false

var CloseChan = make(chan os.Signal, 1)

func init() {

	readArgs()
	setupCloseHandler()

	// Check/make srv.yaml
	err := srv_int.ServerInit(app_dir)
	if err != nil {
		log.Fatal(err)
	}

	// Check/make userdb, usr.yaml & user http.FileSystem
	err = users.UserInit(app_dir)
	if err != nil {
		log.Fatal(err)
	}

	// Check/make appdb, app.yaml & app http.FileSystem
	err = app.AppInit(app_dir)
	if err != nil {
		log.Fatal(err)
	}

	app_conf.StartTime = time.Now().Unix()

}

func main() {

	r := server.InitWebServer()

	addr := fmt.Sprintf(":%d", srv_conf.GetInt("server_port"))
	go printsrvinfo(addr)

	if is_app {
		log.Println("Running as app ...")
		r.Run(addr)
	} else {
		log.Println("Running as server ...")
		r.RunTLS(addr, srv_sec.CertFile, srv_sec.KeyFile)
	}

}

func getWD() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("WorkDir: %s", wd)
	return wd
}

func readArgs() {
	// Check for command line arguments
	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); {
			switch os.Args[i] {
			case "-h", "--help":
				fmt.Println("Usage: srv [options]")
				fmt.Println("Options:")
				fmt.Println("  -h, --help\tShow this help message")
				fmt.Println("  -a, --app\tRun as app")
				os.Exit(0)
			case "-a", "--app":
				_ = os.Setenv("is_app", "true")
				is_app = true
			default:
				fmt.Printf("Unknown option: %s\n", os.Args[i])
				os.Exit(1)
			}
			i++ // Increment the loop variable
		}
	}
}

func printsrvinfo(adr string) {

	Hostip, _ = global.GetIPv4Addresses()

	mode := "https"
	if is_app {
		mode = "http"
	}

	for _, ip := range Hostip {
		log.Printf("Webserver %s://%s%s", mode, ip, adr)
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
