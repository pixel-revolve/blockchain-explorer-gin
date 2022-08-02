import request from '@/utils/request'
import qs from 'qs'

export function getChainCodeList() {
    return request({
        url: '/blockchain/getInstalledCC',
        method: 'get'
    })
}

