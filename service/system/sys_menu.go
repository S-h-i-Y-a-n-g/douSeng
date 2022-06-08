package system

import (
	"errors"
	"project/global"
	"project/model/common/request"
	"project/model/system"
	"strconv"

	"gorm.io/gorm"
)

type MenuService struct{}

//@author: [houruotong](https://github.com/Monkey-Pear)
//@function: GetInfoList
//@description: 获取路由分页
//@return: err error, list interface{}, total int64
func (menuService *MenuService) GetMenuList() (err error, list interface{}, total int64) {
	var menuList []system.SysBaseMenu
	err, treeMap := menuService.getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = menuService.getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList, total
}

//@author: [houruotong](https://github.com/Monkey-Pear)
//@function: getBaseMenuTreeMap
//@description: 获取路由总树map
//@return: err error, treeMap map[string][]model.SysBaseMenu
func (menuService *MenuService) getBaseMenuTreeMap() (err error, treeMap map[string][]system.SysBaseMenu) {
	var allMenu []system.SysBaseMenu
	treeMap = make(map[string][]system.SysBaseMenu)
	err = global.GSD_DB.Order("sort").Preload("Parameters").Find(&allMenu).Error
	for _, v := range allMenu {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//@author: [houruotong](https://github.com/Monkey-Pear)
//@function: getBaseMenuTreeMap
//@description: 获取菜单子菜单
//@param: menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu
//@return: err error
func (menuService *MenuService) getBaseChildrenList(menu *system.SysBaseMenu, treeMap map[string][]system.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@author: [houruotong](https://github.com/Monkey-Pear)
//@function: AddMenu
//@description: 菜单管理-新增菜单
//@param: menu system.SysBaseMenu
//@return: err error

func (menuService *MenuService) AddMenu(menu system.SysBaseMenu) error {
	if !errors.Is(global.GSD_DB.Where("name = ?", menu.Name).First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("已存在name,请修改name")
	}
	return global.GSD_DB.Create(&menu).Error
}

//@author: [houruotong](https://github.com/Monkey-Pear)
//@function: DeleteMenu
//@description: 菜单管理-删除菜单
//@param: id uint
//@return: err error

func (menuService *MenuService) DeleteMenu(id uint) (err error) {
	err = global.GSD_DB.Preload("Parameters").Where("parent_id = ?", id).First(&system.SysBaseMenu{}).Error
	if err != nil {
		var menu system.SysBaseMenu
		db := global.GSD_DB.Preload("SysAuthoritys").Where("id = ?", id).First(&menu).Delete(&menu)
		err = global.GSD_DB.Delete(&system.SysBaseMenuParameter{}, "sys_base_menu_id = ?", id).Error
		if err != nil {
			return err
		}
		if len(menu.SysAuthoritys) > 0 {
			err = global.GSD_DB.Model(&menu).Association("SysAuthoritys").Delete(&menu.SysAuthoritys)
		} else {
			err = db.Error
			if err != nil {
				return
			}
		}
	} else {
		return errors.New("此菜单存在子菜单，不可删除")
	}
	return err
}

//@author: [houruotong](https://github.com/Monkey-Pear)
//@function: UpdateMenu
//@description: 菜单管理-更新菜单
//@param: menu system.SysBaseMenu
//@return: err error
func (menuService *MenuService) UpdateMenu(menu system.SysBaseMenu) (err error) {
	var oldMenu system.SysBaseMenu
	updateMenu := make(map[string]interface{})
	updateMenu["keep_alive"] = menu.KeepAlive
	updateMenu["close_tab"] = menu.CloseTab
	updateMenu["default_menu"] = menu.DefaultMenu
	updateMenu["parent_id"] = menu.ParentId
	updateMenu["path"] = menu.Path
	updateMenu["name"] = menu.Name
	updateMenu["hidden"] = menu.Hidden
	updateMenu["component"] = menu.Component
	updateMenu["title"] = menu.Title
	updateMenu["icon"] = menu.Icon
	updateMenu["sort"] = menu.Sort

	err = global.GSD_DB.Transaction(func(tx *gorm.DB) error {
		find := tx.Where("id = ?", menu.ID).First(&oldMenu)
		txErr := find.Error
		if txErr != nil {
			return txErr
		}
		if oldMenu.Name != menu.Name {
			if !errors.Is(tx.Where("name = ?", menu.Name).First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
				return errors.New("存在相同name修改失败")
			}
		}
		txErr = tx.Unscoped().Delete(&system.SysBaseMenuParameter{}, "sys_base_menu_id = ?", menu.ID).Error
		if txErr != nil {
			return txErr
		}
		if len(menu.Parameters) > 0 {
			for i := range menu.Parameters {
				menu.Parameters[i].SysBaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.Parameters).Error
			if txErr != nil {
				return txErr
			}
		}
		txErr = find.Updates(updateMenu).Error
		if txErr != nil {
			return txErr
		}
		return nil
	})
	return err
}

//@author: [houruotong](https://github.com/Monkey-Pear)
//@function: GetMenuTree
//@description: 获取用户菜单
//@param: ids []system.SysAuthority
//@return: err error, menus []model.SysMenu
func (menuService *MenuService) GetUserMenu(ids []system.SysAuthority) (err error, menus []system.SysMenu) {
	err, menuTree := menuService.getMenuTree(ids)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

//@author: [houruotong](https://github.com/Monkey-Pear)
//@function: GetMenuTree
//@description: 获取用户菜单
//@param: ids []system.SysAuthority
//@return: err error, menus []model.SysMenu
func (menuService *MenuService) getMenuTree(authorities []system.SysAuthority) (err error, treeMap map[string][]system.SysMenu) {
	var authorityIDs []uint
	var allMenus []system.SysMenu
	for _, authority := range authorities {
		authorityIDs = append(authorityIDs, authority.AuthorityId)
	}
	treeMap = make(map[string][]system.SysMenu)
	err = global.GSD_DB.Where("authority_id in (?)", authorityIDs).Order("sort").Preload("Parameters").Find(&allMenus).Error
	if err != nil {
		return
	}
	//list去重
	deduplicationList := make([]system.SysMenu, 0)
	for _, menu := range allMenus {
		flag := false
		for _, sysMenu := range deduplicationList {
			if menu.ID == sysMenu.ID {
				flag = true
				break
			}
		}
		if !flag {
			deduplicationList = append(deduplicationList, menu)
		}
	}
	for _, v := range deduplicationList {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//@author: [houruotong](https://github.com/Monkey-Pear)
//@function: GetMenuTree
//@description: 获取用户菜单的子菜单
//@param: menu *system.SysMenu, treeMap map[string][]system.SysMenu
//@return: err error, menus []model.SysMenu
func (menuService *MenuService) getChildrenList(menu *system.SysMenu, treeMap map[string][]system.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@author: [houruotong](https://github.com/Monkey-Pear)
//@function: GetBaseMenuTree
//@description: 获取基础路由树
//@return: err error, menus []model.SysBaseMenu
func (menuService *MenuService) GetBaseMenuTree() (err error, menus []system.SysBaseMenu) {
	err, treeMap := menuService.getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getBaseChildrenList(&menus[i], treeMap)
	}
	return err, menus
}

//@author: [houruotong](https://github.com/Monkey-Pear)
//@function: GetMenuAuthority
//@description: 查看当前角色树
//@param: info *request.GetAuthorityId
//@return: err error, menus []model.SysMenu
func (menuService *MenuService) GetMenuAuthority(info *request.GetAuthorityId) (err error, menus []system.SysMenu) {
	err = global.GSD_DB.Where("authority_id = ? ", info.AuthorityId).Order("sort").Find(&menus).Error
	// sql := "SELECT authority_menu.keep_alive,authority_menu.default_menu,authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? ORDER BY authority_menu.sort ASC"
	// err = global.GVA_DB.Raw(sql, authorityId).Scan(&menus).Error
	return err, menus
}

//@author: [houruotong](https://github.com/Monkey-Pear)
//@function: GetBaseMenuById
//@description: 返回当前选中menu
//@param: id float64
//@return: err error, menu model.SysBaseMenu
func (menuService *MenuService) GetBaseMenuById(id uint) (err error, menu system.SysBaseMenu) {
	err = global.GSD_DB.Preload("Parameters").Where("id = ?", id).First(&menu).Error
	return
}

//@author: [houruotong](https://github.com/Monkey-Pear)
//@function: AddMenuAuthority
//@description: 为角色增加menu树
//@param: menus []model.SysBaseMenu, authorityId string
//@return: err error
func (menuService *MenuService) AddMenuAuthority(menus []system.SysBaseMenu, authorityId uint) (err error) {
	var auth system.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	err = AuthorityServiceApp.SetMenuAuthority(&auth)
	return err
}
