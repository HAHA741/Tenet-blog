package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type testRouter struct{}

func (s *testRouter) InitTestRouter(Router *gin.RouterGroup) {
	testRouter := Router.Group("test")
	testApi := v1.ApiGroupApp.SystemApiGroup.TestApi
	{
		testRouter.POST("testT", testApi.TestT) // 创建Api
		// initRouter.POST("checkdb", dbApi.CheckDB) // 创建Api
	}
}
