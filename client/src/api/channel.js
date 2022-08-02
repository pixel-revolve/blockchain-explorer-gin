import request from '@/utils/request'

export function getChannelList() {
    return request({
        url: '/blockchain/getChannelList',
        method: 'get'
    })
}

