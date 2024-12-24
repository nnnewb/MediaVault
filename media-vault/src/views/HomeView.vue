<script setup>
const medias = [
  {
    path: '01.mkv',
  },
  {
    path: '02.mkv'
  }
];
</script>

<template>
  <el-container>
    <video-item v-for="media in medias" :media="media">
    </video-item>
  </el-container>
</template>
