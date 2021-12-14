<template>
    <div style="width: 100%;">
        <div :id="Id" :style="{height:height, width:width}" />
        <div>
            <el-dialog :title="dialogTitle" v-model="dialogTableVisible" destroy-on-close @opened="renderDialog">
                <div :id="'dialog-pieChart-' + Id" style="height:350px;width:100%;"/>
            </el-dialog>
        </div>
    </div>
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
            default: 'pie'
        },
        parentName: {type: String},
        chartDatas: {type: Function},
        cacheChartKey: {type: String},
        filter: {type: Object},
        dialog: {type: Function}
    },
    data() {
        return {
            chart: null,
            dialogTitle: '',
            dialogTableVisible: false,
            clickParam: null,
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
            // 外部托管
            if (this.chartDatas) {
                this.chartDatas(pieOptions, this.type_, this.filter, this.reRender)
                this.$store.state.flushQueue[this.parentName][this.Id] = this.refreshDatas
            } else {
                // 缓存托管
                this.$store.state.cacheChart[this.cacheChartKey] = {
                    options: pieOptions, 
                    type_: this.type_,
                    render: this.reRender
                }
            }
            this.chart.on('click', (params)=>{
                // console.log(params)
                if (this.dialog) {
                    this.clickParam = params
                    this.dialogTableVisible = true
                }
            })
        },
        renderDialog() {
            if (this.dialog) {
                let chartDom = document.getElementById('dialog-pieChart-' + this.Id)
                this.dialog(chartDom, this.clickParam, this.filter)
            }
        },
        refreshDatas() {
            this.chartDatas(pieOptions, this.type_, this.filter, this.reRender)
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
