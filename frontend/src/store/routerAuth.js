// 路由权限配置表

// 权限映射 1=>admin 2=>consumer 数值越低权限越大
const sideItem = {
    routers: {
        dashboard: {path: 'dashboard', auth: 3, icon: 'el-icon-s-home'},
        project: {path: 'project', auth: 3, icon: 'el-icon-s-unfold'},
        mongoMonitor: {path: 'group', name: 'mongoMonitor', auth: 3, icon: 'el-icon-s-order', routes: {
            do: {path: 'do', auth: 3, icon: 'el-icon-s-data'},
            analyse: {path: 'analyse', auth: 3, icon: 'el-icon-s-release'},
        }},
        member: {path: 'group', name: 'member', auth: 2, icon: 'el-icon-user-solid', routes: {
            list: {path: 'list', auth: 1, icon: 'el-icon-s-grid'},
            create: {path: 'create', auth: 1, icon: 'el-icon-circle-plus'},
            profile: {path: 'profile', auth: 1, icon: 'el-icon-edit'},
            project: {path: 'project', auth: 2, icon: 'el-icon-s-unfold'},
        }},
    },
    allowRoutes:[
        'home', '404', '401', '402', '405', '500', 'login', 'logout', 'register'
    ]
}

export {
    sideItem,
}