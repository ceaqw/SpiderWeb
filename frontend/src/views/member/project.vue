<template>
    <div id="member-project">
        <el-row>
            <el-collapse accordion style="width: 100%">
                <el-collapse-item title="筛选" name="1">
                    <el-row>
                        <el-col :span="5">
                            <div class="demonstration">平台</div>
                            <el-select v-model="filterForm.platform" @focus="freshFilterData" @change="filterForm.project='all'">
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
                <el-table-column prop="id" label="Id" width="40%"></el-table-column>
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
            <!-- <el-form ref="form" :inline="true" :model="form" label-position="left" label-width="100px"> -->
            <el-form ref="projectForm" :model="projectForm" :rules="rules" label-position="left" label-width="120px">
                <el-form-item label="Platform:" prop="platform">
                    <el-select v-model="projectForm.platform" placeholder="请选择平台">
                        <el-option v-for="platform, index in formPlatform" :key="index" :label="platform" :value="platform"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="Project:" prop="project">
                    <el-input v-model="projectForm.project"></el-input>
                </el-form-item>
                <el-form-item label="Script Id:" prop="scriptId">
                    <el-input v-model.number="projectForm.scriptId"></el-input>
                </el-form-item>
                <el-form-item label="Ip:" prop="ip">
                    <el-input v-model="projectForm.ip"></el-input>
                </el-form-item>
                <el-form-item label="Comment:" prop="comment">
                    <el-input v-model="projectForm.comment"></el-input>
                </el-form-item>
                <el-form-item label="Server:" prop="server">
                    <el-select v-model.number="projectForm.server" placeholder="请选择Server">
                    <el-option label="Japan" :value="0"></el-option>
                    <el-option label="China" :value="1"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="Critical Kpi:" prop="criticalKpi">
                    <el-input v-model="projectForm.criticalKpi"></el-input>
                </el-form-item>
                <el-form-item label="Bind Table:" prop="bindTables">
                    <el-checkbox-group v-model="projectForm.bindTables">
                        <el-checkbox v-for="table, index in formBindTables" :key="index" :label="table" name="type"></el-checkbox>
                    </el-checkbox-group>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="saveProject">保存</el-button>
                    <el-button type="warning" @click="clearForm('projectForm')">清空</el-button>
                    <el-button @click="dialogVisible=false">关闭</el-button>
                </el-form-item>
            </el-form>
        </el-dialog>
    </div>
</template>

<script>
import { getFilterData } from '@/service/datas/base.js'
import { getProjectInfos, getAllPlatformAndProjectMap, createProject } from '../../service/project'
import baseConf from '@/conf/baseConf'
export default {
    data() {
        return {
            filterForm: { platform: 'all', project: 'all', page: 1, pageSize: baseConf.pageSize },
            platFormList: ['all'],
            projectList: ['all'],
            formPlatform: [],
            formBindTables: [],
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
            projectForm: {
                id: 0,
                platform: '',
                project: '',
                scriptId: 0,
                ip: '',
                comment: '',
                server: '',
                criticalKpi: '',
                bindTables: [],
            },
            dialogVisible: false,
            optionType: '',
            rules: {
                platform: { required: true, message: '必选', trigger: 'blur' },
                project: { required: true, message: '必选', trigger: 'blur' },
                scriptId: [
                    { required: true, message: '必选', trigger: 'blur' },
                    { type: 'number', message: '类型错误', trigger: 'blur' }
                ],
                ip: { required: true, message: '必选', trigger: 'blur' },
                server: { required: true, message: '必选', trigger: 'blur' },
                bindTables: { required: true, message: '必选', trigger: 'blur' },
            }
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
    watch: {
        'projectForm.platform': {
            handler () {
                this.freshBindTables()
            }
        },
    },
    methods: {
        freshBindTables() {
            this.formBindTables = this.$store.state.allPlatformAndProjectMap[this.projectForm.platform]
        },
        freshFilterData() {
            this.platFormList = Object.keys(this.$store.state.projectList)
            this.projectList = this.$store.state.projectList[this.filterForm.platform]
        },
        clearForm(formName) {
            if (formName) this.$refs[formName].resetFields()
            this.projectForm.platform = ''
            this.projectForm.project = ''
            this.projectForm.scriptId = ''
            this.projectForm.ip = ''
            this.projectForm.comment = ''
            this.projectForm.server = ''
            this.projectForm.criticalKpi = ''
            this.projectForm.bindTables = []
        },
        changePage(page) {
            this.search(page)
        },
        search(page) {
            this.filterForm.page = page
            getProjectInfos(this.filterForm, this.tableData)
        },
        add() {
            this.formPlatform = Object.keys(this.$store.state.allPlatformAndProjectMap)
            this.optionType = '添加'
            // this.clearForm()
            this.dialogVisible = true
        },
        option(row) {
            this.formPlatform = Object.keys(this.$store.state.allPlatformAndProjectMap)
            this.optionType = '设置'
            this.dialogVisible = true
            this.projectForm.id = row.id
            this.projectForm.platform = row.platform
            this.projectForm.project = row.project
            this.projectForm.scriptId = row.script_id
            this.projectForm.bindTables = row.bind_table.split(',')
            this.projectForm.ip = row.ip
            this.projectForm.comment = row.comment
            this.projectForm.server = row.server
            this.projectForm.criticalKpi = row.critical_kpi
            // console.log(row)
        },
        saveProject() {
            this.$refs.projectForm.validate((valid) => {
                if (valid) {
                    let formData = Object.assign({}, this.projectForm)
                    formData.bindTables = formData.bindTables.join(',')
                    createProject(formData)
                } else {
                    return false
                }
            })
        }
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

.el-dialog .el-form-item {
    text-align: left;
}
</style>
