<script setup>
const medias = [
];
for (let i = 0; i < 20; i++) {
  medias.push({ path: `${i}.mkv`, cover: null })
}
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
