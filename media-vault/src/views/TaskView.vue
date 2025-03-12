<script setup>
const tasks = [
    {
        id: 1,
        name: '提取封面',
        status_string: 'finish',
        progress: 1,
        total: 1,
    }
];
</script>

<template>
    <el-container direction="vertical">
        <el-table :data="tasks">
            <el-table-column prop="id" label="id" />
            <el-table-column prop="name" label="名称" />
            <el-table-column prop="status_string" label="状态" />
            <el-table-column prop="progress" label="进度" />
        </el-table>
    </el-container>
</template>

<style scoped>
</style>