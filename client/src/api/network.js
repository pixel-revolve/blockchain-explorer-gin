import request from '@/utils/request'

export function getPeerList() {
    return request({
        url: '/blockchain/getPeerList',
        method: 'get'
    })
}

