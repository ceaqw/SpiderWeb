<template>
  <div :style="{height:height, width:width}" />
</template>

<script>
import theme from '@/conf/theme'

let options = {
    title: {
        text: '',
        left: 'center',
        textStyle: {
            color: theme.color.black
        }
    },
    tooltip: {
        trigger: 'item',
        backgroundColor: 'rgba(255,255,255,0.8)',
        // formatter: '{b}: {c} ({d}%)',
    },
    legend: {
        orient: 'vertical',
        left: 'left'
    },
    series: [
        {
            type: 'pie',
            // radius: ['40%', '65%'],
            avoidLabelOverlap: false,
            startAngle: -180,
            radius: '60%',
            labelLine: {
                show: true
            },
            label: {
                // position: 'inside',
                fontSize: 14,
                formatter: '{d}%',
            },
            emphasis: {
                label: {
                    show: true,
                    fontSize: '40',
                    fontWeight: 'bold'
                }
            },
            minAngle: 15,
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
            type: Function
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
    },
    beforeDestroy() {
        this.destroy()
    },
    methods: {
        initChart() {
            this.destroy()
            this.chart = this.$echarts.init(this.$el)
            this.chartDatas(options, 'pie', this.filter, this.reRender)
            this.$store.state.flushQueue[this.parentName].unshift(this.refreshDatas)
            this.chart.on('click', (params)=>{
                // 点击图标处理事件
                console.log(params)
            })
        },
        refreshDatas() {
            this.chartDatas(options, 'pie', this.filter, this.reRender)
        },
        reRender() {
            // this.chart.resize()
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
