package main

import (
	"fmt"
	"os"
	"path/filepath"

	controller "github.com/ICOMP-UNC/2023---soii---laboratorio-6-FrancoNB/controllers"
	"github.com/ICOMP-UNC/2023---soii---laboratorio-6-FrancoNB/repositories/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func InitLogs(err_path string, out_path string) (*os.File, *os.File) {
	logFolderPath, err := filepath.Abs("/var/log/lab6")

	if err != nil {
		log.Fatal("Error getting absolute path of logs folder: ", err)
	}

	err = os.MkdirAll(logFolderPath, os.ModePerm)

	if err != nil {
		log.Fatal("Error creating log folder: ", err)
	}

	stdoutLog, err := os.OpenFile(filepath.Join(logFolderPath, out_path), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("Error opening log file for stdout: ", err)
	}

	stderrLog, err := os.OpenFile(filepath.Join(logFolderPath, err_path), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("Error opening log file for stderr: ", err)
	}

	return stdoutLog, stderrLog
}

func InitProcessingDB() {
	folderPath, err := filepath.Abs("/var/lib/lab6")

	if err != nil {
		log.Fatal("Error getting absolute path of logs folder: ", err)
	}

	err = os.MkdirAll(folderPath, os.ModePerm)

	if err != nil {
		log.Fatal("Error creating log folder: ", err)
	}

	database.ProcessingInitDBConnection(folderPath + "/processing.db")
}

func StartUserService() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.POST("/api/users/login", controller.Login)
	r.GET("/api/users/listall", controller.AuthMiddleware(), controller.ListAll)
	r.POST("/api/users/createuser", controller.AuthMiddleware(), controller.Register)

	err := r.Run(":8555")

	if err != nil {
		fmt.Println("Error starting user service: ", err)
	}
}

func StartProcessingService() {
	InitProcessingDB()

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.POST("/api/processing/submit", controller.AuthMiddleware(), controller.Submit)
	r.GET("/api/processing/summary", controller.Summary)

	err := r.Run(":8556")

	if err != nil {
		fmt.Println("Error starting processing service: ", err)
	}

	defer database.ProcessingCloseDBConnection()
}

func main() {
	if os.Geteuid() != 0 {
		log.Fatal("This program needs to run with root privileges.")
	}

	stdoutLog, stderrLog := InitLogs("errors.log", "outputs.log")

	log.SetOutput(stderrLog)

	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}

	logrus.SetFormatter(formatter)

	gin.DefaultWriter = stdoutLog
	gin.DefaultErrorWriter = stderrLog

	go StartUserService()
	go StartProcessingService()

	defer stdoutLog.Close()
	defer stderrLog.Close()

	select {}
}
