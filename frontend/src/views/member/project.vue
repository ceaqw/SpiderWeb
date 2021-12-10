<template>
    <div id="member-project">
        <el-row>
            <el-collapse accordion style="width: 100%">
                <el-collapse-item title="筛选" name="1">
                    <el-row>
                        <el-col :span="5">
                            <div class="demonstration">平台</div>
                            <el-select v-model="filterForm.platForm" @focus="freshFilterData" @change="filterForm.project='all'">
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
                        <el-col :span="2">
                            <div class="demonstration">-</div>
                            <el-button
                                type="primary"
                                @click="search(1)">
                                搜索</el-button>
                        </el-col>
                        <el-col :span="2">
                            <div class="demonstration">-</div>
                            <el-button
                                type="primary"
                                @click="add">
                                添加</el-button>
                        </el-col>
                    </el-row>
                </el-collapse-item>
            </el-collapse>
        </el-row>
        <el-row>
            <!-- <span class="title">详情</span> -->
            <el-divider></el-divider>
            <el-table :data="tableData.datas" stripe>
                <el-table-column prop="rank" label="Id" :formatter="(row,column,val,index)=>{return index+1}" width="40%"></el-table-column>
                <el-table-column prop="project" label="Project"></el-table-column>
                <el-table-column prop="platform" label="Platform"></el-table-column>
                <el-table-column prop="script_id" label="Script Id"></el-table-column>
                <el-table-column prop="ip" label="Ip"></el-table-column>
                <el-table-column prop="comment" label="Comment"></el-table-column>
                <el-table-column prop="bind_table" label="BindTable"></el-table-column>
                <el-table-column label="Set">
                    <template #default="scope">
                        <el-button type="primary" icon="el-icon-setting" size="medium" title="设置" @click.prevent="option(scope.row)"></el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination background layout="prev, pager, next" :total="tableData.totalNumber" @current-change="changePage" :page-size="tableData.pageSize">
            </el-pagination>
        </el-row>
        <el-dialog :title="optionType" v-model="dialogVisible" destroy-on-close>
            <!-- TODO: 更新project -->
        </el-dialog>
    </div>
</template>

<script>
import { getFilterData } from '@/service/datas/base.js'
import { getProjectInfos, getAllPlatformAndProjectMap } from '../../service/project'
import baseConf from '@/conf/baseConf'
export default {
    data() {
        return {
            filterForm: { platForm: 'all', project: 'all', page: 1 },
            platFormList: ['all'],
            projectList: ['all'],
            tableData: {
                datas: [
                    // {
                    //     platform: "shopee",
                    //     project: "item",
                    //     script_id: 7456,
                    //     ip: "192.168.0.1",
                    //     comment: "shopee_item",
                    //     bind_table: "item"
                    // }
                ],
                totalNumber: 0,
                pageSize: baseConf.pageSize,
            },
            dialogVisible: false,
            optionType: '',


        }
    },
    beforeMount() {
        // 更新缓存数据
        getFilterData()
        getAllPlatformAndProjectMap()
    },
    mounted() {
        this.search(1)
    },
    methods: {
        freshFilterData() {
            this.platFormList = Object.keys(this.$store.state.projectList)
            this.projectList = this.$store.state.projectList[this.filterForm.platForm]
        },
        changePage(page) {
            this.search(page)
        },
        search(page) {
            this.filterForm.page = page
            getProjectInfos(this.filterForm, this.tableData)
        },
        add() {
            
        },
        option(row) {
            console.log(row)
        },
    },
}
</script>


<style>
.member-project .title {
    font-size: 140%;
}

.member-project .el-table {
    margin: 5px;
    border-radius: 5px;
    width: 100%;
}

.el-collapse .el-collapse-item, 
.el-collapse-item .el-collapse-item__header, 
.el-collapse-item .el-collapse-item__wrap {
    background-color: #E9EEF3 !important
}
</style>
