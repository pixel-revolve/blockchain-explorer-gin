import {getTransactionList} from "@/api/transaction";

const state = {
    transactionList:[]
}

const mutations={
    SET_TRANSACTION_LIST: (state, transactionList) => {
        state.transactionList = transactionList
    },
}

const actions = {
    getTransactionList({
                         commit
                     }) {
        return new Promise((resolve, reject) => {
            getTransactionList().then(response => {
                let {
                    data
                } = response.data
                // 给home的card添加数据
                commit('home/SET_CARD_NUMBER',{index: 1,number: data.length},{root: true})
                // 直接调用当前模块的mutation
                commit('SET_TRANSACTION_LIST', data)
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