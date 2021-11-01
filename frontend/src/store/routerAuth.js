// 路由权限配置表

// 权限映射 1=>admin 2=>consumer 数值越低权限越大
const sideItem = {
    routers: {
        dashboard: {path: 'dashboard', auth: 2, icon: 'el-icon-s-home'},
        project: {path: 'project', auth: 2, icon: 'el-icon-s-unfold'},
        trend: {path: 'group', name: 'trend', auth: 2, icon: 'el-icon-s-order', routes: {
            kpi: {path: 'kpi', auth: 2, icon: 'el-icon-s-data'},
            fail: {path: 'fail', auth: 2, icon: 'el-icon-s-release'},
        }},
        member: {path: 'group', name: 'member', auth: 2, icon: 'el-icon-user-solid', routes: {
            list: {path: 'list', auth: 1, icon: 'el-icon-s-grid'},
            create: {path: 'create', auth: 1, icon: 'el-icon-circle-plus'},
            update: {path: 'update', auth: 1, icon: 'el-icon-edit'},
            project: {path: 'project', auth: 2, icon: 'el-icon-s-unfold'},
        }},
    },
    allowRoutes:[
        'home', '404', '401', '402', '405', '500'
    ]
}

export {
    sideItem,
}