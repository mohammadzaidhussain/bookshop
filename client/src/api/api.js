import { create_book_url, create_staff_url, delete_book_url, delete_staff_url, get_book_by_id_url, get_books_url, get_staff_by_id_url, get_staffs_url, login_url, staff_login_url, update_book_url, update_staff_url } from "./url.js";

import axios from 'axios';

import { GetApiClient } from "./api-client.js";

const apiClient = GetApiClient(axios);

export const Api = {

    login: async ({ data, params, headers, appOptions }) => {
        const res = await apiClient.getInstance().post(login_url, data, params, headers, appOptions);
        return res;
    },

    staffLogin: async ({ data, params, headers, appOptions }) => {
        const res = await apiClient.getInstance().post(staff_login_url, data, params, headers, appOptions);
        return res;
    },

    createBook: async ({ data, params, headers, appOptions }) => {
        const res = await apiClient.getInstance().post(create_book_url, data, params, headers, appOptions);
        return res;
    },

    getBooks: async ({ data, params, headers, appOptions }) => {
        const res = await apiClient.getInstance().post(get_books_url, data, params, headers, appOptions);
        return res;
    },

    updateBook: async ({ data, params, headers, appOptions }) => {
        const id = data._id || data.id;
        const res = await apiClient.getInstance().post(update_book_url.replace('{id}', id), data, params, headers, appOptions);
        return res;
    },

    deleteBook: async ({ data, params, headers, appOptions }) => {
        const id = data._id || data.id;
        const res = await apiClient.getInstance().post(delete_book_url.replace('{id}', id), data, params, headers, appOptions);
        return res;
    },

    getBookById: async ({ data, params, headers, appOptions }) => {
        const id = data._id || data.id;
        const res = await apiClient.getInstance().get(get_book_by_id_url.replace('{id}', id), params, headers, appOptions);
        return res;
    },

    // ===================== staff apis =========================================

    createStaff: async ({ data, params, headers, appOptions }) => {
        const res = await apiClient.getInstance().post(create_staff_url, data, params, headers, appOptions);
        return res;
    },

    getStaffs: async ({ data, params, headers, appOptions }) => {
        const res = await apiClient.getInstance().post(get_staffs_url, data, params, headers, appOptions);
        return res;
    },

    updateStaff: async ({ data, params, headers, appOptions }) => {
        const id = data._id || data.id;
        const res = await apiClient.getInstance().post(update_staff_url.replace('{id}', id), data, params, headers, appOptions);
        return res;
    },

    deleteStaff: async ({ data, params, headers, appOptions }) => {
        const id = data._id || data.id;
        const res = await apiClient.getInstance().post(delete_staff_url.replace('{id}', id), data, params, headers, appOptions);
        return res;
    },

    getStaffById: async ({ data, params, headers, appOptions }) => {
        const id = data._id || data.id;
        const res = await apiClient.getInstance().get(get_staff_by_id_url.replace('{id}', id), params, headers, appOptions);
        return res;
    },

    // =========================== staff api ============================

}