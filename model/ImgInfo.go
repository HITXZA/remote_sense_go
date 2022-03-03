package model

import "time"

//定义图像基本数据的结构体
type ImageInfo struct {
	Id         int   `db:"id"`
	Band       string    `db:"band"`  //共有多少波段
	//ThingType  string  `db:"thingtype"`  //地物类型
	//Angle      int64 `db:"angle"`  //角度就4个 不用写了
	Uuid       string `  db:"uuid"`   //唯一标识符   这个应该是前面返回给我的mapReduce完了之后 或者直接我后台生成返给他们 他们自己存着 这样方便他们查询

	// 时间
	CreateTime   time.Time `db:"create_time"`
	//ViewCount uint32 `db:"view_count"`
}

