<template>
  <div :class="className" :style="{height:height, width:width}" />
</template>

<script>
import theme from '@/conf/theme'

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
        className: {
            type: String,
            default: 'chart'
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
        return {
            chart: null,
        }
    },
    mounted() {
        new Promise(()=>{this.initChart()})
        // window.addEventListener('resize', this.resize)
    },
    beforeDestroy() {
        this.destroy()
        // window.removeEventListener('resize', this.resize)
    },
    methods: {
        initChart() {
            this.destroy()
            this.chart = this.$echarts.init(this.$el)
            this.chartDatas(options, 'bar', this.filter, this.reRender)
            this.$store.state.flushQueue[this.parentName].unshift(this.refreshDatas)
            this.chart.on('click', (params)=>{
                // 点击图标处理事件
                console.log(params)
            })
        },
        resize() {
            this.chart.resize()
        },
        refreshDatas() {
            this.chartDatas(options, 'bar', this.filter, this.reRender)
        },
        reRender() {
            this.chart.setOption(options)
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
