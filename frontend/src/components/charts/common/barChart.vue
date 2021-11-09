<template>
  <div :class="className" :style="{height:height, width:width}" />
</template>

<script>
import * as echarts from 'echarts'
// import theme from '@/conf/theme'

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
        options: {
            type: Object
        },
    },
    data() {
        return {
            chart: null,
        }
    },
    mounted() {
        this.$nextTick(() => {
            this.initChart()
            window.addEventListener('resize', this.resize)
        })
    },
    beforeDestroy() {
        if (!this.chart) {
            return
        }
        this.chart.dispose()
        this.chart = null
        window.removeEventListener('resize', this.resize)
    },
    methods: {
        initChart() {
            this.chart = echarts.init(this.$el)
            this.chart.setOption(this.options)
        },
        resize() {
            this.chart.resize()
        }
    }
}
</script>
