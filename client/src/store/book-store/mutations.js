import storageService from "../../service/local_storage_service.js";

export const mutations = {

    SET_BOOKS: (state, books) => {
        state.books = books;
        storageService.persistData("books", books);
    },

    SET_TOTAL_BOOK_RECORD: (state, total_records) => {
        state.total_book_records = total_records;
    },

    SET_BOOK: (state, book) => {
        state.books = [
            ...state.books,
            book
        ];
        storageService.persistData("books", state.books);
    },

    SET_OPENED_BOOK: (state, book) => {
        state.opened_book = {
            ...state.opened_book,
            ...book
        };
        storageService.persistData("opened_book", state.opened_book);
    },

    DELETE_BOOK_FROM_LIST: (state, bookId) => {
        const bookIndex = state.books.findIndex((b) => b._id == bookId);
        state.books.splice(bookIndex, 1);
        storageService.persistData("books", state.books);
    }

}