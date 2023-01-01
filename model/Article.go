/* 数据库 Article的model */
package model

import (
	"GinBlog/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title      string   `gorm: "type:varchar(100);not null" json: "title"`
	Desc       string   `gorm: "type:varchar(200)" json: "desc"`
	Content    string   `gorm: "type:longtext" json: "content"`
	Img        string   `gorm: "type:varchar(100)" json: "img"`
	CategoryId int      `gorm: "type:int;not null" json: "category_id"`
	Category   Category `gorm: "foreignkey:CategoryId;  references:Id"` //外键为CategoryId 引用为Category中的Id
}

// 新增文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error // 新增文章
	if err != nil {
		return errmsg.ERROR // 500 创建失败返回ERROR代码
	}
	return errmsg.SUCCESS
}

// 查询分类下所有文章
func GetArtInCate(cid, pageSize, pageNum int) (int, []Article) {
	var artCateList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("category_id = ?", cid).Find(&artCateList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errmsg.ERROR, nil
	}
	if len(artCateList) <= 0 {
		return errmsg.ERROR_CATEGORYNAME_NOT_EXIST, nil
	}
	return errmsg.SUCCESS, artCateList
}

// 查询单个文章
func GetArtInfo(id int) (int, Article) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return errmsg.ERROR_ARTICLE_NOT_EXIST, Article{}
	}
	return errmsg.SUCCESS, art
}

// 查询文章列表 pageSize--每页大小 pageNum--当前页号
func GetArtList(pageSize, pageNum int) (int, []Article) {
	var arts []Article
	// 查询Article时 预加载相关的Category
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arts).Error
	// 查找错误 && 没有找到
	if err != nil && err != gorm.ErrRecordNotFound {
		return errmsg.ERROR, nil
	}
	return errmsg.SUCCESS, arts
}

// 删除文章
func DeleteArticle(id int) int {
	var art Article
	err := db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑文章信息
func EditArticle(id int, data *Article) int {
	var art Article
	mp := map[string]interface{}{}
	mp["title"] = data.Title
	mp["category_id"] = data.CategoryId
	mp["desc"] = data.Desc
	mp["content"] = data.Content
	mp["img"] = data.Img
	err := db.Model(&art).Where("id = ?", id).Updates(mp).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
