<template>
  <div :style="{height:height, width:width}" />
</template>

<script>
import theme from '@/conf/theme'
import baseConf from '@/conf/baseConf'

let pieOptions = {
    title: {
        text: '',
        left: 'center',
        textStyle: {
            color: theme.color.black
        }
    },
    textStyle: {
        fontFamily: 'Microsoft YaHei',
        fontStyle: 'normal',
        fontWeight: 'normal'
    },
    tooltip: {
        trigger: 'item',
        backgroundColor: 'rgba(255,255,255,0.8)',
        // formatter: '{b}: {c} ({d}%)',
    },
    series: [
        {
            type: 'pie',
            avoidLabelOverlap: false,
            startAngle: -180,
            radius: '75%',
            label: {
                position: 'inside',
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
            minAngle: baseConf.minAngle,
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
        chartDatas: {type: Function},
        cacheChartKey: {type: String},
        filter: {type: Object}
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
            // 外部托管
            if (this.chartDatas) {
                this.chartDatas(pieOptions, 'pie', this.filter, this.reRender)
                this.$store.state.flushQueue[this.parentName].unshift(this.refreshDatas)
            } else {
                // 缓存托管
                this.$store.state.cacheChart[this.cacheChartKey] = {
                    options: pieOptions, 
                    type_: 'pie',
                    render: this.reRender
                }
            }
            this.chart.on('click', (params)=>{
                // 点击图标处理事件
                console.log(params)
            })
        },
        refreshDatas() {
            this.chartDatas(pieOptions, 'pie', this.filter, this.reRender)
        },
        reRender() {
            // this.chart.resize()
            this.chart.setOption(pieOptions)
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
