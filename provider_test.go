package articlep

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func connMysql() ArticleProvider {
	// 连接mysql
	DBConn, err := gorm.Open(mysql.Open("root:ZTQ4ZTBjMTViNGMzODgzODUz@tcp(120.53.243.73:13306)/services?charset=utf8"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	provider := GetArticleProvider("default")
	provider.SetDB(DBConn)
	return provider
}

func TestCreateChannel(t *testing.T) {
	provider := connMysql()
	err := provider.CreateChannel("11111111112222222222333333333344", "11111111112222222222333333333344", "0", "国内新闻")
	if err != nil {
		t.Log(err)
	}
}

func TestDeleteChannel(t *testing.T) {
	provider := connMysql()
	err := provider.DeleteChannel("11111111112222222222333333333344")
	if err != nil {
		t.Log(err)
	}
}

func TestListChannel(t *testing.T) {
	provider := connMysql()
	res, count, err := provider.ListChannel("11111111112222222222333333333344", "20", "1")
	if err != nil {
		t.Log(err)
	}
	t.Log(count)
	t.Log(res)
}

func TestCreateContent(t *testing.T) {
	provider := connMysql()
	err := provider.CreateContent("11111111112222222222333333333344", "11111111112222222222333333333344", "11111111112222222222333333333344", "八卦新闻", "八卦新闻内容")
	if err != nil {
		t.Log(err)
	}
}

func TestDeleteContent(t *testing.T) {
	provider := connMysql()
	err := provider.DeleteContent("11111111112222222222333333333344")
	if err != nil {
		t.Log(err)
	}
}

func TestListContent(t *testing.T) {
	provider := connMysql()
	res, count, err := provider.ListContent("11111111112222222222333333333344", "20", "1")
	if err != nil {
		t.Log(err)
	}
	t.Log(count)
	t.Log(res)
}

func TestViewContent(t *testing.T) {
	provider := connMysql()
	res, err := provider.ViewContent("11111111112222222222333333333344")
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
}

func TestCreateSingle(t *testing.T) {
	provider := connMysql()
	err := provider.CreateSingle("11111111112222222222333333333344", "11111111112222222222333333333344", "八卦新闻", "八卦新闻内容")
	if err != nil {
		t.Log(err)
	}
}

func TestDeleteSingle(t *testing.T) {
	provider := connMysql()
	err := provider.DeleteSingle("11111111112222222222333333333344")
	if err != nil {
		t.Log(err)
	}
}

func TestListSingle(t *testing.T) {
	provider := connMysql()
	res, count, err := provider.ListSingle("11111111112222222222333333333344", "20", "1")
	if err != nil {
		t.Log(err)
	}
	t.Log(count)
	t.Log(res)
	t.Log(len(res))
}

func TestViewSingle(t *testing.T) {
	provider := connMysql()
	res, err := provider.ViewSingle("11111111112222222222333333333344")
	if err != nil {
		t.Log(err)
	}
	t.Log(res)
}
