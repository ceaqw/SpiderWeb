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
            <el-collapse accordion class="main-color" style="width: 100%">
                <el-collapse-item title="更多筛选" name="1">
                    <el-row>
                        <el-col :span="7">
                            <span class="demonstration">Start Date</span>
                            <el-date-picker
                                v-model="value1"
                                type="daterange"
                                range-separator="至"
                                start-placeholder="开始日期"
                                end-placeholder="结束日期"
                            >
                            </el-date-picker>
                        </el-col>
                        <el-col :span="4">
                            <span class="demonstration">平台</span>
                            <el-select v-model="componentValue" :placeholder="placeholder" :multiple="multiple" clearable 
                                :filterable="filterable" :remote="remote" :remote-method="remoteMethod" :loading="loading" @change="changeValueFunction"
                                @visible-change="visibleChangeFunction" :style="'width:' + width + '!important'">
                            <el-option v-for="(item,index) in list" :key="item.id + index" :label="item.value || item.name" :value="item.id">
                            </el-option>
                        </el-select>
                        </el-col>
                        <!-- <el-col :span="5"></el-col>
                        <el-col :span="5"></el-col>
                        <el-col :span="4"></el-col> -->
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
            value1: '',
            value2: '',
        }    
    },
    methods: {
        changeRate(index) {
            this.rateFalg[this.rateFalg.indexOf('primary')] = ''
            this.rateFalg[index] = 'primary'
            // TODO: 更新筛选范围
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

#filter #times #primary {
    background-color: #409EFF;
}

.main-color *{
    background-color: #E9EEF3 !important;
}

#filter .demonstration {
    font-size: 130%;
}
</style>