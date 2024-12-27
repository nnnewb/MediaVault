<script setup>
import { ArrowLeft, Check, Document, Folder, Refresh } from "@element-plus/icons-vue";
import dayjs from "dayjs";
import {
  ElButton,
  ElButtonGroup,
  ElCol,
  ElContainer,
  ElDialog,
  ElIcon,
  ElInput,
  ElRow,
  ElTable,
  ElTableColumn,
  ElText,
  ElTooltip,
} from "element-plus";
import { defineEmits, defineProps, inject, reactive, ref } from "vue";

const emits = defineEmits(["choose"]);
const props = defineProps({
  choose_dir: {
    type: Boolean,
    default: false,
  },
  multiple: {
    type: Boolean,
    default: false,
  },
});

const axios = inject("axios");
const choosed_path = ref("");
const current_path = ref("");
const entries = reactive([]);
const history = reactive([]);

function load_entries(path) {
  axios
    .post("/api/v1/path/list", {
      path: path,
    })
    .then((resp) => {
      if (resp.data.code !== 0) {
        console.log(resp.data);
        return;
      }

      entries.length = 0;
      entries.push(...resp.data.data);
    });
}

function choose_entry(row) {
  if (row) {
    choosed_path.value = row.path;
  } else {
    choosed_path.value = "";
  }
}

function enter_dir(row) {
  load_entries(row.path);
  history.push(current_path.value);
  current_path.value = row.path;
  choosed_path.value = row.path;
}

function back() {
  const cur = current_path.value;
  const path = history.pop();
  load_entries(path);
  current_path.value = path;
  choosed_path.value = cur;
}

function refresh() {
  load_entries(current_path.value);
  choosed_path.value = "";
}

function confirm() {
  emits("choose", current_path.value);
}

load_entries("");
</script>

<template>
  <el-dialog ref="dialog" title="选择路径" width="800">
    <el-container direction="vertical">
      <el-row>
        <el-col :span="24">
          <el-input v-model="current_path">
            <template #prepend>
              <el-button @click="back">
                <el-icon>
                  <arrow-left />
                </el-icon>
              </el-button>
            </template>

            <template #append>
              <el-button-group>
                <el-button @click="refresh">
                  <el-icon>
                    <Refresh />
                  </el-icon>
                </el-button>
              </el-button-group>
            </template>
          </el-input>
        </el-col>
      </el-row>

      <el-row>
        <el-col :span="24">
          <el-table
            :data="entries"
            highlight-current-row
            @current-change="choose_entry"
            @row-dblclick="enter_dir"
            :max-height="400"
          >
            <el-table-column prop="is_dir" label="类型" width="80" sortable>
              <template #default="scope">
                <el-icon>
                  <folder v-if="scope.row.is_dir" />
                  <document v-else />
                </el-icon>
              </template>
            </el-table-column>
            <el-table-column prop="name" label="文件名" sortable>
              <template #default="scope">
                <div
                  style="
                    overflow: hidden;
                    white-space: nowrap;
                    text-overflow: ellipsis;
                  "
                >
                  <el-tooltip placement="top" :content="scope.row.name">
                    <el-text style="user-select: none">{{
                        scope.row.name
                      }}
                    </el-text>
                  </el-tooltip>
                </div>
              </template>
            </el-table-column>
            <el-table-column
              prop="updated_at"
              label="修改时间"
              width="180"
              sortable
            >
              <template #default="scope">
                <el-text style="user-select: none">{{
                    dayjs(scope.row.updated_at).format("YYYY-MM-DD HH:mm:ss")
                  }}
                </el-text>
              </template>
            </el-table-column>
          </el-table>
        </el-col>
      </el-row>

      <el-row>
        <el-col
          :span="20"
          style="overflow: hidden; white-space: nowrap; text-overflow: ellipsis"
        >
          <el-input v-model="choosed_path" />
        </el-col>

        <el-col :push="1" :span="2">
          <el-button type="primary" @click="confirm">
            <el-icon>
              <check />
            </el-icon>
            确定
          </el-button>
        </el-col>
      </el-row>
    </el-container>
  </el-dialog>
</template>

<style scoped>
.el-container {
  gap: 7px;
}
</style>
