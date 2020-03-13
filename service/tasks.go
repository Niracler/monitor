package service

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"gamenews.niracler.com/monitor/model"
	"gamenews.niracler.com/monitor/util"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

func UpdateUserOperation() {

	log.Printf("开始更新用户行为数据\n")

	redisPool := RedisPool()
	// db := GetDB()

	// 用户行为的记录审计数据格式(放在redis的列表中)
	// IP 用户ID 资源类型 资源ID 使用设备 时间
	// redis键格式: "action_" + uid
	prefix := "action_*"
	res, _ := redisPool.Cmd("scan", "0", "MATCH", prefix, "COUNT", "10000").Array()
	keys, _ := res[1].Array()

	log.Println(keys)

	var uoList []interface{}
	for _, key := range keys {
		keyStr, _ := key.Str()
		logs, _ := redisPool.Cmd("LRANGE", keyStr, "0", "-1").Array()
		num, _ := redisPool.Cmd("DEL", keyStr).Int64()

		log.Printf("拿出 %s 中的数据条数:%d, 并删除 %d\n", keyStr, len(logs), num)

		for _, clog := range logs {
			logStr, _ := clog.Str()
			uo := util.CutLogFetchData(logStr) //将内容装到对象中
			uo.Uid, _ = strconv.ParseInt(strings.Split(keyStr, "_")[1], 10, 64)
			uoList = append(uoList, uo)
		}
	}

	err := gormbulk.BulkInsert(db, uoList, 3000)
	if err != nil {
		log.Println(err)
	}

	log.Printf("成功插入%d条数据\n", len(uoList))
	log.Printf("更新用户行为数据结束\n")
}

func UpdatePVUV() {

	log.Printf("开始更新pvuv数据\n")
	redisPool := RedisPool()

	var vcList []interface{}

	// PV,UV数据格式(放在多个redis的有序集合中)
	// 统计类型 资源类型 时间类型 时间 资源ID 点击量
	// redis键格式: anyType + resType + timeType + timestamp
	for _, anyType := range []string{"pv", "uv"} {
		prefix := anyType + "_*"
		res, _ := redisPool.Cmd("scan", "0", "MATCH", prefix, "COUNT", "100000").Array()
		keys, _ := res[1].Array()

		for _, key := range keys {
			keyStr, _ := key.Str()
			items := strings.Split(keyStr, "_")
			resources, _ := redisPool.Cmd("ZRANGE", keyStr, "0", "-1").Array()
			timeStr := util.GetTime(items[3])
			theTime, _ := time.Parse("2006-01-02 15:04:05", timeStr)

			//与当前时间进行比较
			k := time.Now()
			var d time.Duration
			if items[2] == "day" {
				d, _ = time.ParseDuration("-24h")
			} else if items[2] == "hour" {
				d, _ = time.ParseDuration("-1h")
			}
			if !theTime.Before(k.Add(d)) {
				log.Println("跳过", theTime)
				continue
			}

			fmt.Printf("拿出 %s 中的数据条数: %d, 并删除\n", keyStr, len(resources))

			for _, resource := range resources {
				resID, _ := resource.Str()
				score, _ := redisPool.Cmd("ZSCORE", keyStr, resource).Int64()
				_, _ = redisPool.Cmd("ZREM", keyStr, resource).Int64()

				//将内容装到对象中
				vc := model.VisitorCount{
					VisType:   items[0],
					ResType:   items[1],
					ResId:     resID,
					TimeType:  items[2],
					TimeLocal: &theTime,
					Click:     score,
				}
				vcList = append(vcList, &vc)
			}
		}
	}

	if len(vcList) != 0 {
		db := GetDB()
		err := gormbulk.BulkInsert(db, vcList, 3000)
		if err != nil {
			log.Println(err)
		}
		log.Printf("成功插入%d条数据\n", len(vcList))
		log.Printf("更新pvuv数据结束\n")
	}
}
