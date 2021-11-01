<template>
    <div id="filter">
        <el-row>
            <el-col id="times">
                <el-button-group>
                    <el-button v-for="(rate, index) in rateSelect"
                        :key="index"
                        :id="rateFalg[index]"
                        size="medium"
                        @click="changeRate(index)" plain>
                        {{ rate }}
                    </el-button>
                </el-button-group>
            </el-col>
        </el-row>
        <el-row>
            <el-collapse accordion style="width: 100%">
                <el-collapse-item title="更多筛选" name="1">
                    <el-row>
                        <el-col :span="6">
                            <div class="demonstration">Start Date</div>
                            <el-date-picker
                                v-model="value1"
                                type="daterange"
                                range-separator="至"
                                start-placeholder="开始日期"
                                end-placeholder="结束日期"
                            >
                            </el-date-picker>
                        </el-col>
                        <el-col :span="5">
                            <div class="demonstration">平台</div>
                            <el-select>
                                <el-option v-for="(item, index) in platFormList" :key="index" :label="item.value || item.name" :value="item.value">
                                </el-option>
                            </el-select>
                        </el-col>
                        <el-col :span="4">
                            <div class="demonstration">项目</div>
                            <el-select>
                                <el-option v-for="(item, index) in platFormList" :key="index" :label="item.value || item.name" :value="item.value">
                                </el-option>
                            </el-select>
                        </el-col>
                        <el-col :span="3">
                            <div class="demonstration">显示类型</div>
                            <el-button-group>
                                <el-button v-for="(showType, index) in showTypeSelect"
                                    :key="index"
                                    :id="showTypeFlag[index]"
                                    @click="changeShowType(index)">
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

export default {
    name: 'Filter',
    data () {
        return {
            rateSelect: ['7天', '3天', '昨天', '当天'],
            rateFalg: ['', '', '', 'primary'],
            showTypeSelect: ['天', '小时'],
            showTypeFlag: ['', 'primary'],
            value1: '',
            value2: '',
            platFormList: [
                {name: 'all', value: 'all'},
                {name: 'Rakuten', value: 'rakuten'},
                {name: 'Amazon', value: 'amazon'},
                {name: 'Yahoo', value: 'yahoo'},
                {name: 'Lohaco', value: 'lohaco'},
                {name: 'smRakuten', value: 'smRakuten'},
            ]
        }    
    },
    methods: {
        changeRate(index) {
            this.rateFalg[this.rateFalg.indexOf('primary')] = ''
            this.rateFalg[index] = 'primary'
            // TODO: 更新筛选范围
        },
        changeShowType(index) {
            this.showTypeFlag[this.showTypeFlag.indexOf('primary')] = ''
            this.showTypeFlag[index] = 'primary'
            // TODO: 更新时间显示类型
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
</style>