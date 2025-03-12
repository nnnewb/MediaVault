<script setup>
import { inject } from 'vue';

const axios = inject('axios');
const tree_props = {
    label: 'name',
    isLeaf: item => !item.is_dir,
};

function load_node(node, resolve, reject) {
    if (!node.data.path) {
        resolve([{
            name: 'D:/',
            is_dir: true,
            path: 'D:/'
        }])
        return
    }
    axios.post('/api/v1/path/list', { path: node.data.path }).then(resp => {
        resolve(resp.data.data)
    }).catch(reject)
}
</script>

<template>
    <el-dialog title="选择路径" width="800">
        <el-tree show-checkbox :load="load_node" :props="tree_props" lazy />
    </el-dialog>
</template>

<style scoped></style>