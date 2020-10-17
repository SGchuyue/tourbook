package model

import (
	"github.com/jinzhu/gorm"
	"tourbook/utils/errmsg"
)

type Article struct {
	gorm.Model
	Title    string   `gorm:"type:varchar(100);not null" json:"title"` // 文章标题
	Category Category `gorm:"foreignkey:Cid"`                          // 物理外建
	Cid      int      `gorm: "type:int;not null" json:"cid"`           // 文章id
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`           // 文章说明
	Content  string   `gorm:"type:longtext" json:"content"`            // 文章内容
	Img      string   `gorm:"type:varchar(100)" json:"img"`            // 上传图像
}

// 添加文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// 查询文章列表
func GetArt(pageSize int, pageNum int) ([]Article, int) {
	var articleList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return articleList, errmsg.SUCCSE
}

// 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCSE
}

// 查询分类下的所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int) {
	var cateArtList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArtList).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_ART
	}
	return cateArtList, errmsg.SUCCSE
}

// 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["content"] = data.Content
	maps["img"] = data.Img
	maps["desc"] = data.Desc
	err := db.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除文章
func DeleteArt(id int) int {
	var art Article
	err := db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
