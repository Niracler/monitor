package util

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gamenews.niracler.com/monitor/model"
	"gamenews.niracler.com/monitor/setting"
)

// 将日志文件中的时间格式化为时间戳的函数
func GetTimeStamp(logTime, timeType string) string {
	var item string

	switch timeType {
	case "day":
		item = "2006-01-02"
		break
	case "hour":
		item = "2006-01-02 15"
		break
	case "min":
		item = "2006-01-02 15:04"
		break
	}
	theTime, _ := time.Parse("02/Jan/2006:15:04:05 -0700", logTime)
	t, _ := time.Parse(item, theTime.Format(item))
	return strconv.FormatInt(t.Unix(), 10)
}

// 将一行的日志切割到结构体中
func CutLogFetchData(logStr string) *model.UserOperation {
	values := strings.Split(logStr, "\"")
	var res []string
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value != "" {
			res = append(res, value)
		}
	}
	if len(res) > 0 {
		r := strings.Split(res[3], " ")
		if len(r) < 3 {
			log.Fatalf("Some different", res[3])
			return nil
		}
		// 将数据放到 Channel
		r1, _ := regexp.Compile(setting.ResourceType)
		r2, _ := regexp.Compile("/([0-9]+)")
		resType := r1.FindString(r[1])
		if resType == "" {
			resType = "other"
		}

		resId := r2.FindString(r[1])
		if resId != "" {
			resId = resId[1:]
		} else {
			resId = "list"
		}
		theTime, _ := time.Parse("02/Jan/2006:15:04:05 -0700", res[2])
		bbs, _ := strconv.ParseInt(res[5], 10, 64)

		data := model.UserOperation{
			RemoteAddr:    res[0],
			TimeLocal:     &theTime,
			HttpMethod:    r[0],
			HttpUrl:       r[1],
			Status:        res[4],
			BodyBytesSent: bbs,
			HttpReferer:   res[6],
			HttpUserAgent: res[7],
			ResType:       resType,
			ResId:         resId,
		}

		return &data
	}

	return nil
}

func GetTime(timeStamp string) string {
	//0000-00-00 00:00:00
	sec, _ := strconv.ParseInt(timeStamp, 10, 64)
	tm := time.Unix(sec, 0)
	return tm.Format("2006-01-02 15:04:05")
}
