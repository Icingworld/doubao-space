<script setup>
import { useRoute } from 'vue-router'

const route = useRoute()
</script>

<template>
  <el-card shadow="never" class="placeholder-card">
    <el-empty :description="route.meta.description || '功能正在准备中'">
      <template #image>
        <div class="placeholder-icon">✦</div>
      </template>
    </el-empty>
  </el-card>
</template>
