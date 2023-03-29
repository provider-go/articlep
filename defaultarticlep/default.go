package defaultarticlep

import (
	"encoding/json"
	"gorm.io/gorm"
	"strconv"
)

type DefaultArticleProvider struct {
	MysqlDB *gorm.DB
}

// NewDefaultArticleProvider 初始化供应商
func NewDefaultArticleProvider() *DefaultArticleProvider {
	p := &DefaultArticleProvider{}
	return p
}

func (d *DefaultArticleProvider) SetDB(db *gorm.DB) {
	d.MysqlDB = db
}

func (d *DefaultArticleProvider) CreateChannel(id, belongId, pid, name string) error {
	return d.MysqlDB.Table("articlep_default_channel").Select("id", "belong_id", "pid", "name").
		Create(&ArticlePDefaultChannel{Id: id, BelongId: belongId, Pid: pid, Name: name}).Error
}

func (d *DefaultArticleProvider) DeleteChannel(id string) error {
	return d.MysqlDB.Table("articlep_default_channel").Where("id = ?", id).Delete(&ArticlePDefaultChannel{}).Error
}

func (d *DefaultArticleProvider) ListChannel(belongId, pageSize, pageNum string) (string, int64, error) {
	var rows []*ArticlePDefaultChannel
	//计算列表数量
	var count int64
	d.MysqlDB.Table("articlep_default_channel").Where("belong_id = ?", belongId).Count(&count)
	// 获取分页数据
	pageSizeInt, _ := strconv.Atoi(pageSize)
	pageNumInt, _ := strconv.Atoi(pageNum)
	if err := d.MysqlDB.Table("articlep_default_channel").Where("belong_id = ?", belongId).Order("create_time desc").Limit(pageSizeInt).Offset((pageNumInt - 1) * pageSizeInt).Find(&rows).Error; err != nil {
		return "", 0, err
	}
	rowsByte, _ := json.Marshal(rows)
	return string(rowsByte), count, nil
}

func (d *DefaultArticleProvider) CreateContent(id, belongId, tid, title, content string) error {
	return d.MysqlDB.Table("articlep_default_content").Select("id", "belong_id", "tid", "title", "content").
		Create(&ArticlePDefaultContent{Id: id, BelongId: belongId, Tid: tid, Title: title, Content: content}).Error
}

func (d *DefaultArticleProvider) DeleteContent(id string) error {
	return d.MysqlDB.Table("articlep_default_content").Where("id = ?", id).Delete(&ArticlePDefaultContent{}).Error
}

func (d *DefaultArticleProvider) ListContent(belongId, pageSize, pageNum string) (string, int64, error) {
	var rows []*ArticlePDefaultContent
	//计算列表数量
	var count int64
	d.MysqlDB.Table("articlep_default_content").Where("belong_id = ?", belongId).Count(&count)
	// 获取分页数据
	pageSizeInt, _ := strconv.Atoi(pageSize)
	pageNumInt, _ := strconv.Atoi(pageNum)
	if err := d.MysqlDB.Table("articlep_default_content").Where("belong_id = ?", belongId).Order("create_time desc").Limit(pageSizeInt).Offset((pageNumInt - 1) * pageSizeInt).Find(&rows).Error; err != nil {
		return "", 0, err
	}
	rowsByte, _ := json.Marshal(rows)
	return string(rowsByte), count, nil
}

func (d *DefaultArticleProvider) ViewContent(id string) (string, error) {
	row := new(ArticlePDefaultContent)
	if err := d.MysqlDB.Table("articlep_default_content").Where("id = ?", id).First(&row).Error; err != nil {
		return "", err
	}
	rowByte, _ := json.Marshal(row)
	return string(rowByte), nil
}

func (d *DefaultArticleProvider) CreateSingle(id, belongId, title, content string) error {
	return d.MysqlDB.Table("articlep_default_single").Select("id", "belong_id", "title", "content").
		Create(&ArticlePDefaultSingle{Id: id, BelongId: belongId, Title: title, Content: content}).Error
}

func (d *DefaultArticleProvider) DeleteSingle(id string) error {
	return d.MysqlDB.Table("articlep_default_single").Where("id = ?", id).Delete(&ArticlePDefaultSingle{}).Error
}

func (d *DefaultArticleProvider) ListSingle(belongId, pageSize, pageNum string) (string, int64, error) {
	var rows []*ArticlePDefaultSingle
	//计算列表数量
	var count int64
	d.MysqlDB.Table("articlep_default_single").Where("belong_id = ?", belongId).Count(&count)
	// 获取分页数据
	pageSizeInt, _ := strconv.Atoi(pageSize)
	pageNumInt, _ := strconv.Atoi(pageNum)
	if err := d.MysqlDB.Table("articlep_default_single").Where("belong_id = ?", belongId).Order("create_time desc").Limit(pageSizeInt).Offset((pageNumInt - 1) * pageSizeInt).Find(&rows).Error; err != nil {
		return "", 0, err
	}
	rowsByte, _ := json.Marshal(rows)
	return string(rowsByte), count, nil
}

func (d *DefaultArticleProvider) ViewSingle(id string) (string, error) {
	row := new(ArticlePDefaultSingle)
	if err := d.MysqlDB.Table("articlep_default_single").Where("id = ?", id).First(&row).Error; err != nil {
		return "", err
	}
	rowByte, _ := json.Marshal(row)
	return string(rowByte), nil
}
