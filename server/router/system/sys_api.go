package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("api")
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		apiRouter.POST("createApi", apiRouterApi.CreateApi)               // 创建Api
		apiRouter.POST("deleteApi", apiRouterApi.DeleteApi)               // 删除Api
		apiRouter.POST("getApiById", apiRouterApi.GetApiById)             // 获取单条Api消息
		apiRouter.POST("updateApi", apiRouterApi.UpdateApi)               // 更新api
		apiRouter.POST("addCustom", apiRouterApi.AddCustom)               //添加自定义布局
		apiRouter.POST("getCustomList", apiRouterApi.GetCustomList)       //查询自定义布局列表
		apiRouter.POST("editCustom", apiRouterApi.EditCustom)             //编辑自定义布局
		apiRouter.POST("delCustom", apiRouterApi.DelCustom)               //删除自定义布局
		apiRouter.POST("getCustomDetail", apiRouterApi.GetCustomDetail)   //获取自定义布局详情
		apiRouter.POST("saveCustomDetail", apiRouterApi.SaveCustomDetail) //保存自定义布局详情
		apiRouter.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		apiRouterWithoutRecord.POST("getAllApis", apiRouterApi.GetAllApis) // 获取所有api
		apiRouterWithoutRecord.POST("getApiList", apiRouterApi.GetApiList) // 获取Api列表
	}
}
