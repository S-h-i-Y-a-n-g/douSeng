package system

import (
	v1 "project/api"

	"github.com/gin-gonic/gin"
)

type DeptRouter struct {
}

func (s *DeptRouter) InitDeptRouter(Router *gin.RouterGroup) {
	deptRouter := Router.Group("department")
	deptApi := v1.ApiGroupApp.SystemApiGroup.DeptApi
	{
		deptRouter.POST("addDept", deptApi.AddDepartment)       //添加部门
		deptRouter.POST("deleteDept", deptApi.DeleteDepartment) //删除部门
		deptRouter.POST("updateDept", deptApi.UpdateDepartment) //修改部门
		deptRouter.POST("lists", deptApi.GetDeptList)           //分页获取部门列表
		deptRouter.POST("users", deptApi.GetDeptUser)           //根据部门id获取用户
		deptRouter.POST("id", deptApi.GetDeptListById)          //根据pid获取部门列表
	}
}
