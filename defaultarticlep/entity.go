package defaultarticlep

type ArticlePDefaultChannel struct {
	Id         string `json:"id"`          // 主键ID
	BelongId   string `json:"belong_id"`   // 所属
	Pid        string `json:"pid"`         // 父类ID
	Name       string `json:"name"`        // 类型名称
	Recommend  int    `json:"recommend"`   // 推荐排序
	CreateTime string `json:"create_time"` // 创建时间
	UpdateTime string `json:"update_time"` // 更新时间
}

type ArticlePDefaultContent struct {
	Id         string `json:"id"`          // 主键ID
	BelongId   string `json:"belong_id"`   // 归属
	Tid        string `json:"tid"`         // 类型ID
	Title      string `json:"title"`       // 文章标题
	Content    string `json:"content"`     // 文章内容
	Status     int    `json:"status"`      // 文章状态：0：正常；1：下架
	UpdateTime string `json:"update_time"` // 更新时间
	CreateTime string `json:"create_time"` // 创建时间
}

type ArticlePDefaultSingle struct {
	Id         string `json:"id"`          // 主键ID
	BelongId   string `json:"belong_id"`   // 所属
	Title      string `json:"title"`       // 文章标题
	Content    string `json:"content"`     // 文章内容
	Status     int    `json:"status"`      // 文章状态：0：正常；1：隐藏
	CreateTime string `json:"create_time"` // 创建时间
	UpdateTime string `json:"update_time"` // 更新时间
}
