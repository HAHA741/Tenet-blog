package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct{}

func (s *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	articleRouter := Router.Group("article").Use(middleware.OperationRecord())
	// articleRouterWithoutRecord := Router.Group("article")
	articleRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemArticleApi
	{
		articleRouter.POST("addArticle", articleRouterApi.AddArticle)           // 新建文章
		articleRouter.POST("editArticle", articleRouterApi.EditArticle)         // 编辑文章
		articleRouter.GET("getArticlesList", articleRouterApi.GetArticlesList)  // 获取文章列表
		articleRouter.POST("delArticleById", articleRouterApi.DelArticleById)   //删除文章
		articleRouter.POST("uploadArticleMD", articleRouterApi.UploadArticleMD) //上传mD文件解析
		articleRouter.POST("saveArticleImg", articleRouterApi.SaveArticleImg)   //保存图片
	}

}
