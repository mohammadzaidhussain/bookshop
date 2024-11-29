export const mutations = {

    SET_TOAST_STATE(state, payload) {
        state.toast = {
            ...state.toast,
            ...payload
        };
    }

}