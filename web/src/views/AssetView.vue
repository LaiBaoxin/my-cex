<template>
  <div class="asset-container">
    <el-row :gutter="20" class="stat-row">
      <el-col :xs="24" :sm="12" :lg="8">
        <el-card shadow="hover" class="stat-card blue">
          <div class="stat-info">
            <div class="label">当前系统总用户</div>
            <div class="value">1,284</div>
          </div>
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="12" :lg="8">
        <el-card shadow="hover" class="stat-card green">
          <div class="stat-info">
            <div class="label">查询用户余额 (UID: {{ form.userId }})</div>
            <div class="value">{{ balance }} <span class="unit">ETH</span></div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-card class="list-card">
      <template #header>
        <div class="header-content">
          <div class="title">资产列表管理</div>
          <div class="actions">
            <el-button type="primary" @click="fetchBalance">刷新查询</el-button>
            <el-button type="success" @click="handleDeposit">一键充值 10 ETH</el-button>
          </div>
        </div>
      </template>

      <el-table :data="tableData" border stripe style="width: 100%" v-loading="loading">
        <el-table-column prop="uid" label="用户 ID" width="120" align="center" />
        <el-table-column prop="address" label="钱包地址" min-width="200" show-overflow-tooltip />
        <el-table-column prop="currency" label="币种" width="100" align="center">
          <template #default="scope">
            <el-tag>{{ scope.row.currency }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="balance" label="当前余额" width="180">
          <template #default="scope">
            <span class="balance-text">{{ scope.row.balance }} ETH</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" align="center">
          <template #default>
            <el-button link type="primary">流水明细</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getBalance, deposit } from '@/api/account'

const loading = ref(false)
const form = ref({ userId: 10010 })
const balance = ref('0.0000')
const tableData = ref([])

// 查询余额逻辑
const fetchBalance = async () => {
  loading.value = true
  try {
    const data = await getBalance({
      userId: parseInt(form.value.userId),
      currency: 'ETH'
    })

    balance.value = data.balance
    tableData.value = [{
      uid: form.value.userId,
      address: data.address,
      currency: 'ETH',
      balance: data.balance
    }]
    ElMessage.success('同步成功')
  } catch (err) {
    console.error(err)
  } finally {
    loading.value = false
  }
}

const handleDeposit = async () => {
  try {
    await deposit({
      userId: parseInt(form.value.userId),
      amount: "10"
    })
    ElMessage.success('充值成功')
    await fetchBalance()
  } catch (err) {}
}

onMounted(async () => {
  console.log("初始化执行")
  await fetchBalance()
})
</script>

<style scoped>

.asset-container {
  padding: 20px;
  background-color: #f5f7fa;
  min-height: calc(100vh - 100px);
}

.stat-row {
  margin-bottom: 20px;
}

.stat-card {
  height: 100px;
  display: flex;
  align-items: center;
  color: white;
}

.stat-card.blue { background: linear-gradient(135deg, #3a7bd5, #00d2ff); }
.stat-card.green { background: linear-gradient(135deg, #11998e, #38ef7d); }

.stat-info .label { font-size: 14px; opacity: 0.8; }
.stat-info .value { font-size: 28px; font-weight: bold; margin-top: 5px; }
.unit { font-size: 16px; margin-left: 5px; }

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.title { font-size: 18px; font-weight: bold; color: #303133; }
.balance-text { font-weight: bold; color: #67C23A; }

@media (max-width: 768px) {
  .actions { margin-top: 10px; display: flex; flex-direction: column; gap: 10px; }
  .stat-row .el-col { margin-bottom: 10px; }
}
</style>