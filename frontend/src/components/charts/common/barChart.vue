<template> 
    <div style="width: 100%;">
        <div :id="Id" :style="{height:height, width:width}" />
        <div>
            <el-dialog title="详细数据" v-model="dialogTableVisible" destroy-on-close @opened="renderDialog">
                <div :id="'dialog-barChart-' + Id" style="height:350px;width:100%;"/>
            </el-dialog>
        </div>
    </div>
</template>

<script>
import theme from '@/conf/theme'
import { ElMessage } from 'element-plus'

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
        Id: {
            type: String,
            required: true
        },
        width: {
            type: String,
            default: '100%'
        },
        height: {
            type: String,
            default: '350px'
        },
        type_: {
            type: String,
            default: 'bar'
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
            dialogTableVisible: false,
            xKey: null,
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
            this.chart = this.$echarts.init(document.getElementById(this.Id))
            this.chartDatas(barOptions, this.type_, this.filter, this.reRender)
            this.$store.state.flushQueue[this.parentName][this.Id] = this.refreshDatas
            this.chart.on('click', (params)=>{
                const name = params.name
                if (name[name.length-1] != '日') {
                    ElMessage({
                        message: '当前已是最小单位，无法展开',
                        type: 'warning'
                    })
                } else {
                    this.xKey = params.name
                    this.dialogTableVisible = true
                }
            })
        },
        renderDialog() {
            let tmpChart = this.$echarts.init(document.getElementById('dialog-barChart-' + this.Id))
            let tmpFilter = Object.assign({}, this.filter)
            tmpFilter.showType = 1
            tmpFilter.startDate = this.xKey.slice(0, 10)
            tmpFilter.endDate = this.xKey.slice(0, 10)
            this.chartDatas(barOptions, this.type_ + '-single', tmpFilter, () => {tmpChart.setOption(barOptions)})
            // tmpChart = null
        },
        refreshDatas() {
            this.chartDatas(barOptions, this.type_, this.filter, this.reRender)
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
