import { Api } from "../../api/api.js";

export const actions = {

    async createStaff({ commit, dispatch }, staffData) {

        const staff = await Api.createStaff({
            data: staffData.data,
            appOptions: staffData.appOptions
        });

        commit('SET_STAFF', staff);
        commit('SET_OPENED_STAFF', staff?.data);

    },

    async getAllStaffs({commit}, filters) {

        const staffs = await Api.getStaffs({
            data: filters.data || {},
        });

        commit('SET_STAFFS', staffs?.data)
    },

    async getStaffById({commit}, filters) {

        const staff = await Api.getStaffById({
            data: filters.data,
        });

        commit('SET_OPENED_STAFF', staff?.data);
    },

    async updateStaffById({commit}, staffData) {

        const staff = await Api.updateStaff({
            data: staffData.data,
        });

        commit('SET_OPENED_STAFF', staff?.data);
    },

    async deleteStaffById({commit}, staffData) {

        const staff = await Api.deleteStaff({
            data: staffData.data,
        });

        commit('DELETE_STAFF_FROM_LIST', staffData.data._id)

    }

}