<script setup>
import { Picture as IconPicture } from "@element-plus/icons-vue";
import { defineProps } from "vue";
const props = defineProps({
  data: {
    type: Array,
    default: () => [],
  },
});

function build_cover_url(id) {
  return `/api/v1/media/cover/${id}`;
}
</script>

<template>
  <el-container>
    <el-row class="grid">
      <el-card v-for="item in props.data" class="card" shadow="hover">
        <template #header>
          {{ item.name }}
        </template>
        <el-image
          :src="build_cover_url(item.cover_id)"
          v-if="item.cover_id"
          width="400"
          height="300"
          alt="media cover"
        />
        <el-image v-else>
          <template #error>
            <el-icon :size="40" class="image-slot">
              <icon-picture />
            </el-icon>
          </template>
        </el-image>
      </el-card>
    </el-row>
  </el-container>
</template>

<style scoped>
.grid {
  gap: 7px;
}

.card {
  width: 440px;
}

.el-image {
  width: 400px;
  height: 300px;
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
</style>
