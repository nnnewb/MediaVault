<script setup>
import { AnimeClient } from "@/api.js";
import { Picture as IconPicture } from "@element-plus/icons-vue";
import ChooseMedia from "@/components/ChooseMedia";

defineProps({
  id: { required: true },
});

const anime_client = new AnimeClient(inject("axios"));
const anime_info = ref({});
const select_loading = ref(false);
const select_options = ref([]);
const select_value = ref(0);
const choose_media_dialog_visible = ref(false);

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

function choose_episode_media() {
  ChooseMedia.choose_single().then((result) => {
    console.log("choose single", result);
  });
}
</script>

<template>
  <el-container direction="vertical">
    <el-row :gutter="7">
      <el-col :span="22">
        <el-form-item label="搜索">
          <el-select remote :remote-method="fetch_suggestions" :loading="select_loading" filterable
                     v-model="select_value" @change="on_select_change">
            <el-option v-for="item in select_options" :label="item.label" :value="item.value"></el-option>
          </el-select>
        </el-form-item>
      </el-col>
      <el-col :span="2">
        <el-button type="primary">保存</el-button>
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

    <el-row>
      <el-card shadow="never" class="episodes-card">
        <template #header>
          <el-row>
            <el-text size="large">剧集</el-text>
            <el-button type="primary" class="push" @click="choose_episode_media">
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

      <choose-media v-model="choose_media_dialog_visible"></choose-media>
    </el-row>
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

.episodes-card {
  width: 100%;
}

.bold {
  font-weight: bold;
}

.push {
  margin-left: auto;
}
</style>
