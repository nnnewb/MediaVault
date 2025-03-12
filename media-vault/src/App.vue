<script setup>
import AsideNav from '@/components/AsideNav.vue';
import HeadNav from '@/components/HeadNav.vue';
import zhCn from 'element-plus/es/locale/lang/zh-cn';
import { RouterView } from 'vue-router';

const size = 'default';
const zIndex = 3000;
</script>

<template>
  <el-config-provider :size="size" :z-index="zIndex" :locale="zhCn">
    <el-container class="expand-height">
      <el-aside class="expand-height">
        <aside-nav />
      </el-aside>

      <el-container class="expand-height">
        <el-header class="header">
          <head-nav />
        </el-header>
        <el-main>
          <RouterView />
        </el-main>
      </el-container>
    </el-container>
  </el-config-provider>
</template>

<style scoped>
.expand-height {
  height: 100%;
}

.header {
  padding: 0;
}

.aside {
  height: 100%;
}
</style>
