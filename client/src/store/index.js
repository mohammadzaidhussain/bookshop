import Vuex from 'vuex';
import * as authStore from './auth-store/index.js';
import * as bookStore from './book-store/index.js';
import * as commonStore from './common-store/index.js';
import * as staffStore from './staff-store/index.js';


export const store = new Vuex.Store({
    modules: {
        authStore: {
            ...authStore,
            namespaced: true
        },
        bookStore: {
            ...bookStore,
            namespaced: true
        },
        commonStore: {
            ...commonStore,
            namespaced: true
        },
        staffStore: {
            ...staffStore,
            namespaced: true
        }
    },
});