<script setup>
import { onMounted, ref } from 'vue'
import { getHealth } from '../api'

const health = ref(null)
const error = ref('')

async function checkHealth() {
  error.value = ''
  try {
    health.value = await getHealth()
  } catch (requestError) {
    health.value = null
    error.value = requestError.message
  }
}

onMounted(checkHealth)
</script>

<template>
  <div class="page-stack">
    <section class="hero-card">
      <div>
        <p class="eyebrow">WELCOME BACK</p>
        <h2>把今天安排得刚刚好。</h2>
        <p class="hero-copy">这是一个从本地开始的个人工作台，数据保存在你自己的设备上。</p>
      </div>
      <div class="hero-orb">✦</div>
    </section>

    <section class="stats-grid">
      <el-card shadow="never">
        <div class="stat-label">今日待办</div>
        <div class="stat-value">0</div>
        <div class="stat-hint">准备开始记录</div>
      </el-card>
      <el-card shadow="never">
        <div class="stat-label">本周日程</div>
        <div class="stat-value">0</div>
        <div class="stat-hint">保持节奏</div>
      </el-card>
      <el-card shadow="never">
        <div class="stat-label">收集内容</div>
        <div class="stat-value">0</div>
        <div class="stat-hint">灵感不会丢失</div>
      </el-card>
    </section>

    <el-card shadow="never" class="status-card">
      <template #header>
        <div class="card-header">
          <span>运行状态</span>
          <el-button link type="primary" @click="checkHealth">重新检查</el-button>
        </div>
      </template>
      <div v-if="health" class="status-row">
        <span class="status-dot is-online"></span>
        <span>后端与数据库运行正常</span>
        <el-tag size="small" type="success">{{ health.version }}</el-tag>
      </div>
      <div v-else-if="error" class="status-row">
        <span class="status-dot is-error"></span>
        <span>{{ error }}</span>
      </div>
      <div v-else class="status-row">
        <el-skeleton :rows="1" animated />
      </div>
    </el-card>
  </div>
</template>
