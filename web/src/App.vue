<template>
  <el-config-provider>
    <div class="common-layout">
      <el-container style="height: 100vh;">
        <el-aside width="200px" style="background-color: #304156; transition: width 0.3s;">
          <div class="logo-container">
            <img src="@/assets/logo.svg" alt="logo" class="logo" v-if="false" />
            <span class="logo-title">My-CEX Admin</span>
          </div>
          <el-menu
              active-text-color="#ffd04b"
              background-color="#304156"
              text-color="#fff"
              default-active="1"
              class="el-menu-vertical"
          >
            <el-menu-item index="1">
              <el-icon><Money /></el-icon>
              <span>资产管理</span>
            </el-menu-item>
          </el-menu>
        </el-aside>

        <el-container>
          <el-header class="main-header">
            <div class="header-left">
              <h2>My-CEX 后台管理系统</h2>
            </div>
            <div class="header-right">
              <el-avatar :size="32" src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" />
              <span class="user-name">管理员</span>
            </div>
          </el-header>

          <el-main class="main-content">
            <AssetView />
          </el-main>
        </el-container>
      </el-container>
    </div>
  </el-config-provider>
</template>

<script setup>
import AssetView from './views/AssetView.vue'
</script>

<style>
body {
  margin: 0;
  padding: 0;
  font-family: 'Helvetica Neue', Helvetica, 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
}

.logo-container {
  height: 60px;
  line-height: 60px;
  text-align: center;
  color: #fff;
  font-size: 18px;
  font-weight: bold;
  background-color: #2b2f3a;
}

.main-header {
  background-color: #fff;
  border-bottom: 1px solid #e6e6e6;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 1px 4px rgba(0,21,41,.08);
  padding: 0 20px;
  color: #000;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 10px;
}

.user-name {
  font-size: 14px;
  color: #606266;
}

.main-content {
  background-color: #f0f2f5;
  padding: 0;
}

.el-menu {
  border-right: none;
}
</style>