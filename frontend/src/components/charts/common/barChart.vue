<template>
  <div :style="{height:height, width:width}" />
</template>

<script>
import theme from '@/conf/theme'


let barOptions = {
    title: {
        text: '',
        left: 'center',
        textStyle: {
            color: theme.color.black
        }
    },
    tooltip: {
        trigger: 'item',
        backgroundColor: 'rgba(255,255,255,0.8)'
    },
    textStyle: {
        fontFamily: 'Microsoft YaHei',
        fontStyle: 'normal',
        fontWeight: 'normal'
    },
    legend: {
        top: 'bottom'
    },
    grid: {
        left: '3%',
        right: '4%',
        bottom: '7%',
        containLabel: true
    },
    xAxis: [{type: 'category', data: []}],
    yAxis: [{
        type: 'value',
        axisLabel: {
            formatter: function(value){
                if(value >= 1000000){
                    value = value/1000000+'M'
                } else if (value >= 1000) {
                    value = value/1000+'k'
                }
                return value
            }
        }
    }],
    series: 
    [
        {
            name: 'Undone',
            type: 'bar',
            stack: 'Task',
            emphasis: {
                focus: 'series'
            },
            // itemStyle: {
            //     color: '#5470C6'
            // },
            data: []
        },
        {
            name: 'Fail',
            type: 'bar',
            stack: 'Task',
            emphasis: {
                focus: 'series'
            },
            // itemStyle: {
            //     color: '#EE6666'
            // },
            data: []
        },
        {
            name: 'Success',
            type: 'bar',
            stack: 'Task',
            emphasis: {
                focus: 'series'
            },
            // itemStyle: {
            //     color: '#91CC75'
            // },
            data: []
        },
        {
            name: 'Total',
            type: 'bar',
            emphasis: {
                focus: 'series'
            },
            // itemStyle: {
            //     color: '#12B8C5'
            // },
            data: []
        },
        {
            name: 'Retry',
            type: 'bar',
            emphasis: {
                focus: 'series'
            },
            // itemStyle: {
            //     color: '#FAC858'
            // },
            data: []
        }
    ]
}

export default {
    props: {
        width: {
            type: String,
            default: '100%'
        },
        height: {
            type: String,
            default: '350px'
        },
        parentName: {type: String},
        chartDatas: {
            type: Function,
            required: true
        },
        filter: {
            type: Object
        }
    },
    data() {
        return {
            chart: null
        }
    },
    mounted() {
        new Promise(()=>{this.initChart()})
    },
    beforeDestroy() {
        this.destroy()
    },
    methods: {
        initChart() {
            this.destroy()
            this.chart = this.$echarts.init(this.$el)
            this.chartDatas(barOptions, 'bar', this.filter, this.reRender)
            this.$store.state.flushQueue[this.parentName].unshift(this.refreshDatas)
            this.chart.on('click', (params)=>{
                // 点击图标处理事件
                console.log(params)
            })
        },
        refreshDatas() {
            this.chartDatas(barOptions, 'bar', this.filter, this.reRender)
        },
        reRender() {
            // this.chart.resize()
            this.chart.setOption(barOptions)
        },
        destroy() {
            if (!this.chart) {
                return
            }
            this.chart.dispose()
            this.chart = null
        }
    }
}
</script>
