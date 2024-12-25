<script setup>
import defaultCover from '@/assets/cover.png'
const props = defineProps({
    media: {
        type: Object,
        required: true
    }
});

const cover = props.media.cover || defaultCover;
</script>

<template>
    <el-card class="card">
        <template #header>
            <el-text>{{ props.media.path }}</el-text>
        </template>

        <el-image :src="cover" alt="default-cover" fit="cover" />

        <template #footer>
            <el-row>
                <el-col>
                    <el-tooltip placement="top" content="not implemented">
                        <el-button disabled type="primary" icon="VideoPlay" />
                    </el-tooltip>
                </el-col>
            </el-row>
        </template>
    </el-card>
</template>

<style scoped>
.card {
    width: 440px;
}
</style>
