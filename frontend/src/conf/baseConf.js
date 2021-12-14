/*
 * @Date: 2021-11-22 09:49:48
 * @LastEditTime: 2021-12-14 14:55:06
 * @Author: ceaqw
 */
// 基本配置

const baseConf = {
    pageSize: 20,
    role: 2,
    apiVersion: '/api/v1',
    filterForm: {
        dateRangeType: 3,
        startDate: '',
        endDate: '',
        platForm: 'all',
        project: 'all',
        showType: 0
    },
    showTypeLimitIndex: 1,
    showTypeLimitCount: 3,
    minAngle: 15,
}

export default baseConf