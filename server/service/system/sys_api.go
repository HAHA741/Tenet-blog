package system

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateApi
//@description: 新增基础api
//@param: api model.SysApi
//@return: err error

type ApiService struct{}

var ApiServiceApp = new(ApiService)

func (apiService *ApiService) CreateApi(api system.SysApi) (err error) {
	if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.GVA_DB.Create(&api).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApi
//@description: 删除基础api
//@param: api model.SysApi
//@return: err error

func (apiService *ApiService) DeleteApi(api system.SysApi) (err error) {
	err = global.GVA_DB.Delete(&api).Error
	CasbinServiceApp.ClearCasbin(1, api.Path, api.Method)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAPIInfoList
//@description: 分页获取数据,
//@param: api model.SysApi, info request.PageInfo, order string, desc bool
//@return: err error

func (apiService *ApiService) GetAPIInfoList(api system.SysApi, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysApi{})
	var apiList []system.SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			// 设置有效排序key 防止sql注入
			// 感谢 Tom4t0 提交漏洞信息
			orderMap := make(map[string]bool, 5)
			orderMap["id"] = true
			orderMap["path"] = true
			orderMap["api_group"] = true
			orderMap["description"] = true
			orderMap["method"] = true
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order
				}
			}

			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return err, apiList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllApis
//@description: 获取所有的api
//@return: err error, apis []model.SysApi

func (apiService *ApiService) GetAllApis() (err error, apis []system.SysApi) {
	err = global.GVA_DB.Find(&apis).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetApiById
//@description: 根据id获取api
//@param: id float64
//@return: err error, api model.SysApi

func (apiService *ApiService) GetApiById(id float64) (err error, api system.SysApi) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateApi
//@description: 根据id更新api
//@param: api model.SysApi
//@return: err error

func (apiService *ApiService) UpdateApi(api system.SysApi) (err error) {
	var oldA system.SysApi
	err = global.GVA_DB.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = CasbinServiceApp.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = global.GVA_DB.Save(&api).Error
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApis
//@description: 删除选中API
//@param: apis []model.SysApi
//@return: err error

func (apiService *ApiService) DeleteApisByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]system.SysApi{}, "id in ?", ids.Ids).Error
	return err
}

func (apiService *ApiService) DeleteApiByIds(ids []string) (err error) {
	return global.GVA_DB.Delete(&system.SysApi{}, "id in ?", ids).Error
}

// //保存自定义布局api
func (apiService *ApiService) AddCustom(custom system.SysCustom) (err error) {
	// fmt.Println("custom", custom)
	layoutType := reflect.TypeOf(custom.CustomLayout)
	fmt.Println("变量类型", layoutType)
	custom.CustomLayout = "1"
	sqlWords := fmt.Sprintf("Insert Into sys_custom (customName,customLayout,permissionId) values ('%s','%s','%s')", custom.CustomName, custom.CustomLayout, custom.PermissionId)
	fmt.Println("sqlWords:", sqlWords)
	result := global.GVA_DB.Exec(sqlWords)
	// fmt.Println("错误", layoutType)
	return result.Error
}

// 查询子定义布局
func (apiService *ApiService) GetCustomList() (err error, customList []system.SysCustom) {
	sqlWords := fmt.Sprintf("Select * From sys_custom")
	err = global.GVA_DB.Raw(sqlWords).Scan(&customList).Error
	// fmt.Println("自定义布局列表", customList)

	// db := global.GVA_DB.Model(&system.SysCustom{})
	// err = db.Table("sys_custom").Find(&customList).Error
	fmt.Printf("自定义布局列表: %+v\n", customList)
	return
}

// 编辑自定义布局
func (apiService *ApiService) EditCustom(custom system.SysCustom) (err error) {
	sqlWords := fmt.Sprintf("Update sys_custom Set CustomName='%s' , PermissionId='%s' Where customId = '%s'", custom.CustomName, custom.PermissionId, custom.CustomId)
	err = global.GVA_DB.Exec(sqlWords).Error
	return err
}

// 删除自定义布局
func (apiService *ApiService) DelCustom(customId string) (err error) {
	sqlWords := fmt.Sprintf("Delete From sys_custom Where customId= '%s'", customId)
	err = global.GVA_DB.Exec(sqlWords).Error
	return err

}

// 获取自定义布局详情
func (apiService *ApiService) GetCustomDetail(customId string) (err error, componentList []system.CustomLayout) {
	sqlWords := fmt.Sprintf("Select * From custom_layout Where customId = %s", customId)
	err = global.GVA_DB.Raw(sqlWords).Scan(&componentList).Error
	return

}

// 保存自定义布局详情
func (apiService *ApiService) SaveCustomDetail(customId string, componentList []system.CustomLayout) (err error) {
	var customLayout system.CustomLayout
	for _, value := range componentList {
		value.CustomId = customId
		fmt.Printf("value-----: %+v\n", value)
		if err := global.GVA_DB.Table("custom_layout").Where("customId = ?", customId).Where("i=?", value.I).First(&customLayout).Error; err != nil {
			if err := global.GVA_DB.Table("custom_layout").Create(value).Error; err != nil {
				return err
			}
		} else {
			if err := global.GVA_DB.Table("custom_layout").Model(&customLayout).Where("customId = ?", customId).Where("i=?", value.I).Updates(value).Error; err != nil {
				return err
			}
		}

	}
	// if err := global.GVA_DB.Table("custom_layout").First(&customLayout).Error; err != nil {
	// 	return err
	// }
	// customLayout = componentList
	// if err := global.GVA_DB.Save(&customLayout).Error; err != nil {
	// 	return err
	// }
	return nil
}
