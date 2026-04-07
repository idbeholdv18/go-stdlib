package main

import (
	"errors"
	"log"
	"os"
	"time"
)

type FileState struct {
	modTime time.Time
	size    int64
	exists  bool
}

func InitFileState(filename string) *FileState {
	info, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return &FileState{
			exists:  false,
			size:    0,
			modTime: time.Time{},
		}
	}
	if err != nil {
		log.Fatal(err)
	}
	return &FileState{
		exists:  true,
		modTime: info.ModTime(),
		size:    info.Size(),
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("filename must be specified")
	}

	filename := os.Args[1]

	ticker := time.NewTicker(time.Duration(time.Second))
	defer ticker.Stop()

	fileState := InitFileState(filename)

	for range ticker.C {
		fileState.CheckIsUpdated(filename)
	}
}

func (prev *FileState) CheckIsUpdated(filename string) {
	info, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		// файл был удален
		if prev.exists {
			log.Printf("file %s has been removed", filename)
			prev.exists = false
		}
		return
	}
	if err != nil {
		log.Println(err)
		return
	}

	// файл был создан
	if !prev.exists {
		log.Printf("file %s has been created", filename)
		prev.exists = true
		prev.modTime = info.ModTime()
		prev.size = info.Size()
		return
	}

	// файл был обновлен
	if info.ModTime() != prev.modTime || info.Size() != prev.size {
		log.Printf("file %s has been updated", filename)
		prev.modTime = info.ModTime()
		prev.size = info.Size()
	}
}
