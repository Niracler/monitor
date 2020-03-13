package service

import "gamenews.niracler.com/monitor/model"

func GetUserOperation(pageNum int, pageSize int, maps interface{}) (count int, uoList []model.UserOperation) {
	db := GetDB()
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&uoList)
	db.Model(&model.UserOperation{}).Where(maps).Count(&count)

	return count, uoList
}

func GetVisitorCount(pageNum int, pageSize int, maps interface{}) (count int, vcList []model.VisitorCount) {
	db := GetDB()
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&vcList)
	db.Model(&model.VisitorCount{}).Where(maps).Count(&count)

	return count, vcList
}
