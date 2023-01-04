package store

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

/**
 * @Description
 * @Author duanfeng.zhang
 * @Date 2023/1/4 12:42
 **/
func LogInit() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	io1 := &bytes.Buffer{}
	io2 := os.Stderr
	file := "/data/blog/log/blog" + time.Now().Format("20060102") + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if nil != err {
		panic(err)
	}
	logrus.SetReportCaller(true)
	logrus.SetOutput(io.MultiWriter(io1, io2, logFile))
}
