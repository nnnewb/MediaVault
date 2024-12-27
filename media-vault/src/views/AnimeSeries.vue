<script setup>
import { reactive } from "vue";

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
const data = reactive([]);
for (let index = 0; index < 10; index++) {
  data.push({
    title: `yuru yuri`,
    release: `2022-01-01`,
    episode: `12`,
    status: `已完结`,
  });
}
</script>
<template>
  <el-container direction="vertical">
    <!-- action bar -->
    <el-row>
      <el-form :inline="true" :model="searchForm">
        <el-form-item>
          <el-input
            type="text"
            placeholder="输入开始搜索"
            v-model="searchForm.search"
            prefix-icon="Search"
            clearable
          />
        </el-form-item>

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
          <el-button type="primary">
            <el-icon>
              <plus />
            </el-icon>
            添加
          </el-button>
        </el-form-item>
      </el-form>
    </el-row>

    <el-table :data="data">
      <el-table-column prop="title" label="标题"></el-table-column>
      <el-table-column prop="release" label="发行日期"></el-table-column>
      <el-table-column prop="episode" label="集数"></el-table-column>
      <el-table-column prop="status" label="状态"></el-table-column>
    </el-table>
  </el-container>
</template>
<style scoped></style>
