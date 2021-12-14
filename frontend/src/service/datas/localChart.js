/*
 * @Date: 2021-11-22 09:49:48
 * @LastEditTime: 2021-12-14 15:18:06
 * @Author: ceaqw
 */
// 本地图表数据处理模块

import theme from '../../conf/theme'
import store from '@/store'
import baseConf from '@/conf/baseConf'

const topError = (rowData) => {
    const chartItem = ['success', 'fail', 'undone']
    const leftChart = store.state.cacheChart['topError']
    // 更新左侧饼图，默认未返回结果第一个
    leftChart.options.series[0].data.splice(0)
    for (const key of chartItem) {
        leftChart.options.series[0].data.push({
            name: key,
            value: rowData[key],
            itemStyle: {
                color: theme.color[key]
            },
        })
    }
    leftChart.options.title.text = rowData['project']
    leftChart.options.legend = {top: 'bottom'}
    leftChart.render()
}

export {
    topError
}