package model

import "time"

type UserOperation struct {
	Model
	Uid           int64      `gorm:"index:idx_uid" json:"uid"`
	RemoteAddr    string     `gorm:"type:VARCHAR(64)" json:"remote_addr"`
	TimeLocal     *time.Time `json:"time_local"`
	HttpMethod    string     `gorm:"type:VARCHAR(32)" json:"http_method"`
	HttpUrl       string     `gorm:"type:VARCHAR(128)" json:"http_url"`
	Status        string     `gorm:"type:VARCHAR(32)" json:"status"`
	BodyBytesSent int64      `json:"bodyBytesSent"`
	HttpReferer   string     `gorm:"type:VARCHAR(128)" json:"http_referer"`
	HttpUserAgent string     `gorm:"type:VARCHAR(256)" json:"http_user_agent"`
	ResType       string     `gorm:"type:VARCHAR(64)" json:"res_type"`
	ResId         string     `gorm:"type:VARCHAR(64)" json:"res_id"`
}

// PV,UV数据格式(放在多个redis的有序集合中)
// 统计类型 资源类型 时间类型 时间 资源ID 点击量
// redis键格式: anyType + resType + timeType + timestamp
type VisitorCount struct {
	Model
	VisType   string     `gorm:"type:VARCHAR(32)" json:"vis_type"`
	ResType   string     `gorm:"type:VARCHAR(64)" json:"res_type"`
	ResId     string     `gorm:"type:VARCHAR(64)" json:"res_id"`
	TimeType  string     `gorm:"type:VARCHAR(32)" json:"time_type"`
	TimeLocal *time.Time `json:"time_local"`
	Click     int64      `gorm:"type:VARCHAR(64)" json:"click"`
}
