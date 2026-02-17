<template>
  <div class="asset-view-wrapper">
    <el-row :gutter="20">
      <el-col :xs="24" :md="16" class="mb-20">
        <el-card shadow="never" class="custom-card">
          <template #header>
            <div class="card-title">资产管理控制台</div>
          </template>

          <el-form :inline="true" class="search-form">
            <el-form-item label="用户 UID">
              <el-input-number v-model="form.userId" :min="1" controls-position="right" />
            </el-form-item>
            <div class="button-group">
              <el-button type="primary" @click="fetchBalance">查询刷新</el-button>
              <el-button type="success" @click="handleDeposit">一键充值 10 ETH</el-button>
              <el-button type="danger" @click="handleWithdraw" :loading="withdrawLoading">安全提现 5 ETH</el-button>
            </div>
          </el-form>

          <el-table :data="tableData" border stripe class="mt-20" v-loading="loading">
            <el-table-column prop="uid" label="用户 ID" width="80" align="center" />
            <el-table-column prop="address" label="钱包地址" min-width="180" show-overflow-tooltip />
            <el-table-column prop="balance" label="当前余额" width="130" align="right">
              <template #default="scope">
                <span class="balance-num">{{ scope.row.balance }}</span> <small>ETH</small>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>

      <el-col :xs="24" :md="8">
        <el-card class="log-card custom-card" shadow="never">
          <template #header>
            <div class="log-header">
             <div class="box">
               <div class="status-dot"></div>
               <span>系统实时流水审计</span>
             </div>
              <el-button class="ml-10" @click="fetchHistoryLogs">刷新流水</el-button>
            </div>
          </template>

          <div class="timeline-wrapper">
            <el-timeline v-if="logs.length > 0">
              <el-timeline-item
                  v-for="(log, index) in logs"
                  :key="index"
                  :color="log.type === 'DEPOSIT' ? '#67c23a' : '#f56c6c'"
                  :timestamp="log.time"
                  hollow
              >
                <div class="log-text">{{ log.message }}</div>
              </el-timeline-item>
            </el-timeline>
            <el-empty v-else :image-size="60" description="暂无实时动态" />
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { getBalance, deposit, getSystemLogs, withdraw } from '@/api/account'
import { ElMessage } from 'element-plus'

const loading = ref(false)
const form = ref({ userId: 10010 })
const tableData = ref([])
const logs = ref([])
const withdrawLoading = ref(false)

let socket = null
const initWebSocket = () => {
  socket = new WebSocket('ws://localhost:8888/ws')
  socket.onmessage = (event) => {
    const data = JSON.parse(event.data)
    const newLog = {
      type: data.type,
      userId: data.userId || data.uid,
      message: data.message || data.content,
      time: data.time || data.createdAt
    }

    logs.value.unshift(newLog)

    if (logs.value.length > 15) logs.value.pop()

    if (newLog.userId === form.value.userId) {
      fetchBalance()
    }
  }
  socket.onclose = () => {
    setTimeout(initWebSocket, 3000)
  }
}

// fetchHistoryLogs 获取历史流水
const fetchHistoryLogs = async () => {
  try {
    const res = await getSystemLogs({ limit: 10 })
    logs.value = res.list.map(item => ({
      type: item.opType,
      userId: item.uid,
      message: item.content,
      time: item.createdAt
    }))
  } catch (err) {}
}

// handleWithdraw 提现
const handleWithdraw = async () => {
  if (parseFloat(tableData.value[0]?.balance || 0) < 5) {
    return ElMessage.error('链上余额不足，无法提现')
  }

  try {
    withdrawLoading.value = true
    const res = await withdraw({
      userId: form.value.userId,
      amount: "5" // 每次提现 5 枚
    })

    if (res.ok) {
      ElMessage({
        message: `提现交易已发送！Hash: ${res.txHash.slice(0, 10)}...`,
        type: 'success',
        duration: 5000
      })
    }
  } catch (err) {
    ElMessage({
      message: `提现交易失败${err}`,
      type: 'error',
      duration: 5000
    })
  } finally {
    withdrawLoading.value = false
  }
}

const fetchBalance = async () => {
  loading.value = true
  try {
    const data = await getBalance({ userId: form.value.userId, currency: 'ETH' })
    tableData.value = [{
      uid: form.value.userId,
      address: data.address,
      balance: data.balance
    }]
  } finally {
    loading.value = false
  }
}

const handleDeposit = async () => {
  try {
    await deposit({ userId: form.value.userId, amount: "10" })
    ElMessage.success('充值成功')
  } catch (err) {}
}

onMounted(() => {
  fetchBalance()
  fetchHistoryLogs()
  initWebSocket()
})

onUnmounted(() => {
  if (socket) socket.close()
})
</script>

<style scoped>
.asset-view-wrapper {
  padding: 20px;
  box-sizing: border-box;
  width: 100%;
}

.custom-card {
  border-radius: 8px;
  border: 1px solid #ebeef5;
}

.card-title {
  font-weight: 600;
  color: #303133;
}

.mt-20 { margin-top: 20px; }
.mb-20 { margin-bottom: 20px; }

.search-form {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.button-group {
  display: flex;
  gap: 10px;
}

.balance-num {
  font-weight: bold;
  color: #67c23a;
}

.log-card {
  display: flex;
  flex-direction: column;
  height: calc(100vh - 100px);
}

.timeline-wrapper {
  flex: 1;
  overflow-y: auto;
  padding-right: 10px;
}

.log-header {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: bold;
  justify-content: space-between;
}

.box {
  display: flex;
  align-items: center;
  gap: 4px;
}

.log-text {
  font-size: 13px;
  color: #606266;
  line-height: 1.5;
}

/* 呼吸灯 */
.status-dot {
  width: 8px;
  height: 8px;
  background-color: #67c23a;
  border-radius: 50%;
  box-shadow: 0 0 5px #67c23a;
  animation: blink 2s infinite;
  margin-right: 10px;
}

@keyframes blink {
  50% { opacity: 0.3; }
}

@media (max-width: 768px) {
  .asset-view-wrapper {
    padding: 10px;
  }

  .log-card {
    height: 500px;
    margin-bottom: 20px;
  }

  .el-form-item {
    margin-right: 0 !important;
    width: 100%;
  }

  .button-group {
    width: 100%;
    justify-content: space-between;
  }

  .button-group .el-button {
    flex: 1;
  }
}
</style>