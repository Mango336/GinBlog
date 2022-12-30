/* 数据库 Category的model */
package model

import (
	"GinBlog/utils/errmsg"

	"gorm.io/gorm"
	// "fmt"
)

type Category struct {
	ID   uint   `gorm: "primary_key; auto_increment" json: "id"`
	Name string `gorm: "type:varchar(20);not null" json: "name"`
}

// 查询分类是否存在
func CheckCategory(name string) int {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 { // id > 0  该用户已经存在 放回存在的user数据
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 获取分类的所有信息
func GetCateInfo(id int) (int, Category) {
	var cate Category
	db.First(&cate, id)
	if cate.ID <= 0 {
		return errmsg.ERROR, cate
	}
	return errmsg.SUCCESS, cate
}

// 创建分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error // 创建用户
	if err != nil {
		return errmsg.ERROR // 500 创建失败返回ERROR代码
	}
	return errmsg.SUCCESS
}

// 查询分类列表 pageSize--每页大小 pageNum--当前页号
func GetCategory(pageSize, pageNum int) []Category {
	var cate []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	// 上面一行写法==sql语句: select * from categories offset (pageNum-1)*pageSize limit pageSize;
	// fmt.Printf("%v\n", users)
	// 查找错误 && 没有找到
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

// 删除用户
func DeleteCategory(id int) int {
	var cate Category
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑用户信息
func EditCategory(id int, data *Category) int {
	// 更新多列--两种方法map和struct
	// 1. map
	var cate Category
	mp := map[string]interface{}{}
	mp["name"] = data.Name
	// mp["role"] = data.Role
	err := db.Model(&cate).Where("id = ?", id).Updates(mp).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
