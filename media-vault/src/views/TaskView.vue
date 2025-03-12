<script setup>
import { inject, onMounted, reactive, ref } from "vue";
import { TaskClient } from "@/api.js";

const axios = inject("axios");
const client = new TaskClient(axios);
const page = ref(1);
const page_size = ref(20);
const total = ref(10);
const data = reactive([]);

function change_page(value) {
  console.log("change page ", value);
  page.value = value;
  refresh();
}

function change_page_size(value) {
  page_size.value = value;
  refresh();
}

function refresh() {
  console.log("refresh", page.value, page_size.value);
  client.list(page.value, page_size.value).then((resp) => {
    if (resp.data.code !== 0) {
      console.error(resp.data.message);
      return;
    }

    total.value = resp.data.data.total;
    data.length = 0;
    data.push(...resp.data.data.data);
  });
}

onMounted(refresh);
// refresh();
// load_data(page.value, page_size.value);
</script>

<template>
  <el-container direction="vertical">
    <el-table :data="data">
      <el-table-column prop="id" label="id" width="100" />
      <el-table-column prop="name" label="名称" width="300" />
      <el-table-column prop="description" label="描述" />
      <el-table-column prop="status_string" label="状态" width="300" />
      <el-table-column prop="progress" label="进度" width="300">
        <template #default="scope">
          <el-progress :percentage="(scope.row.progress_complete/scope.row.progress_total)*100"
                       :status="scope.row.progress_complete===scope.row.progress_total?'success':''" />
        </template>
      </el-table-column>
    </el-table>
    <el-footer>
      <el-pagination background layout="prev, pager, next, jumper"
                     :total="total" :page-size="page_size" @current-change="change_page" />
    </el-footer>
  </el-container>
</template>

<style scoped>
.el-container {
  gap: 7px;
}
</style>
