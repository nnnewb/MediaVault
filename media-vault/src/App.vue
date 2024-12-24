<script setup>
import zhCn from 'element-plus/es/locale/lang/zh-cn';
import { RouterView } from 'vue-router';
import HeadNav from '@/components/HeadNav.vue'
import AsideNav from '@/components/AsideNav.vue'
import FootNote from '@/components/FootNote.vue'

const size = 'small';
const zIndex = 3000;
</script>

<template>
  <el-config-provider :size="size" :z-index="zIndex" :locale="zhCn">
    <el-container>
      <el-header class="header">
        <head-nav />
      </el-header>
      <el-container>
        <el-aside>
          <aside-nav />
        </el-aside>
        <el-container>
          <el-main>
            <RouterView />
          </el-main>
          <el-footer>
            <foot-note />
          </el-footer>
        </el-container>
      </el-container>
    </el-container>
  </el-config-provider>
</template>

<style scoped>
.header {
  padding: 0;
}
</style>
