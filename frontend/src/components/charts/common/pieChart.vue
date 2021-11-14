<template>
  <div :class="className" :style="{height:height, width:width}" />
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
            this.chartDatas(options, 'pie', this.filter, this.reRender)
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
            this.chartDatas(options, 'pie', this.filter, this.reRender)
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
