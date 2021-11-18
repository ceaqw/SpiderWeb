<template>
    <div id="filter">
        <el-row>
            <el-col id="times" :span="19">
                <el-button-group>
                    <el-button v-for="(rate, index) in rateSelect"
                        :key="index"
                        :id="index == filterForm.dateRangeType ? 'primary' : ''"
                        size="medium"
                        @click="changeFilter('rate', index)" plain>
                        {{ rate }}
                    </el-button>
                </el-button-group>
                <span id="filter-shareing">
                    <el-switch v-model="$store.state.FilterSharing" active-text="共享筛选条件" @change="changeFilter('filter', 1)"></el-switch>
                </span>
            </el-col>
        </el-row>
        <el-row>
            <el-collapse accordion style="width: 100%">
                <el-collapse-item title="更多筛选" name="1">
                    <el-row>
                        <el-col :span="5">
                            <div class="demonstration">Start Date</div>
                            <el-date-picker v-model="filterForm.startDate" type="date" placeholder="startDate" value-format="YYYY-MM-DD"></el-date-picker>
                        </el-col>
                        <el-col :span="4">
                            <div class="demonstration">End Date</div>
                             <el-date-picker v-model="filterForm.endDate" type="date" placeholder="endDate" value-format="YYYY-MM-DD"></el-date-picker>
                        </el-col>
                        <el-col :span="5">
                            <div class="demonstration">平台</div>
                            <el-select v-model="filterForm.platForm" @focus="freshFilterData" @change="changeFilter('platform', 1)">
                                <el-option v-for="(item, index) in platFormList" :key="index" :label="item" :value="item">
                                </el-option>
                            </el-select>
                        </el-col>
                        <el-col :span="4">
                            <div class="demonstration">项目</div>
                            <el-select v-model="filterForm.project" @focus="freshFilterData">
                                <el-option v-for="(item, index) in projectList" :key="index" :label="item" :value="item">
                                </el-option>
                            </el-select>
                        </el-col>
                        <el-col :span="3">
                            <div class="demonstration">显示类型</div>
                            <el-button-group>
                                <el-button v-for="(showType, index) in showTypeSelect"
                                    :key="index"
                                    :id="index == filterForm.showType ? 'primary' : ''"
                                    @click="changeFilter('showType', index)">
                                    {{ showType }}
                                </el-button>
                            </el-button-group>
                        </el-col>
                        <el-col :span="2">
                            <div class="demonstration">-</div>
                            <el-button
                                type="primary"
                                @click="search">
                                搜索</el-button>
                        </el-col>
                    </el-row>
                </el-collapse-item>
            </el-collapse>
        </el-row>
    </div>
</template>

<script>
import { getFilterData } from '@/service/datas/base.js'
import { ElMessage } from 'element-plus'    
export default {
    props: {
        filter: {type: Object},
        parentName: {type: String}
    },
    data () {
        return {
            rateSelect: ['7天', '3天', '2天', '当天'],
            showTypeSelect: ['天', '小时'],
            filterForm: this.filter,
            platFormList: ['all'],
            projectList: ['all'],
        }    
    },
    beforeMount() {
        getFilterData()
    },
    watch: {
        'filterForm.startDate': {
            handler () {
                this.verifyDataRange()
            }
        },
        'filterForm.endDate': {
            handler () {
                this.verifyDataRange()
            }
        }
    },
    methods: {
        changeFilter(type, index) {
            if (type == 'rate') {
                this.filterForm.dateRangeType = index
                this.$store.state.shareFilter.dateRangeType = index
            } else if (type == 'showType') {
                this.filterForm.showType = index
                this.$store.state.shareFilter.showType = index
                this.verifyDataRange()
            } else if (type == 'platform') {
                this.filterForm.project = 'all'
            } else if (type == 'filter') {
                // 重置共享筛选
                this.$store.state.shareFilter.dateRangeType = 3
                this.$store.state.shareFilter.startDate = ''
                this.$store.state.shareFilter.endDate = ''
                this.$store.state.shareFilter.platForm = 'all'
                this.$store.state.shareFilter.project = 'all'
                this.$store.state.shareFilter.showType = 1
                this.filterForm = this.$store.state.FilterSharing ? this.$store.state.shareFilter : this.filter
            }
        },
        search() {
            // alert(this.$store.state.filterForm)
            const flushQueue = this.$store.state.flushQueue[this.parentName]
            // 异步更新
            for (const index in flushQueue) {
                new Promise(() => {
                    flushQueue[index]()
                })
            }
        },
        freshFilterData() {
            this.platFormList = Object.keys(this.$store.state.projectList)
            this.projectList = this.$store.state.projectList[this.filterForm.platForm]
        },
        verifyDataRange() {
            if (this.filterForm.startDate && this.filterForm.endDate) {
                if ((new Date(this.filterForm.endDate) - new Date(this.filterForm.startDate))/(24*3600*1000) > this.$conf.showTypeLimitCount) {
                    ElMessage({
                        message: '时间跨度较大，已禁止小时单位',
                        type: 'warning'
                    })
                    this.filterForm.dateRangeType = 0
                    this.$store.state.shareFilter.showType = 0
                }
            }
        }
    }
}
</script>

<style>

#filter {
    padding-bottom: 10px;
}

#filter #times {
    text-align: left;
}

#filter #times .el-button {
    background-color: #F0F2F5;
    border: 0px;
    color: black;
    padding-bottom: 5px;
}

#filter #primary {
    background-color: #409EFF !important;
    color: black;
}

#filter .demonstration {
    font-size: 130%;
}

.el-collapse .el-collapse-item, 
.el-collapse-item .el-collapse-item__header, 
.el-collapse-item .el-collapse-item__wrap {
    background-color: #E9EEF3 !important
}

#filter #filter-shareing {
    padding-left: 20px;
}
</style>