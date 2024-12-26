<script setup>
</script>

<template>
    <el-menu :router="true" default-active="1" class="aside-menu">
        <el-menu-item route="/" index="1">
            <el-icon><VideoCamera /></el-icon>
            媒体库
        </el-menu-item>
        <el-menu-item route="/task" index="2">
            <el-icon><Suitcase /></el-icon>
            后台任务
        </el-menu-item>
    </el-menu>
</template>

<style scoped>
.aside-menu {
    height: 100%;
}
</style>