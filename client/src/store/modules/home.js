import {
    getChannelNameList,
    getHighestBlock,
    getPeerList, getPeerNameList,
} from '@/api/home'

const state ={
    peerNameList: [],
    channelNameList: [],
    cardList: [
        {
            title: "BLOCKS",
            number: 0,
            icon: "block"
        },
        {
            title: "TRANSACTIONS",
            number: 0,
            icon: "transaction"
        },
        {
            title: "NODES",
            number: 0,
            icon: "node"
        },
        {
            title: "CHAINCODES",
            number: 0,
            icon: "chaincode"
        }
    ]
}

const mutations={
    SET_PEER_NAME_LIST: (state, peerNameList) => {
        state.peerNameList = peerNameList
    },
    SET_CHANNEL_NAME_LIST: (state, channelNameList) => {
        state.channelNameList = channelNameList
    },
    SET_CARDLIST: (state, cardList) => {
        state.cardList = cardList
    },
    SET_CARD_NUMBER: (state, data) => {
        state.cardList[data.index]["number"] = data.number
    },
}

const actions = {
    // 获取Peer节点列表
    getPeerNameList({
                commit
            }) {
        return new Promise((resolve, reject) => {
            getPeerNameList().then(response => {
                let {
                    data
                } = response.data
                let count=0
                for (const index in data) {
                    if (data[index].search("peer")!==-1)
                        count++
                }
                commit('SET_CARD_NUMBER',{index: 2,number: count})
                data = data.map(item=>({peerName:item}))
                commit('SET_PEER_NAME_LIST', data)
                resolve(data)
            }).catch(error => {
                reject(error)
            })
        })
    },
    getChannelNameList({
                    commit
                }) {
        return new Promise((resolve, reject) => {
            getChannelNameList().then(response => {
                let {
                    data
                } = response.data
                data = data.map(item=>({channelName:item}))
                commit('SET_CHANNEL_NAME_LIST', data)
                resolve(data)
            }).catch(error => {
                reject(error)
            })
        })
    },
    getHighestBlock({
                       commit
                   }) {
        return new Promise((resolve, reject) => {
            getHighestBlock().then(response => {
                let {
                    data
                } = response.data
                commit('SET_CARD_NUMBER', {index:0,number: data.Height})
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