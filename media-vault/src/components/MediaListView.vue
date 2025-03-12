<script setup>
import dayjs from 'dayjs';
const props = defineProps({
    data: {
        type: Array,
        default: () => []
    }
});
</script>

<template>
    <el-table :data="props.data">
        <el-table-column prop="id" label="id" width="100" />
        <el-table-column label="封面" width="200">
            <template #default="scope">
                <el-image :src="'/api/v1/media/cover/' + scope.row.cover_id" v-if="scope.row.cover_id" width="400"
                    height="300" alt="media cover" />
                <el-empty :image-size="40" description="无封面" v-else />
            </template>
        </el-table-column>
        <el-table-column label="创建时间" width="300">
            <template #default="scope">
                <el-icon>
                    <timer />
                </el-icon>
                {{ dayjs(scope.row.created_at).format('YYYY-MM-DD HH:mm:ssZ') }}
            </template>
        </el-table-column>
        <el-table-column prop="path" label="路径" />
    </el-table>
</template>

<style scoped></style>