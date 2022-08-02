import {getPeerList} from "@/api/network";

const state = {
    peerList:[]
}

const mutations={
    SET_PEER_LIST: (state, peerList) => {
        state.peerList = peerList
    },
}

const actions = {
    getPeerList({
                       commit
                   }) {
        return new Promise((resolve, reject) => {
            getPeerList().then(response => {
                let {
                    data
                } = response.data
                // 直接调用当前模块的mutation
                commit('SET_PEER_LIST', data)
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