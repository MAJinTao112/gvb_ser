package comment_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_ser"
	"gvb_server/service/redis_ser"
	"gvb_server/utils/jwts"
)

type CommentRequest struct {
	ArticleID       string `json:"article_id" binding:"required" msg:"请选择文章"`
	Content         string `json:"content" binding:"required" msg:"请选择内容"`
	ParentCommentID *uint  `json:"parent_comment_id"` //父评论id
}

func (CommentApi) CommentCreateView(c *gin.Context) {
	var cr CommentRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	//文章是否存在
	_, err = es_ser.CommDetail(cr.ArticleID)
	if err != nil {
		res.FailWithMessage("文章不存在", c)
		return
	}
	//fmt.Println(article, claims.UserID)
	//
	//fmt.Println(cr)
	//判断是否是子评论
	if cr.ParentCommentID != nil {
		var parentComment models.CommentModel
		err = global.DB.Take(&parentComment, cr.ParentCommentID).Error
		if err != nil {
			res.FailWithMessage("负评论不存在", c)
			return
		}
		if parentComment.ArticleID != cr.ArticleID {
			res.FailWithMessage("评论文章不一致", c)
			return
		}
		global.DB.Model(&parentComment).Update("comment_count", gorm.Expr("comment_count + 1"))
	}
	//添加评论
	global.DB.Create(&models.CommentModel{
		ArticleID:       cr.ArticleID,
		Content:         cr.Content,
		UserID:          claims.UserID,
		ParentCommentID: cr.ParentCommentID,
	})
	//拿到文章数，新的文章评论数存在缓存里

	//给文章评论数+1
	err = redis_ser.NewCommentCount().Set(cr.ArticleID)
	if err != nil {
		return
	}
	res.OkWithMessage("文章评论成功", c)
	fmt.Println(claims)
}
