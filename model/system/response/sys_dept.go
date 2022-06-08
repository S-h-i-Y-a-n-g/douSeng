package response

import "project/model/system"

type SysDeptResponse struct {
	Dept system.SysDept `json:"Dept"`
}

type GetUserByDeptId struct {
	Id       uint   `json:"dept_id"`
	DeptName string `json:"dept_name"`
	Count    int    `json:"count"`
}
