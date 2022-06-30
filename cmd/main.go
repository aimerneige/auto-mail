package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/aimerneige/auto-mail/config"
	"github.com/aimerneige/auto-mail/mail"
	"github.com/spf13/viper"
)

const timeFormatStr = "2006-01-02 15:04:05"

var frequency time.Duration
var subject string
var content string
var core int

func init() {
	// init config file
	configFileName := "config"
	configFileType := "yaml"
	configFilePath := "./config"
	config.InitConfig(configFileName, configFileType, configFilePath)
	// read config
	mailHost := viper.GetString("host")
	mailPort := viper.GetInt("port")
	mailUser := viper.GetString("user")
	mailPass := viper.GetString("pass")
	mailTo := viper.GetStringSlice("to")
	mailFreq := viper.GetInt("freq")
	mailSubject := viper.GetString("subject")
	mailContent := viper.GetString("content")
	sysCore := viper.GetInt("core")
	// mail service
	mail.InitMailService(mailHost, mailPort, mailUser, mailPass, mailTo)
	// frequency
	frequency = time.Millisecond * time.Duration(mailFreq)
	// mail info
	subject = mailSubject
	content = mailContent
	// sys core
	core = sysCore
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < core; i++ {
		wg.Add(1)
		go func() {
			for {
				err := mail.SendMail(subject, content)
				msg := "Successful!"
				if err != nil {
					msg = "Fail!"
				}
				loc, _ := time.LoadLocation("Asia/Shanghai")
				fmt.Println(msg, time.Now().In(loc).Format(timeFormatStr))
				time.Sleep(frequency)
			}
		}()
	}
	wg.Wait()
}
