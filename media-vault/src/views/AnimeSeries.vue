<script setup>
import { reactive } from 'vue';

const searchForm = reactive({
    search: '',
    displayType: 'list'
});
const data = reactive([]);
for (let index = 0; index < 10; index++) {
    data.push({
        title: `yuru yuri`,
        release: `2022-01-01`,
        episode: `12`,
        status: `已完结`
    })
}
</script>
<template>
    <el-container direction="vertical">
        <!-- action bar -->
        <el-row>
            <el-form :inline="true" :model="searchForm">
                <el-form-item>
                    <el-input type="text" placeholder="输入开始搜索" v-model="searchForm.search" prefix-icon="Search"
                        clearable />
                </el-form-item>
                <el-form-item label="视图模式">
                    <el-radio-group v-model="searchForm.displayType">
                        <el-radio-button value="grid">
                            照片墙
                            <el-icon>
                                <grid />
                            </el-icon>
                        </el-radio-button>
                        <el-radio-button value="list">
                            列表
                            <el-icon>
                                <list />
                            </el-icon>
                        </el-radio-button>
                    </el-radio-group>
                </el-form-item>
            </el-form>
        </el-row>

        <el-table>
            <el-table-column prop="title" label="标题"></el-table-column>
            <el-table-column prop="release" label="发行日期"></el-table-column>
            <el-table-column prop="episode" label="集数"></el-table-column>
            <el-table-column prop="status" label="状态"></el-table-column>
        </el-table>
    </el-container>
</template>
<style scoped></style>