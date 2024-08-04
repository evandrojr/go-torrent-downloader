package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type BindFile struct {
	Magnet string `form:"magnet" binding:"required"`
}

const ShellToUse = "bash"

func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")
	router.POST("/download", func(c *gin.Context) {
		var bindFile BindFile

		// Bind file
		if err := c.ShouldBind(&bindFile); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}

		// Save uploaded file
		file := bindFile.Magnet

		out, errout, err := Shellout(`cd ~/Downloads && aria2c "` + file + `" > ~/p/torrent-downloader/public/log.txt`)
		if err != nil {
			log.Printf("error: %v\n", err)
			// c.String(http.ErrBodyNotAllowed, err.Error())
		}
		fmt.Println("--- stdout ---")
		fmt.Println(out)
		fmt.Println("--- stderr ---")
		fmt.Println(errout)

		c.String(http.StatusOK, fmt.Sprintf("File %s .", file))
	})
	router.Run(":8091")
}
