// 路由权限配置表

// 权限映射 1=>admin 2=>consumer 多个权限值为各权限值的和
const sideItem = {
    routers: [
        {path: 'dashboard', auth: 3, icon: 'el-icon-s-home'},
        {path: 'project', auth: 3, icon: 'el-icon-s-unfold'},
        {path: 'group', name: 'trend', auth: 3, icon: 'el-icon-s-order', routes: [
            {path: 'kpi', icon: 'el-icon-s-data'},
            {path: 'fail', icon: 'el-icon-s-release'},
        ]},
        {path: 'group', name: 'member', auth: 3, icon: 'el-icon-user-solid', routes: [
            {path: 'list', icon: 'el-icon-s-grid'},
            {path: 'create', icon: 'el-icon-circle-plus'},
            {path: 'update', icon: 'el-icon-edit'},
            {path: 'project', icon: 'el-icon-s-unfold'},
        ]},
    ]
}

export {
    sideItem,
}