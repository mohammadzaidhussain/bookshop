import axios from "axios";
import storageService from "../service/local_storage_service";


export const setupVueInterceptor = (store, router) => {

    axios.interceptors.request.use(
        (config) => {
            const userData = storageService.getPersistedData("logged_in_user");
            config.headers["authorization"] = 'Bearer ' + userData?.token;
            return config;
        },
        (error) => {
            return Promise.reject(error)
        }
    );

    axios.interceptors.response.use((response) => {
        const responseCode = response?.data.statusCode;
        switch (responseCode) {
            case 200:
                if(response?.data?.operation == "delete") {
                    store.dispatch('commonStore/toggleToast', {
                        message: 'Deleted successfully',
                        visible: true,
                        type: 'success',
                    });
                }
                break;

            case 201:
                store.dispatch('commonStore/toggleToast', {
                    message: response.data?.message || 'success',
                    visible: true,
                    type: 'success',
                });
                break;
        
            default:
                break;
        }
        return Promise.resolve(response?.data);
    }, async (error) => {
        const responseCode = error.response?.data?.statusCode;
        const message = error.response?.data?.message;
        switch (responseCode) {
            case 401: {
                store.dispatch('commonStore/toggleToast', {
                    message: message || 'Unauthorized Access!',
                    visible: true,
                    type: 'error'
                });
                store.commit('CLEAR_DATA');
                const route = await router.resolve()
                if(!route.fullPath.includes('/login')) {
                    router.push('/login');
                }
                break;
            }
            case 500: {
                store.dispatch('commonStore/toggleToast', {
                    message: message || 'Internal server error!',
                    visible: true,
                    type: 'error'
                });
                break;
            }
        
            default: {
                store.dispatch('commonStore/toggleToast', {
                    message: 'Something went wrong!',
                    visible: true,
                    type: 'error'
                });
                break;
            }
        }
        return Promise.reject(error);
    })
}