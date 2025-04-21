package filefunc

import (

	"log"
	"fmt"
	"os"
	"embed"
	"io/fs"
	"path/filepath"
	"strings"

	"app/app_models"
	"app/app_menu"

)

func IsExists(path string) bool {
	// if folder/file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// Create folder
func CreateFolder(path string) error {
	// create folder
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	// log.Println("Created folder:", path)
	return nil
}

// Create file
func CreateFile(path string) (*os.File, error) {
	// create or overwrite file
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	// log.Println("Created file:", path)
	return file, nil
}

// Delete file
func DeleteFile(filepath string) error {

	_, err := os.Stat(filepath)
	os.IsExist(err)
	if err != nil {
		return err
	}
	os.Remove(filepath)
	log.Println("Deleted file:", filepath)
	return nil
}

// Delete folder with content
func DeleteFolder_FR(path string) error {
	// force delete folder and all sub folders with content
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	log.Println("Deleted folder and content:", path)
	return nil
}

func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetFileList returns a list of files in a directory
func GetFileList(dir string) []string {
	var files []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files
}

// GetFileListByExt returns a list of files in a directory with a specific extension
func GetFileListByExt(dir string, ext string) []string {
	var files []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, ext) {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func ExportSearchResult(filepath, filename string, items []app_models.ItemsWeb) error {
	// make a csv file
	
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	// write the header
	menutitle := app_menu.GetMenuTitles()
	header := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s\n", "Item ID", menutitle[0], menutitle[1], menutitle[3], menutitle[2], menutitle[4], "Status", "Date")
	f.WriteString(header)

	// write the items
	for _, item := range items {
		// csv format: remove commas
		serial := strings.ReplaceAll(item.Serial, ",", "")
		desc := strings.ReplaceAll(item.Description, ",", "") 
		line := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s\n",
			item.Itmid, item.Loc, item.Typ, serial, item.Man, desc, item.Sta, item.Updtime)
		f.WriteString(line)
	}

	return nil

}

func WriteWebFSToDisk(folder string, wfs embed.FS) error {

	log.Println("Writing embed.FS to disk:", folder)
	// loop thru static and write folders and files to static on disk
	err := fs.WalkDir(wfs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			// create folder
			folder := folder + "/" + path
			if !IsExists(folder) {
				CreateFolder(folder)
			}
		} else {
			// create file
			folder := folder + "/" + path
			if !IsExists(folder) {
				file, err := CreateFile(folder)
				if err != nil {
					return err
				}
				data, err := wfs.ReadFile(path)
				if err != nil {
					return err
				}
				file.Write(data)
				file.Close()
			}
		}

		return nil
	},
	)
	if err != nil {
		return err
	}

	return nil

}
