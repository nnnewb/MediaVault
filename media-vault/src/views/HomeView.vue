<script setup>
import { inject, onMounted, ref, watch } from "vue";
import PathSelectDialog from "@/components/PathSelectDialog.vue";
import MediaListView from "@/components/MediaListView.vue";
import MediaGridView from "@/components/MediaGridView.vue";
import { MediaClient } from "@/api.js";

const page = ref(1);
const page_size = ref(10);
const total = ref(0);
const axios = inject("axios");
const mediaClient = new MediaClient(axios);
const show_path_select_dialog = ref(false);
const search = ref("");
const view_type = ref("list");
const view_types = [
  { label: "网格", value: "grid", icon: "grid" },
  { label: "列表", value: "list", icon: "list" },
];
const medias = ref([]);

watch(page, () => load_medias(page.value, page_size.value));
watch(page_size, () => load_medias(page.value, page_size.value));

function confirm_scan(path) {
  axios.post("/api/v1/media/scan", { paths: [path] }).then(resp => {
    if (resp.data.code !== 0) {
      console.error(resp.data.message);
    }
  }).then(() => {
    show_path_select_dialog.value = false;
    load_medias(1, 20);
  });
}

function load_medias(page, page_size) {
  const payload = { page, page_size };
  mediaClient.list(page, page_size).then(resp => {
    if (resp.code !== 0) {
      console.error(resp.message);
      return;
    }

    medias.value = resp.data.data;
    total.value = resp.data.total;
  });
}

onMounted(() => load_medias(1, 20));
</script>

<template>
  <el-container direction="vertical">
    <el-row :gutter="7">
      <el-col :span="12">
        <el-input type="text" placeholder="输入开始搜索" v-model="search" prefix-icon="Search" clearable />
      </el-col>

      <el-col :span="2">
        <el-segmented v-model="view_type" :options="view_types" size="default">
          <template #default="{ item }">
            <el-icon>
              <component :is="item.icon" />
            </el-icon>
            {{ item.label }}
          </template>
        </el-segmented>
      </el-col>

      <el-col :span="4">
        <el-button type="primary" @click="show_path_select_dialog = true">
          <el-icon>
            <plus />
          </el-icon>
          添加
        </el-button>
      </el-col>
    </el-row>

    <path-select-dialog v-model="show_path_select_dialog" @choose="confirm_scan" />
    <media-list-view :data="medias" v-if="view_type === 'list'" />
    <media-grid-view :data="medias" v-else-if="view_type === 'grid'" />

    <el-footer>
      <el-pagination layout="prev, pager, next, jumper" v-model:current-page="page" v-model:page-size="page_size"
                     :total="total" />
    </el-footer>
  </el-container>
</template>

<style scoped></style>
