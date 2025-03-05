package app_embed

import (
	"embed"

	"srv/filefunc"
	"srv/srv_conf"
)

//go:embed assets/*
//go:embed css/*
//go:embed html/*
//go:embed js/*
//go:embed media/*
var static embed.FS

func App_EmbedFiles() error {

	// write app_embed to disk
	err := filefunc.WriteWebFSToDisk(srv_conf.StaticDir, static)
	if err != nil {
		return err
	}
	
	return nil

}
