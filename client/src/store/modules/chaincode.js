import {getChainCodeList} from "@/api/chaincode";

const state = {
    chainCodeList:[]
}

const mutations={
    SET_CHAINCODE_LIST: (state, chainCodeList) => {
        state.chainCodeList = chainCodeList
    },
}

const actions = {
    getChainCodeList({
                    commit
                }) {
        return new Promise((resolve, reject) => {
            getChainCodeList().then(response => {
                let {
                    data
                } = response.data
                // 给home的card添加数据
                commit('home/SET_CARD_NUMBER',{index: 3,number: data.length},{root: true})
                for (const index in data) {
                    let references=data[index].references
                    let channelName=Object.keys(references)[0]
                    data[index]['channelName']=channelName
                    let chaincodeDetail=references[channelName][0]
                    data[index]['chaincodeName']=chaincodeDetail.name
                    data[index]['version']=chaincodeDetail.version
                    data[index]['path']='-'
                }

                // 直接调用当前模块的mutation
                commit('SET_CHAINCODE_LIST', data)
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