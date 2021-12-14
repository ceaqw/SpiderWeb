/*
 * @Date: 2021-11-29 09:13:10
 * @LastEditTime: 2021-12-14 15:59:21
 * @Author: ceaqw
 */
// project 页面 图表组数据处理模块

import theme from '@/conf/theme'
import chartApi from '../../api/chart'
import store from '@/store'
import baseConf from '@/conf/baseConf'
import { ElMessage } from 'element-plus'

function resetChartDatas(datas, type_, options, render) {
    const data = datas.pieDatas
    let pieDatas = [
        { 
            value: 0, name: 'success',
            itemStyle: {
                color: theme.color.success
            },
        },
        { 
            value: 0, name: 'fail',
            itemStyle: {
                color: theme.color.fail
            },
        },
        // {
        //     value: 0, name: 'undone',
        //     itemStyle: {
        //         color: theme.color.undone
        //     },
        // }
    ]
    for (const index in data) {
        if (index == 2) break
        pieDatas[index].value = data[index]
    }
    options.series[0].radius = '75%'
    options.legend = {
        orient: 'vertical',
        left: 'left'
    },
    options.series[0].minAngle = baseConf.minAngle
    options.series[0].data = pieDatas
    options.title.text = 'analyse-job'
    // 异步渲染
    render()
}

function getDataMiddleware(api, options, type_, filter, render) {
    // 先入当前页面任务队列
    if (store.state.pageTaskQueue[filter.project]) {
        store.state.pageTaskQueue[filter.project].push({options, type_, render})
    } else {
        store.state.pageTaskQueue[filter.project] = [{options, type_, render}]
    }
    // 队列满，符合发送任务条件
    if (store.state.pageTaskQueue[filter.project].length == 2) {
        chartApi[api](filter).then((result)=> {
            for (let chart of store.state.pageTaskQueue[filter.project]) {
                resetChartDatas(result.data, chart.type_, chart.options, chart.render)
            }
            store.state.pageTaskQueue[filter.project] = null
            ElMessage({
                message: filter.project + '-同步完成',
                type: 'success'
            })
        }).catch((err)=> {
            console.log(err)
        })
    }
}

const projectChartDatas = (options, type_, filter, render) => {
    let tmpFilter = Object.assign({}, filter)
    tmpFilter.project = type_.split('-')[2]
    getDataMiddleware('project', options, type_, tmpFilter, render)
}


export {
    projectChartDatas,
}