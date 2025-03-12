<script setup>
import { inject, ref, reactive } from 'vue';

const page = ref(1);
const page_size = ref(10);
const axios = inject('axios');

const searchForm = reactive({
  search: '',
  displayType: 'list',
});
const medias = reactive([]);

axios.post('/api/v1/media/list', { 'page': page.value, 'page_size': page_size.value, }).then(resp => {
  if (resp.data.code !== 0) {
    console.error(resp.data.message);
    return
  }

  medias.push(...resp.data.data);
})
</script>

<template>
  <el-container direction="vertical">
    <el-row>
      <el-form :inline="true" :model="searchForm">
        <el-form-item>
          <el-input type="text" placeholder="输入开始搜索" v-model="searchForm.search" prefix-icon="Search" clearable />
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

    <media-list-view :data="medias" v-if="searchForm.displayType === 'list'" />
    <media-grid-view :data="medias" v-else-if="searchForm.displayType == 'grid'" />
  </el-container>
</template>

<style scoped></style>
