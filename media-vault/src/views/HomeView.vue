<script setup>
import { inject, ref, reactive } from 'vue';

const page = ref(1);
const page_size = ref(10);
const axios = inject('axios');

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
  <el-container>
    <el-row class="row">
      <video-item v-for="media in medias" :media="media" />
    </el-row>
  </el-container>
</template>

<style scoped>
.row {
  gap: 7px;
}
</style>
