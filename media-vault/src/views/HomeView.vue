<script setup>
import { inject, reactive, ref } from "vue";
import { ElButton, ElContainer, ElForm, ElFormItem, ElIcon, ElInput, ElRow, ElSegmented } from "element-plus";
import { Plus } from "@element-plus/icons-vue";
import PathSelectDialog from "@/components/PathSelectDialog.vue";
import MediaListView from "@/components/MediaListView.vue";
import MediaGridView from "@/components/MediaGridView.vue";

const page = ref(1);
const page_size = ref(10);
const axios = inject("axios");
const show_path_select_dialog = ref(false);
const view_types = [
  {
    label: "网格",
    value: "grid",
    icon: "grid",
  },
  {
    label: "列表",
    value: "list",
    icon: "list",
  },
];

const searchForm = reactive({
  search: "",
  displayType: "list",
});
const medias = reactive([]);

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
  axios
    .post("/api/v1/media/list", payload)
    .then((resp) => {
      if (resp.data.code !== 0) {
        console.error(resp.data.message);
        return;
      }

      medias.push(...resp.data.data);
    });
}

load_medias(1, 20);
</script>

<template>
  <el-container direction="vertical">
    <el-row>
      <el-form :inline="true" :model="searchForm">
        <!-- 搜索 -->
        <el-form-item>
          <el-input
            type="text"
            placeholder="输入开始搜索"
            v-model="searchForm.search"
            prefix-icon="Search"
            clearable
          />
        </el-form-item>

        <!-- 视图切换 -->
        <el-form-item label="视图">
          <el-segmented
            v-model="searchForm.displayType"
            :options="view_types"
            size="default"
          >
            <template #default="{ item }">
              <el-icon>
                <component :is="item.icon" />
              </el-icon>
              {{ item.label }}
            </template>
          </el-segmented>
        </el-form-item>

        <!-- 添加资源 -->
        <el-form-item>
          <el-button type="primary" @click="show_path_select_dialog = true">
            <el-icon>
              <plus />
            </el-icon>
            添加
          </el-button>
        </el-form-item>
      </el-form>
    </el-row>

    <path-select-dialog v-model="show_path_select_dialog" @choose="confirm_scan" />
    <media-list-view :data="medias" v-if="searchForm.displayType === 'list'" />
    <media-grid-view
      :data="medias"
      v-else-if="searchForm.displayType === 'grid'"
    />
  </el-container>
</template>

<style scoped></style>
