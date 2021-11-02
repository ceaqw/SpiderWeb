<template>
  <div :class="className" :style="{height:height, width:width}" />
</template>

<script>
import * as echarts from 'echarts'
import theme from '@/conf/theme'

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
            default: '300px'
        }
  },
  data() {
    return {
      chart: null
    }
  },
  mounted() {
    this.$nextTick(() => {
      this.initChart()
    })
  },
  beforeDestroy() {
    if (!this.chart) {
      return
    }
    this.chart.dispose()
    this.chart = null
  },
  methods: {
    initChart() {
        this.chart = echarts.init(this.$el, 'macarons')

        this.chart.setOption({
            title: {
                text: 'Access From',
                left: 'center',
                textStyle: {
                    color: theme.color.black
                }
            },
            tooltip: {
                trigger: 'item',
                backgroundColor: "rgba(255,255,255,0.8)"
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
                    startAngle: 135,
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
                    data: [
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
                }
            ]
        })
    }
  }
}
</script>
