<template>
  <div :id="chartId" :style="{height:height, width:width}" />
</template>

<script>
import theme from '@/conf/theme'

// chart不由组件托管，防止chart属性冲突
let chart = null

let options = {
    title: {
        text: '',
        left: 'left',
        textStyle: {
            color: theme.color.black
        }
    },
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
            name: 'Undone',
            type: 'bar',
            stack: 'Task',
            emphasis: {
                focus: 'series'
            },
            itemStyle: {
                color: '#FAC858'
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
            itemStyle: {
                color: '#5470C6'
            },
            data: []
        },
        {
            name: 'Success',
            type: 'bar',
            stack: 'Task',
            emphasis: {
                focus: 'series'
            },
            itemStyle: {
                color: '#91CC75'
            },
            data: []
        }
    ]
}

export default {
    props: {
        chartId: {
            type: String,
            required: true,
        },
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
            chart = this.$echarts.init(document.getElementById(this.chartId))
            this.chartDatas(options, 'bar', this.filter, this.reRender)
            this.$store.state.flushQueue[this.parentName].push(this.refreshDatas)
            chart.on('click', (params)=>{
                // 点击图标处理事件
                console.log(params)
            })
        },
        refreshDatas() {
            this.chartDatas(options, 'bar', this.filter, this.reRender)
        },
        reRender() {
            chart.resize()
            chart.setOption(options)
        },
        destroy() {
            if (!chart) {
                return
            }
            chart.dispose()
            chart = null
        }
    }
}
</script>
