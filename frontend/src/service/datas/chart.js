import theme from '@/conf/theme'
import chartApi from '@/api/chart'

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

const allChartDatas = (options, type_, filter, render) => {
    chartApi.all(filter, type_).then((result)=> {
        resetChartDatas(result.data, type_, options, render)
    }).catch((err)=> {
        console.log(err)
    })
}

export {
    allChartDatas
}