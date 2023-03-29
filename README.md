# articlep

## ArticleProvider 模块接口
```
type ArticleProvider interface {
	// SetDB 设置数据库连接
	SetDB(db *gorm.DB)
	// CreateChannel 创建文章分类
	CreateChannel(id, belongId, pid, name string) error
	// DeleteChannel 删除文章分类
	DeleteChannel(id string) error
	// ListChannel 获取文章分类列表
	ListChannel(belongId, pageSize, pageNum string) (string, int64, error)
	// CreateContent 创建文章
	CreateContent(id, belongId, tid, title, content string) error
	// DeleteContent 删除文章
	DeleteContent(id string) error
	// ListContent 获取文章列表
	ListContent(belongId, pageSize, pageNum string) (string, int64, error)
	// ViewContent 查看文章详情
	ViewContent(id string) (string, error)
	// CreateSingle 创建单页文章
	CreateSingle(id, belongId, title, content string) error
	// DeleteSingle 删除单页文章
	DeleteSingle(id string) error
	// ListSingle 获取单页文章列表
	ListSingle(belongId, pageSize, pageNum string) (string, int64, error)
	// ViewSingle 查看单页文章详情
	ViewSingle(id string) (string, error)
}
```

## ArticleProvider 使用示例
```
package main

import (
	"fmt"
	"github.com/provider-go/articlep"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	DBConn, err := gorm.Open(mysql.Open("root:pawword@tcp(ip:3306)/db?charset=utf8"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	provider := articlep.GetArticleProvider("default")
	provider.SetDB(DBConn)
	res, err := provider.ViewContent("11111111112222222222333333333344")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
```