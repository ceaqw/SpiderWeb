import theme from '@/conf/theme'
import chartApi from '@/api/chart'
import store from '@/store'
import baseConf from '@/conf/baseConf'
import { ElMessage } from 'element-plus'
import { topError } from '../datas/localChart'

function resetChartDatas(datas, type_, options, render) {
    if (type_.indexOf('single') > -1) type_ = type_.split('-')[0]
    if (type_ == 'pie') {
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
                    color: theme.color.danger
                },
            },
            {
                value: 0, name: 'undone',
                itemStyle: {
                    color: theme.color.undone
                },
            }
        ]
        for (const index in data) {
            pieDatas[index].value = data[index]
        }
        options.series[0].radius = '75%'
        options.legend = {
            orient: 'vertical',
            left: 'left'
        },
        options.series[0].minAngle = baseConf.minAngle
        options.series[0].data = pieDatas
    } else if (type_ == 'bar') {
        let barDatas = {
            xData: [],
            yData: {'undone': [], 'fail': [], 'success': [], 'total': [], 'retry': []},
            yLabel: ['undone', 'fail', 'success', 'total', 'retry']
        }
        const data = datas.barDatas
        if (data) {
            for (const index in data) {
                barDatas.xData.push(data[index].date)
                for (const keyIndex in barDatas.yLabel) {
                    const key = barDatas.yLabel[keyIndex]
                    barDatas.yData[key].push(data[index][key])
                }
            }
        }
        options.xAxis[0].data = barDatas.xData
        for (const index in barDatas.yLabel) {
            options.series[index].data = barDatas.yData[barDatas.yLabel[index]]
        }
    }
    options.title.text = datas.title
    // 异步渲染
    render()

}

function getDataMiddleware(api, options, type_, filter, render) {
    // 先入当前页面任务队列
    if (store.state.pageTaskQueue[api]) {
        store.state.pageTaskQueue[api].push({options, type_, render})
    } else {
        store.state.pageTaskQueue[api] = [{options, type_, render}]
    }
    // 队列满，符合发送任务条件
    if (store.state.pageTaskQueue[api].length == 2 || type_.indexOf('single') > -1) {
        chartApi[api](filter).then((result)=> {
            for (let chart of store.state.pageTaskQueue[api]) {
                resetChartDatas(result.data, chart.type_, chart.options, chart.render)
            }
            store.state.pageTaskQueue[api] = null
            ElMessage({
                message: api + '-同步完成',
                type: 'success'
            })
        }).catch((err)=> {
            console.log(err)
        })
    }
}

const allChartDatas = (options, type_, filter, render) => {
    getDataMiddleware('all', options, type_, filter, render)
}

const analyseDatas = (options, type_, filter, render) => {
    getDataMiddleware('analyse', options, type_, filter, render)
}

const topErrorDatas = (options, filter, render) => {
    chartApi.topError(filter).then((result)=> {
        let tableDatas = Object.values(result.data.tableDatas)
        tableDatas.sort((a, b)=>{return b.fail - a.fail})
        render.splice(0)
        render.push(...tableDatas)
        if (render.length > 0) {
            options.data = render.slice(0, baseConf.pageSize/2)
            // 更新左侧饼图，默认未返回结果第一个
            topError(render[0])
        }
        options.totalNumber = result.data.count
        ElMessage({
            message: 'topError' + '-同步完成',
            type: 'success'
        })
    }).catch((err)=> {
        console.log(err)
    })
}

function platformDatas (platform, options, type_, filter, render) {
    let tmpFilter = Object.assign({}, filter)
    tmpFilter.platForm = platform
    tmpFilter.project = 'all'
    // 各平台bar图表高度对齐
    if (type_ == 'bar') options.grid.bottom = '13%'
    getDataMiddleware(platform, options, type_, tmpFilter, render)
}
const rakutenDatas = (options, type_, filter, render) => {
    platformDatas('rakuten', options, type_, filter, render)
}

const yahooDatas = (options, type_, filter, render) => {
    platformDatas('yahoo', options, type_, filter, render)
}

const amazonDatas = (options, type_, filter, render) => {
    platformDatas('amazon', options, type_, filter, render)
}

export {
    allChartDatas,
    analyseDatas,
    rakutenDatas,
    yahooDatas,
    amazonDatas,
    topErrorDatas,
}
