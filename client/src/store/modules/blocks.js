import {getBlockList} from "@/api/blocks";

const state = {
    blockList:[]
}

const mutations={
    SET_BLOCK_LIST: (state, blockList) => {
        state.blockList = blockList
    },
}

const actions = {
    getBlockList({
                         commit
                     }) {
        return new Promise((resolve, reject) => {
            getBlockList().then(response => {
                let {
                    data
                } = response.data
                // 直接调用当前模块的mutation
                commit('SET_BLOCK_LIST', data)
                resolve(data)
            }).catch(error => {
                reject(error)
            })
        })
    },
}

export default {
    namespaced: true,
    state,
    mutations,
    actions
}