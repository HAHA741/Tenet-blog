package system

import (
	"fmt"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemArticleApi struct {
}

// 添加文章
func (s *SystemArticleApi) AddArticle(c *gin.Context) {
	var article system.Article
	_ = c.ShouldBindJSON(&article)
	if err := articleService.AddArticle(article); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

// 编辑文章
func (s *SystemArticleApi) EditArticle(c *gin.Context) {
	var article system.Article
	_ = c.ShouldBindJSON(&article)
	if err := articleService.EditArticle(article); err != nil {
		global.GVA_LOG.Error("编辑失败!", zap.Error(err))
		response.FailWithMessage("编辑失败", c)
	} else {
		response.OkWithMessage("编辑成功", c)
	}
}

// 获取文章列表
func (s *SystemArticleApi) GetArticlesList(c *gin.Context) {
	var page system.Page
	// _ = c.ShouldBindJSON(&page)
	currentStr := c.Query("current")
	sizeStr := c.Query("size")
	page.Current, _ = strconv.Atoi(currentStr)
	page.Size, _ = strconv.Atoi(sizeStr)

	if err, params := articleService.GetArticlesList(page); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(params, c)
	}

}

// 删除文章
func (s *SystemArticleApi) DelArticleById(c *gin.Context) {
	type Params struct {
		Id int
	}
	var params Params

	_ = c.ShouldBindJSON(&params)
	fmt.Printf("id", params)
	if err := articleService.DelArticleById(params.Id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// 上传markdown文件解析
func (s *SystemArticleApi) UploadArticleMD(c *gin.Context) {
	file, _ := c.FormFile("file")
	if err := articleService.UploadArticleMD(file); err != nil {
		global.GVA_LOG.Error("上传失败!", zap.Error(err))
		response.FailWithMessage("上传失败", c)
	} else {
		response.OkWithMessage("上传成功", c)
	}
}

// 保存图片并返回url
func (s *SystemArticleApi) SaveArticleImg(c *gin.Context) {
	file, _ := c.FormFile("file")
	if err, url := articleService.SaveArticleImg(file); err != nil {
		global.GVA_LOG.Error("图片上传失败!", zap.Error(err))
		response.FailWithMessage("图片上传失败!", c)
	} else {
		response.OkWithMessage(url, c)
	}
}
