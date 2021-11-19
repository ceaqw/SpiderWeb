<template>
  <div :id="chartId" :style="{height:height, width:width}" />
</template>

<script>
import theme from '@/conf/theme'

let chart = null

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
            // radius: ['40%', '70%'],
            avoidLabelOverlap: false,
            startAngle: -180,
            // label: {
            //     show: false,
            //     position: 'center'
            // },
            labelLine: {
                show: false
            },
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
            type: Function
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
            this.chartDatas(options, 'pie', this.filter, this.reRender)
            this.$store.state.flushQueue[this.parentName].push(this.refreshDatas)
            chart.on('click', (params)=>{
                // 点击图标处理事件
                console.log(params)
            })
        },
        refreshDatas() {
            this.chartDatas(options, 'pie', this.filter, this.reRender)
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
