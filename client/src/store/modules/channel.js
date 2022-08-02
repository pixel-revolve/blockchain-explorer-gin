import {getChannelList} from "@/api/channel";

const state = {
    channelList:[]
}

const mutations={
    SET_CHANNEL_LIST: (state, channelList) => {
        state.channelList = channelList
    },
}

const actions = {
    getChannelList({
                         commit
                     }) {
        return new Promise((resolve, reject) => {
            getChannelList().then(response => {
                let {
                    data
                } = response.data
                // 直接调用当前模块的mutation
                commit('SET_CHANNEL_LIST', data)
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