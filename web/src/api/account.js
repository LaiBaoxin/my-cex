import request from '@/utils/request'

// 查询余额
export function getBalance(data) {
    return request({
        url: '/account/balance',
        method: 'post',
        data
    })
}

// 充值
export function deposit(data) {
    return request({
        url: '/account/deposit',
        method: 'post',
        data
    })
}

// 提现
export function withdraw(data) {
    return request({
        url: '/account/withdraw',
        method: 'post',
        data
    })
}

// getSystemLogs 查询系统日志
export function getSystemLogs(params) {
    return request({
        url: '/account/logs',
        method: 'get',
        params
    })
}