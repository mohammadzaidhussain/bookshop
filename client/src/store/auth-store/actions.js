import { Api } from "../../api/api.js"

export const actions = {

    async login({ commit, dispatch }, loginData) {
        const res = await Api.login({
            data: loginData.data,
            params: loginData.params,
            headers: loginData.headers,
            appOptions: loginData.appOptions
        });

        commit('SET_LOGGED_IN_USER', res?.data);
        
    },


    async staffLogin({ commit, dispatch }, loginData) {
        const res = await Api.staffLogin({
            data: loginData.data,
            params: loginData.params,
            headers: loginData.headers,
            appOptions: loginData.appOptions
        });

        commit('SET_LOGGED_IN_USER', res?.data);
        
    },

}