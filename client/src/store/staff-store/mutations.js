import storageService from "../../service/local_storage_service.js";

export const mutations = {

    SET_STAFFS: (state, staffs) => {
        if(!staffs) {
            return;
        }
        state.staffs = staffs;
        storageService.persistData("staffs", staffs);
    },

    SET_STAFF: (state, staff) => {
        if(!staff) {
            return;
        }
        state.staffs = [
            ...state.staffs,
            staff
        ];
        storageService.persistData("staffs", state.staffs);
    },

    SET_OPENED_STAFF: (state, staff) => {
        if(!staff) {
            return;
        }
        state.opened_staff = {
            ...state.opened_staff,
            ...staff
        };
        storageService.persistData("opened_staff", state.opened_staff);
    },

    DELETE_STAFF_FROM_LIST: (state, staffId) => {
        const staffIndex = state.staffs.findIndex((b) => b._id == staffId);
        state.staffs.splice(staffIndex, 1);
        storageService.persistData("staffs", state.staffs);
    }

}