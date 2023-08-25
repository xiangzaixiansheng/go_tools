package log_util

import (
	"io"
	"log"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

var LogrusObj *logrus.Logger

func init() {
	if LogrusObj != nil {
		src, _ := setOutputFile()
		//设置输出
		LogrusObj.Out = src
		return
	}
	//实例化
	logger := logrus.New()
	writer1_file, _ := setOutputFile() //文件
	writer2_console := os.Stdout       //终端

	//同时写入终端和文件中
	logger.SetOutput(io.MultiWriter(writer1_file, writer2_console))
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 取消线程安全
	logger.SetNoLock()
	//设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//增加行号和文件名
	logger.SetReportCaller(true)

	/*
		加个hook形成ELK体系
		但是考虑到一些同学一下子接受不了那么多技术栈，
		所以这里的ELK体系加了注释，如果想引入可以直接注释去掉，
		如果不想引入这样注释掉也是没问题的。
	*/
	//hook := model.EsHookLog()
	//logger.AddHook(hook)
	LogrusObj = logger
}

func setOutputFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return src, nil
}

/*
    log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	//log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")

	log.WithFields(log.Fields{
		"animal": "dog", //增加字段打印
	}).Info("dog is here")
**/
