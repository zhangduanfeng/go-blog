package g

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

/**
 * @Description
 * @Author duanfeng.zhang
 * @Date 2022/12/30 14:47
 **/
func Initlog() {
	writer, _ := rotatelogs.New(
		"/data/ddd/blog.%Y%m%d",                    //每天
		rotatelogs.WithRotationTime(24*time.Hour),  //最小为1分钟轮询。默认60s  低于1分钟就按1分钟来
		rotatelogs.WithRotationCount(3),            //设置3份 大于3份 或到了清理时间 开始清理
		rotatelogs.WithRotationSize(100*1024*1024), //设置100MB大小,当大于这个容量时，创建新的日志文件

		// logFile+".%Y%m%d%H%M",                      //每分钟
		// rotatelogs.WithLinkName(logFile),           //生成软链，指向最新日志文件
		// rotatelogs.WithRotationTime(time.Minute),   //最小为1分钟轮询。默认60s  低于1分钟就按1分钟来
		// rotatelogs.WithRotationCount(3),            //设置3份 大于3份 或到了清理时间 开始清理
		// rotatelogs.WithRotationSize(100*1024*1024), //设置100MB大小,当大于这个容量时，创建新的日志文件
	)

	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006/01/02 - 15:04:05",
	})
	//同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(os.Stdout, writer)
	log.SetOutput(fileAndStdoutWriter)
	//设置最低loglevel
	log.SetLevel(log.DebugLevel)
}
