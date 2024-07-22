package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type TestApi struct {
}

func (t *TestApi) TestT(c *gin.Context) {
	response.Ok(c)

}
