<script setup>
import { AnimeClient } from "@/api.js";
import { Picture as IconPicture } from "@element-plus/icons-vue";

defineProps({
  id: { required: true },
});

const anime_client = new AnimeClient(inject("axios"));
const anime_info = ref({});
const select_loading = ref(false);
const select_options = ref([]);
const select_value = ref(0);

function fetch_suggestions(query) {
  if (query) {
    select_loading.value = true;
    anime_client
      .search(query)
      .then((resp) => {
        if (resp.code !== 0) {
          console.error(resp.message);
          return;
        }
        select_options.value = resp.data.data.map(item => ({ label: item.title, value: item.id }));
        select_loading.value = false;
      })
      .catch((err) => {
        console.error(err);
      });
  }
}

function on_select_change(value) {
  if (value) {
    anime_client.info(value).then((resp) => {
      if (resp.code !== 0) {
        console.error(resp.message);
        return;
      }

      anime_info.value = resp.data;
    }).catch((err) => {
      console.error(err);
    });
  }
}
</script>

<template>
  <el-container direction="vertical">
    <el-row>
      <el-col>
        <el-select remote :remote-method="fetch_suggestions" :loading="select_loading" filterable v-model="select_value"
                   @change="on_select_change" placeholder="从动画数据库搜索...">
          <el-option v-for="item in select_options" :label="item.label" :value="item.value"></el-option>
        </el-select>
      </el-col>
    </el-row>

    <el-row>
      <el-col :span="24">
        <el-descriptions border direction="horizontal" :column="3" label-width="60">
          <el-descriptions-item label="封面" :rowspan="5">
            <el-image :src="anime_info.picture" fit="cover">
              <template #error>
                <el-icon :size="40" class="image-slot">
                  <icon-picture />
                </el-icon>
              </template>
            </el-image>
          </el-descriptions-item>
          <el-descriptions-item label="标题" :span="2">
            <el-text size="large" type="primary" class="bold">
              {{ anime_info.title }}
            </el-text>
          </el-descriptions-item>
          <el-descriptions-item label="年份">{{ anime_info.year }}</el-descriptions-item>
          <el-descriptions-item label="季度">{{ anime_info.season }}</el-descriptions-item>
          <el-descriptions-item label="集数">{{ anime_info.episodes }}</el-descriptions-item>
          <el-descriptions-item label="播出状态" label-width="120">
            <el-tag type="success">
              {{ anime_info.status }}
            </el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="别名" :span="2">
            <el-space wrap>
              <el-tag type="primary" v-for="synonym in anime_info.synonyms">{{ synonym }}</el-tag>
            </el-space>
          </el-descriptions-item>
          <el-descriptions-item label="标签" :span="2">
            <el-space wrap>
              <el-tag v-for="t in anime_info.tags">{{ t }}</el-tag>
            </el-space>
          </el-descriptions-item>
        </el-descriptions>
      </el-col>
    </el-row>

    <el-row :gutter="7">
      <el-col :span="12">
        <el-card shadow="never">
          <template #header>
            <el-row>
              <el-text size="large">剧集</el-text>
              <el-button type="primary" class="push">
                <el-icon>
                  <Plus />
                </el-icon>
                添加
              </el-button>
            </el-row>
          </template>
          <el-table height="500">
            <el-table-column label="文件名" />
            <el-table-column label="修改时间" />
            <el-table-column label="操作"></el-table-column>
          </el-table>
        </el-card>
      </el-col>

      <el-col :span="12">
        <el-card shadow="never">
          <template #header>
            <el-row>
              <el-text size="large">OVA</el-text>
              <el-button type="primary" class="push">
                <el-icon>
                  <Plus />
                </el-icon>
                添加
              </el-button>
            </el-row>
          </template>
          <el-table height="500">
            <el-table-column label="文件名" />
            <el-table-column label="修改时间" />
            <el-table-column label="操作"></el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
    <el-footer></el-footer>
  </el-container>
</template>

<style scoped>
.el-container {
  gap: 7px;
}

.el-image {
  width: 225px;
  height: 320px;
}

.el-image .image-slot {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  background: var(--el-fill-color-light);
  color: var(--el-text-color-secondary);
  font-size: 30px;
}

.bold {
  font-weight: bold;
}

.push {
  margin-left: auto;
}
</style>
