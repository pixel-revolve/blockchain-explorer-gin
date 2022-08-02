import request from '@/utils/request'
import qs from 'qs'

export function getTransactionList() {
    return request({
        url: '/blockchain/getTransactionList',
        method: 'get'
    })
}

