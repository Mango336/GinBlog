/*
数据库 User的model
相关的数据库操作
*/
package model

import (
	"GinBlog/utils/errmsg"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password string `gorm: "type:varchar(20);not null" json: "password"`
	Username string `gorm: "type:varchar(20);not null" json: "username"`
	Role     int    `gorm: "type: int" json: "role"`
}

// 查询用户是否存在
func CheckUser(name string) int {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 { // id > 0  该用户已经存在 放回存在的user数据
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 获取用户的所有信息
func GetUserInfo(id int) (int, User) {
	var user User
	db.First(&user, id)
	if user.ID <= 0 {
		return errmsg.ERROR, user
	}
	return errmsg.SUCCESS, user
}

// 创建用户
func CreateUser(data *User) int {
	data.BeforeSave()             // 密码加密
	err := db.Create(&data).Error // 创建用户
	if err != nil {
		return errmsg.ERROR // 500 创建失败返回ERROR代码
	}
	return errmsg.SUCCESS
}

// 查询用户列表 pageSize--每页大小 pageNum--当前页号
func GetUsers(pageSize, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	// 上面一行写法==sql语句: select * from users offset (pageNum-1)*pageSize limit pageSize;
	// 查找错误 && 没有找到
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 密码加密--使用方法
func (u *User) BeforeSave() {
	// u.Password = ScryptPwd(u.Password)
	u.Password = BcryptPwd(u.Password)
}

// 密码加密--使用scrypt包
func ScryptPwd(password string) string {
	const KeyLen = 10 // 返回生成密钥的字节片
	// 加盐 随机数 随机生成8个
	salt := make([]byte, 8)
	salt = []byte{11, 22, 33, 56, 78, 96, 123, 99}
	// N 这里取1<<14 表示的是CPU/内存成本参数
	HashPwd, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	pwd := base64.StdEncoding.EncodeToString(HashPwd)
	return pwd
}

// 密码验证--scrypt包
func IsScryptPwd(dbPwd, userPwd string) bool {
	return ScryptPwd(userPwd) == dbPwd
}

// 密码加密--bcrypt包
func BcryptPwd(password string) string {
	// GenerateFromPassword中携带两个参数--要哈希的密码、创建哈希密码的哈希成本 Default为10
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	pwd := base64.StdEncoding.EncodeToString(hashPwd)
	return pwd
}

// 密码验证--bcrypt包
func IsBcryptPwd(dbPwd, userPwd string) bool {
	// 直接将用户输入的密码（pwd）加密后与数据库中存储的密码（dbPwd）对比 会有错误
	// 需要将dbPwd base64解码 得到hashPwd
	// 再使用CompareHashAndPassword()函数来验证密码是否正确
	hashPwd, _ := base64.StdEncoding.DecodeString(dbPwd)           // 将数据库中存储的加密密码解码
	err := bcrypt.CompareHashAndPassword(hashPwd, []byte(userPwd)) // 哈希密码与用户输入密码对比
	if err != nil {                                                // 密码错误
		return false
	}
	return true
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑用户信息
func EditUser(id int, data *User) int {
	// 更新多列--两种方法map和struct
	// 1. map
	var user User
	mp := map[string]interface{}{}
	mp["username"] = data.Username
	mp["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Updates(mp).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 用户登录验证
func CheckLogin(username, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID <= 0 {
		return errmsg.ERROR_USERNAME_NOT_EXIST
	}
	if !IsBcryptPwd(user.Password, password) {
		return errmsg.ERROR_PASSWORD_WORNG
	}
	// 使用Scrypt加密时
	// if !IsScryptPwd(user.Password, password) {
	// 	return errmsg.ERROR_PASSWORD_WORNG
	// }
	if user.Role != 0 {
		return errmsg.ERROR_USER_HAVE_NO_RIGHT
	}
	return errmsg.SUCCESS
}
