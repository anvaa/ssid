package server

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"srv/srv_conf"
)

var LogDb *gorm.DB
var logChannel = make(chan GinLog, 100)

type GinLog struct {
	ID        uint `gorm:"primaryKey"`
	Timestamp time.Time
	Method    string
	Path      string
	Status    int
	Latency   time.Duration
	ClientIP  string
}

func init() {
	go func() {
		for logEntry := range logChannel {
			LogDb.Create(&logEntry)
		}
	}()
}

func ginLoggerDatabase(r *gin.Engine) {
	dbpath := srv_conf.DataDir + "/ginlogs.db"
	var err error
	LogDb, err = gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database", dbpath)
	}
	LogDb.Exec("PRAGMA foreign_keys = ON;")
	LogDb.Exec("PRAGMA journal_mode=WAL;")
	// LogDb.Exec("PRAGMA synchronous=NORMAL;")
	// LogDb.Exec("PRAGMA cache_size=10000;")
	LogDb.Exec("PRAGMA temp_store=MEMORY;")
	// LogDb.Exec("PRAGMA mmap_size=30000000000;")
	// LogDb.Exec("PRAGMA threads=4;")
	// LogDb.Exec("PRAGMA page_size=4096;")
	LogDb.Exec("PRAGMA auto_vacuum=FULL;")

	LogDb.AutoMigrate(&GinLog{})

	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		gl := GinLog{
			Timestamp: start,
			Method:    c.Request.Method,
			Path:      c.Request.URL.Path,
			Status:    c.Writer.Status(),
			Latency:   latency,
			ClientIP:  c.ClientIP(),
		}
		logChannel <- gl
	})
}
