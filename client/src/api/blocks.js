import request from '@/utils/request'
import qs from 'qs'

export function getBlockList() {
    return request({
        url: '/blockchain/getBlockList',
        method: 'get'
    })
}

