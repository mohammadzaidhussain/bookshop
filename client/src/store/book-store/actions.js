import { Api } from "../../api/api.js";

export const actions = {

    async createBook({ commit, dispatch }, bookData) {

        const book = await Api.createBook({
            data: bookData.data,
            appOptions: bookData.appOptions
        });

        commit('SET_BOOK', book);
        commit('SET_OPENED_BOOK', book?.data);

    },

    async getAllBooks({commit}, filters) {

        const books = await Api.getBooks({
            data: filters.data || {},
        });

        commit('SET_BOOKS', books?.data?.items);
        commit('SET_TOTAL_BOOK_RECORD', books?.data?.total_records);
    },

    async getBookById({commit}, filters) {

        const book = await Api.getBookById({
            data: filters.data,
        });

        commit('SET_OPENED_BOOK', book?.data);
    },

    async updateBookById({commit}, bookData) {

        const book = await Api.updateBook({
            data: bookData.data,
        });

        commit('SET_OPENED_BOOK', book?.data);
    },

    async deleteBookById({commit}, bookData) {

        const book = await Api.deleteBook({
            data: bookData.data,
        });

        commit('DELETE_BOOK_FROM_LIST', bookData.data._id)

    }

}