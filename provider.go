package articlep

import (
	"github.com/provider-go/articlep/defaultarticlep"
	"gorm.io/gorm"
)

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

// GetArticleProvider 获取用户模块供应商
func GetArticleProvider(name string) ArticleProvider {
	if name == "default" {
		return defaultarticlep.NewDefaultArticleProvider()
	}

	return nil
}
