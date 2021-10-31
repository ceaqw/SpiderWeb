// 路由权限配置表

// 权限映射 1=>admin 2=>consumer 数值越低权限越大
const sideItem = {
    routers: [
        {path: 'dashboard', auth: 2, icon: 'el-icon-s-home'},
        {path: 'project', auth: 2, icon: 'el-icon-s-unfold'},
        {path: 'group', name: 'trend', auth: 2, icon: 'el-icon-s-order', routes: [
            {path: 'kpi', auth: 2, icon: 'el-icon-s-data'},
            {path: 'fail', auth: 2, icon: 'el-icon-s-release'},
        ]},
        {path: 'group', name: 'member', auth: 2, icon: 'el-icon-user-solid', routes: [
            {path: 'list', auth: 1, icon: 'el-icon-s-grid'},
            {path: 'create', auth: 1, icon: 'el-icon-circle-plus'},
            {path: 'update', auth: 1, icon: 'el-icon-edit'},
            {path: 'project', auth: 2, icon: 'el-icon-s-unfold'},
        ]},
    ]
}

export {
    sideItem,
}