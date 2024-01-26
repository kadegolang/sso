package common

type Meta struct {
	Id         int64 `json:"id" gorm:"column:id"`
	CreateTime int64 `json:"create_time" gorm:"column:create_time;comment:创建时间"` //;comment:创建时间
	UpdateTime int64 `json:"update_time" gorm:"column:update_time;comment:更新时间"` //;comment:更新时间
	//Label 没法存入数据库，不是一个结构化的数据
	//比如就存储在数据里面，存储为json，需要0RM来為我们完成 json的序列化和存储
	//直接序列化为json 存储到 Lable 字段
	Label map[string]string `json:"label" gorm:"column:label;serializer:json"`
}
