package user_embed

import (
	"embed"

	"srv/filefunc"
	"srv/srv_conf"
)

//go:embed css/*
//go:embed html/*
//go:embed js/*
var static embed.FS

func User_EmbedFiles() error {

	if !srv_conf.IsGinModDebug() {
		
		// write user_embed to disk
		err := filefunc.WriteWebFSToDisk(srv_conf.StaticDir, static)
		if err != nil {
			return err
		}

	}

	return nil

}