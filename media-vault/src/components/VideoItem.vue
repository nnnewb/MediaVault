<script setup>
const props=defineProps({
    media: {
        type: Object,
        required: true
    }
});

const media = props.media;
</script>
<template>
    <el-card>
        {{ media.path }}
    </el-card>
</template>