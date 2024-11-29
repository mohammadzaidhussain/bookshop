import storageService from "../../service/local_storage_service.js";
import { getDefaultState } from "./state.js";

export const mutations = {

    SET_LOGGED_IN_USER: (state, data) => {
        if(!data) {
            return;
        } 
        state.logged_in_user = {
            ...state.logged_in_user,
            ...data
        }
        if(Object.keys(state.logged_in_user).length > 0) {
            storageService.persistData("logged_in_user", state.logged_in_user);
        }
    },

    SET_USER_ID: (state, user_id) => {
        state.user_id = user_id;
    },

    CLEAR_DATA: (state) => {
        state.logged_in_user = null;
        storageService.clearAll();
    }

}