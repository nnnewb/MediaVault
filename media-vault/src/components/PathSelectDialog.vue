<script setup>
const tree_props = {
    label: 'label',
    isLeaf: 'leaf'
};

function load_node(node, resolve, reject) {
    resolve([
        {
            label: 'hello',
            leaf: true
        },
        {
            label: 'folder',
            leaf: false
        },
    ])
}
</script>

<template>
    <el-dialog title="选择路径" width="800">
        <el-tree show-checkbox :load="load_node" :props="tree_props" lazy />
    </el-dialog>
</template>

<style scoped></style>