<script setup>
import { ElIcon, ElMenu, ElMenuItem } from "element-plus";
import { Suitcase, VideoCamera, VideoCameraFilled } from "@element-plus/icons-vue";
</script>

<template>
  <el-menu :router="true" default-active="1" class="aside-menu">
    <el-menu-item route="/" index="1">
      <el-icon>
        <video-camera />
      </el-icon>
      媒体库
    </el-menu-item>
    <el-menu-item route="/anime" index="2">
      <el-icon>
        <video-camera-filled />
      </el-icon>
      动画
    </el-menu-item>
    <el-menu-item route="/task" index="3">
      <el-icon>
        <suitcase />
      </el-icon>
      后台任务
    </el-menu-item>
  </el-menu>
</template>

<style scoped>
.aside-menu {
  height: 100%;
}
</style>
