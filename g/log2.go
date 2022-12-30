package g

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

/**
 * @Description
 * @Author duanfeng.zhang
 * @Date 2022/12/30 15:24
 **/

var Logger = &lumberjack.Logger{
	Filename:   "/data/blog/log/blog.%Y%m%d",
	MaxSize:    10, // megabytes
	MaxBackups: 3,
	MaxAge:     7, //days
}

//logrus同时写文件和终端
func InitLogrus() {
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006/01/02 - 15:04:05",
	})
	//同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(os.Stdout, Logger)
	logrus.SetOutput(fileAndStdoutWriter)
	//设置最低loglevel
	logrus.SetLevel(logrus.InfoLevel)
}
