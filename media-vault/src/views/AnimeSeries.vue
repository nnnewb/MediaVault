<script setup>
import { inject, onMounted, ref, watch } from "vue";
import { AnimeClient } from "@/api.js";

const axios = inject("axios");
const anime_client = new AnimeClient(axios);

const page = ref(1);
const page_size = ref(20);
const total = ref(0);
const search = ref("");
const data = ref([]);

function fetch_data(page, page_size) {
  anime_client
    .list(page, page_size)
    .then((resp) => {
      if (resp.code !== 0) {
        console.error(resp.message);
        return;
      }

      total.value = resp.data.total;
      data.value = resp.data.data;
    })
    .catch((err) => {
      console.error(err);
    });
}

onMounted(() => {
  fetch_data(1, 20);
  watch(page, () => fetch_data(page.value, page_size.value));
  watch(page_size, () => fetch_data(page.value, page_size.value));
});
</script>

<template>
  <el-container direction="vertical">
    <!-- action bar -->
    <el-row :gutter="7">
      <el-col :span="12">
        <el-input type="text" placeholder="输入开始搜索" v-model="search" prefix-icon="Search" clearable />
      </el-col>
      <el-col :span="4">
        <router-link to="/anime/edit/new">
          <el-button type="primary">
            <el-icon>
              <plus />
            </el-icon>
            添加
          </el-button>
        </router-link>
      </el-col>
    </el-row>

    <el-table :data="data">
      <el-table-column prop="title" label="标题"></el-table-column>
      <el-table-column prop="release_year" label="发行日期"></el-table-column>
      <el-table-column prop="season" label="季度"></el-table-column>
      <el-table-column prop="total_episodes" label="集数"></el-table-column>
      <el-table-column prop="status" label="状态"></el-table-column>
    </el-table>

    <el-footer>
      <el-pagination v-model:current-page="page" v-model:page-size="page_size" :total="total"
                     layout="prev, pager, next, jumper"></el-pagination>
    </el-footer>
  </el-container>
</template>
<style scoped></style>
