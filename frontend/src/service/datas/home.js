import theme from '@/conf/theme'

const allPlatFormDatas = (options, type_) => {
    let pieDatas = [
        { 
            value: 1048, name: 'success',
            itemStyle: {
                color: theme.color.success
            },
        },
        { 
            value: 735, name: 'fail',
            itemStyle: {
                color: theme.color.danger
            },
        },
        {
            value: 580, name: 'undone',
            itemStyle: {
                color: theme.color.undone
            },
        }
    ]
    let barDatas = {
        xData: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
        yData: [
            [120, 132, 101, 134, 90, 230, 210],
            [220, 182, 191, 234, 290, 330, 310],
            [150, 232, 201, 154, 190, 330, 410]
        ]
    }
    if (type_ == 'pie') {
        options.series[0].data = pieDatas
    } else if (type_ == 'bar') {
        options.xAxis[0].data = barDatas.xData
        for (const index in barDatas.yData) {
            options.series[index].data = barDatas.yData[index]
        }
    }
}

export {
    allPlatFormDatas
}