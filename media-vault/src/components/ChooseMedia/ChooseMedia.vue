<script setup>
import { MediaClient } from "@/api.js";

const props = defineProps({
  multiple: {
    type: Boolean,
    default: false,
  },
});
const emit = defineEmits(["cancel", "confirm"]);

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

const visible = ref(true);
const dialog = useTemplateRef("dialog");
const tbl = useTemplateRef("tbl");
const cur = ref(null);

function confirm() {
  if (props.multiple) {
    emit("confirm", tbl.value.getSelectionRows());
    visible.value = false;
  } else {
    emit("confirm", cur.value);
    visible.value = false;
  }
}

function cancel() {
  emit("cancel");
  visible.value = false;
}

onMounted(() => {
  refresh();
});
</script>

<template>
  <el-dialog v-model="visible" ref="dialog" :show-close="false">
    <el-row :gutter="7">
      <el-input type="text" placeholder="搜索..." prefix-icon="Search" v-model="search_text"></el-input>
    </el-row>
    <el-table ref="tbl" :highlight-current-row="!multiple" :data="data" max-height="500px" :default-sort="default_sort"
              @current-change="(row)=>cur=row" scrollbar-always-on show-overflow-tooltip>
      <el-table-column type="selection" v-if="multiple" :selectable="selectable" width="55" />
      <el-table-column label="文件名" prop="name" />
      <el-table-column label="路径" prop="path" />
    </el-table>
    <template #footer>
      <el-row>
        <el-button type="primary" icon="Check" class="push" @click="confirm">确定</el-button>
        <el-button type="info" icon="Close" @click="cancel">取消</el-button>
      </el-row>
    </template>
  </el-dialog>
</template>

<style scoped>
.push {
  margin-left: auto;
}
</style>
