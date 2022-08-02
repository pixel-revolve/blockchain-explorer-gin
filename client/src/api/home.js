import request from '@/utils/request'
import qs from 'qs'

export function getPeerNameList() {
    return request({
        url: '/blockchain/getPeerNameList',
        method: 'get'
    })
}

export function getChannelNameList() {
    return request({
        url: '/blockchain/getChannels',
        method: 'get'
    })
}

export function getHighestBlock() {
    return request({
        url: '/blockchain/getHighestBlock',
        method: 'get'
    })
}