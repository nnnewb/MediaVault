<script setup>
import { onMounted, ref, watch } from "vue";
import { useRoute } from "vue-router";

const default_active = ref("/");
const routes = [
  { path: "/", label: "媒体库", icon: "VideoCamera" },
  { path: "/anime", label: "动画", icon: "VideoCameraFilled" },
  { path: "/task", label: "后台任务", icon: "Suitcase" },
];

watch(useRoute(), (route) => {
  console.log(route.path);
  default_active.value = route.path;
});

onMounted(() => {
  default_active.value = useRoute().path;
});
</script>

<template>
  <el-menu mode="vertical" :router="true" :default-active="default_active" class="aside-menu">
    <el-menu-item :route="route.path" :index="route.path" v-for="route in routes">
      <el-icon>
        <component :is="route.icon" />
      </el-icon>
      {{ route.label }}
    </el-menu-item>
  </el-menu>
</template>

<style scoped>
.aside-menu {
  height: 100%;
}
</style>
