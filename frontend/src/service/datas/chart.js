import theme from '@/conf/theme'
import chartApi from '@/api/chart'
import store from '@/store'
import { ElMessage } from 'element-plus'

function resetChartDatas(datas, type_, options, render) {
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
        options.series[0].data = pieDatas
    } else if (type_ == 'bar') {
        let barDatas = {
            xData: [],
            yData: {'undone': [], 'fail': [], 'success': []},
            yLabel: ['undone', 'fail', 'success']
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
    if (store.state.pageTaskQueue[api].length == 2) {
        chartApi[api](filter).then((result)=> {
            for (let chart of store.state.pageTaskQueue[api]) {
                console.log(render)
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

const allChartDatas = (options, type_, filter, render, chartId) => {
    getDataMiddleware('all', options, type_, filter, render, chartId)
}

const analyseDatas = (options, type_, filter, render, chartId) => {
    getDataMiddleware('analyse', options, type_, filter, render, chartId)
}

export {
    allChartDatas,
    analyseDatas
}
