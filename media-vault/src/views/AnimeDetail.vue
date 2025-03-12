<script setup>
import { ref } from 'vue';

const searchPath = ref('');
const animeSearch = ref('');
const animeInfo = ref('');
const episodeList = ref('');
const spAnimation = ref('');
const ovaAnimation = ref('');
const ost = ref('');
</script>

<template>
  <el-container>
    <el-header>
      <el-input placeholder="路径展示" v-model="searchPath" clearable></el-input>
    </el-header>

    <el-main>
      <el-row :gutter="20">
        <el-col :span="12">
          <el-card>
            <el-input placeholder="搜索单选番剧（自动识别）" v-model="animeSearch" clearable></el-input>
          </el-card>
        </el-col>
        <el-col :span="12">
          <el-card>
            <el-input placeholder="番剧信息" v-model="animeInfo" clearable></el-input>
          </el-card>
        </el-col>
      </el-row>

      <el-row :gutter="20" style="margin-top: 20px;">
        <el-col :span="24">
          <el-card>
            <el-input placeholder="剧集清单（自动识别）" v-model="episodeList" clearable></el-input>
          </el-card>
        </el-col>
      </el-row>

      <el-row :gutter="20" style="margin-top: 20px;">
        <el-col :span="12">
          <el-card>
            <el-input placeholder="SP 动画（自动识别）" v-model="spAnimation" clearable></el-input>
          </el-card>
        </el-col>
        <el-col :span="12">
          <el-card>
            <el-input placeholder="OVA动画（自动识别）" v-model="ovaAnimation" clearable></el-input>
          </el-card>
        </el-col>
      </el-row>

      <el-row :gutter="20" style="margin-top: 20px;">
        <el-col :span="24">
          <el-card>
            <el-input placeholder="OST（自动识别）" v-model="ost" clearable></el-input>
          </el-card>
        </el-col>
      </el-row>
    </el-main>
  </el-container>
</template>

<style scoped>

</style>