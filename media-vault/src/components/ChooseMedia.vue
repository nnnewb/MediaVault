<script setup>
import { MediaClient } from "@/api.js";

const media_client = new MediaClient(inject("axios"));
const page = ref(1);
const page_size = ref(20);
const search_text = ref("");
const data = ref([]);
const default_sort = { prop: "path", order: "ascending" };
const selectable = (row) => true;


function refresh() {
  fetch_data(search_text.value, page.value, page_size.value);
}

watch(page, refresh);
watch(page_size, refresh);
watch(search_text, refresh);

/**
 * 获取数据
 * @param {string} q
 * @param {number} page
 * @param {number} page_size
 */
function fetch_data(q, page, page_size) {
  media_client.list(q, page, page_size).then((resp) => {
    if (resp.code !== 0) {
      console.error(resp.message);
      return;
    }

    data.value = resp.data.data;
  }).catch((err) => {
    console.error(err);
  });
}

onMounted(() => {
  refresh();
});
</script>

<template>
  <el-dialog>
    <el-row :gutter="7">
      <el-input type="text" placeholder="搜索..." prefix-icon="Search" v-model="search_text"></el-input>
    </el-row>
    <el-table :data="data" max-height="500px" :default-sort="default_sort" scrollbar-always-on show-overflow-tooltip>
      <el-table-column type="selection" :selectable="selectable" width="55" />
      <el-table-column label="文件名" prop="name" />
      <el-table-column label="路径" prop="path" />
    </el-table>
    <template #footer>
      <el-row>
        <el-button type="primary" icon="Check" class="push">确定</el-button>
        <el-button type="info" icon="Close">取消</el-button>
      </el-row>
    </template>
  </el-dialog>
</template>

<style scoped>
.push {
  margin-left: auto;
}
</style>
