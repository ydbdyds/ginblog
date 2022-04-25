package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Uid      uint   `json:"uid"`
	Aid      uint   `json:"aid"`
	Title    string `json:"title"`
	Username string `json:"username"`
	Content  string `gorm:"type:varchar(500);not null;" json:"content"`
	Status   int8   `gorm:"type:tinyint;default:2" json:"status"`
}

//新增评论
func AddComment(data *Comment) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//获取评论列表
func GetCommentList(pageSize int, pageNum int) ([]Comment, int64, int) {
	var commentList []Comment
	var total int64
	db.Find(&Comment{}).Count(&total)
	err = db.Model(&Comment{}).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Select("comment.id, article.title,uid,aid, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").Joins("LEFT JOIN article ON comment.aid = article.id").Joins("LEFT JOIN user ON comment.uid = USER.id").Scan(&commentList).Error
	if err != nil {
		return commentList, 0, errmsg.ERROR
	}
	return commentList, total, errmsg.SUCCESS
}

//根据文章id获取评论
func GetComment(id int, pageSize int, pageNum int) ([]Comment, int64, int) {
	var commentList []Comment
	var total int64
	db.Where("aid = ?", id).Find(&commentList).Count(&total)
	err = db.Model(&Comment{}).Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Select("comment.id, article.title,uid,aid, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").Joins("LEFT JOIN article ON comment.aid = article.id").Joins("LEFT JOIN user ON comment.uid = USER.id").
		Where(
			"aid = ?", id,
		).Scan(&commentList).Error
	if err != nil {
		return commentList, 0, errmsg.ERROR
	}
	return commentList, total, errmsg.SUCCESS
}

// 删除评论
func DeleteComment(id uint) int {
	var comment Comment
	err = db.Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 通过评论
func CheckComment(id int, data *Comment) int {
	var comment Comment
	var maps = make(map[string]interface{})
	maps["status"] = data.Status

	err = db.Model(&comment).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
