package system

import (
	"project/global"
	"project/model/common/request"
	"project/model/common/response"
	"project/model/system"
	systemRes "project/model/system/response"
	"project/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeptApi struct {
}

// @Tags Department
// @Summary 新增部门
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body system.SysDept true "部门名称, 是否父子级"
// @Success 200 {object} response.Response{msg=string} "新增部门"
// @Router /api/department/addDept [post]
func (d *DeptApi) AddDepartment(c *gin.Context) {
	var dept system.SysDept
	_ = c.ShouldBindJSON(&dept)
	if err := utils.Verify(dept, utils.DeptVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, dept := DeptService.AddDepartment(dept); err != nil {
		global.GSD_LOG.Error("添加部门失败", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("添加部门失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysDeptResponse{Dept: dept}, "添加成功", c)
	}
}

// @Tags Department
// @Summary 删除部门
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body system.SysDept true "删除部门"
// @Success 200 {object} response.Response{msg=string} "删除部门"
// @Router /api/department/deleteDept [post]
func (d *DeptApi) DeleteDepartment(c *gin.Context) {
	var dept system.SysDept
	_ = c.ShouldBindJSON(&dept)
	if err := utils.Verify(dept, utils.DeleteDeptVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := DeptService.DeleteDepartment(&dept); err != nil {
		global.GSD_LOG.Error("删除部门失败", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("删除部门失败", c)
	} else {
		response.OkWithMessage("删除部门成功", c)
	}
}

// @Tags Department
// @Summary 修改部门
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body system.SysDept true "部门名称, 是否父子级"
// @Success 200 {object} response.Response{msg=string} "修改部门"
// @Router /api/department/updateDept [post]
func (d *DeptApi) UpdateDepartment(c *gin.Context) {
	var dept system.SysDept
	_ = c.ShouldBindJSON(&dept)
	if err := utils.Verify(dept, utils.DeptVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, dept := DeptService.UpdateDepartment(dept); err != nil {
		global.GSD_LOG.Error("更新部门失败", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("更新部门失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysDeptResponse{Dept: dept}, "修改成功", c)
	}
}

// @Tags Department
// @Summary 查询部门列表
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取部门列表,返回包括列表,总数,页码,每页数量"
// @Router /api/department/lists [post]
func (d *DeptApi) GetDeptList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, deptList, total := DeptService.GetDeptList(pageInfo); err != nil {
		global.GSD_LOG.Error("获取失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     deptList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags Department
// @Summary 根据pid查询部门列表
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body request.GetById true "部门pid"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取部门列表,返回包括列表,总数,页码,每页数量"
// @Router /api/department/id [post]
func (d *DeptApi) GetDeptListById(c *gin.Context) {
	var Pid request.GetById
	_ = c.ShouldBindJSON(&Pid)
	if err := utils.Verify(Pid, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, deptList, total := DeptService.GetDeptListById(Pid.ID); err != nil {
		global.GSD_LOG.Error("获取失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  deptList,
			Total: total,
		}, "获取成功", c)
	}
}

// @Tags Department
// @Summary 获取部门下的用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/octet-stream
// @Param data body system.SysDept true "部门id"
// @Success 200
// @Router /api/department/users [post]
func (d *DeptApi) GetDeptUser(c *gin.Context) {
	var dept system.SysDept
	_ = c.ShouldBindJSON(&dept)
	if err := utils.Verify(dept, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, user := DeptService.GetUserByDeptId(dept); err != nil {
		global.GSD_LOG.Error("获取部门用户失败", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("获取部门用户失败", c)
		return
	} else {
		response.OkWithDetailed(gin.H{"userList": user}, "获取部门下用户成功", c)
	}
}
