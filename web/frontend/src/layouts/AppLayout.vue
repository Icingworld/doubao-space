<script setup>
import { computed } from 'vue'
import { RouterLink, RouterView, useRoute } from 'vue-router'

const route = useRoute()
const pageTitle = computed(() => route.meta.title || '豆包空间')

const navigation = [
  { label: '工作台', to: '/dashboard' },
  { label: '日程管理', to: '/dashboard/calendar' },
  { label: '内容收集', to: '/dashboard/notes' },
  { label: '设置', to: '/dashboard/settings' },
]
</script>

<template>
  <el-container class="app-shell">
    <el-aside width="220px" class="app-sidebar">
      <div class="brand">
        <div class="brand-mark">豆</div>
        <div>
          <div class="brand-title">豆包空间</div>
          <div class="brand-subtitle">个人工作台</div>
        </div>
      </div>

      <el-menu :default-active="route.path" router class="app-menu">
        <el-menu-item v-for="item in navigation" :key="item.to" :index="item.to">
          <RouterLink :to="item.to" class="menu-link">{{ item.label }}</RouterLink>
        </el-menu-item>
      </el-menu>

      <div class="sidebar-footer">轻量 · 本地 · 私人</div>
    </el-aside>

    <el-container>
      <el-header class="app-header">
        <div>
          <div class="eyebrow">DOUBAO SPACE</div>
          <h1>{{ pageTitle }}</h1>
        </div>
        <el-tag type="success" effect="light">本地运行</el-tag>
      </el-header>

      <el-main class="app-main">
        <RouterView />
      </el-main>
    </el-container>
  </el-container>
</template>
