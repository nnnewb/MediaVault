<script setup>
import { inject, ref, reactive } from 'vue';

const page = ref(1);
const page_size = ref(10);
const axios = inject('axios');

const searchForm = reactive({
  search: '',
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
      </el-form>
    </el-row>

    <el-table :data="medias">
      <el-table-column prop="id" label="id" width="100" />
      <el-table-column prop="created_at" label="创建时间" width="300" />
      <el-table-column prop="updated_at" label="更新时间" width="300" />
      <el-table-column prop="information_id" label="信息id" width="100" />
      <el-table-column prop="cover_id" label="封面id" width="100" />
      <el-table-column prop="path" label="路径" />
    </el-table>
  </el-container>
</template>

<style scoped>
.row {
  gap: 7px;
}
</style>
