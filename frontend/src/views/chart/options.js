// 图表option配置
import theme from '@/conf/theme'

let allPlatFormOption = {
    pieOption: {
        title: {
            text: 'Access From',
            left: 'center',
            textStyle: {
                color: theme.color.black
            }
        },
        tooltip: {
            trigger: 'item',
            backgroundColor: 'rgba(255,255,255,0.8)'
        },
        legend: {
            orient: 'vertical',
            left: 'left'
        },
        series: [
            {
                type: 'pie',
                radius: ['40%', '70%'],
                avoidLabelOverlap: false,
                startAngle: 0,
                endAngle: 180,
                label: {
                    show: false,
                    position: 'center'
                },
                labelLine: {
                    show: false
                },
                emphasis: {
                    label: {
                        show: true,
                        fontSize: '40',
                        fontWeight: 'bold'
                    }
                },
                data: []
            }
        ]
    },
    barOption: {
        tooltip: {
            trigger: 'item',
            backgroundColor: 'rgba(255,255,255,0.8)'
        },
        legend: {},
        grid: {
            left: '3%',
            right: '4%',
            bottom: '3%',
            containLabel: true
        },
        xAxis: [{type: 'category', data: []}],
        yAxis: [{type: 'value'}],
        series: 
        [
            {
                name: 'Success',
                type: 'bar',
                stack: 'Task',
                emphasis: {
                    focus: 'series'
                },
                data: []
            },
            {
                name: 'Fail',
                type: 'bar',
                stack: 'Task',
                emphasis: {
                    focus: 'series'
                },
                data: []
            },
            {
                name: 'Undone',
                type: 'bar',
                stack: 'Task',
                emphasis: {
                    focus: 'series'
                },
                data: []
            }
        ]
    }
}

export {
    allPlatFormOption,
}
