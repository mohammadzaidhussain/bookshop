
const StorageService = () => {

    const persistData = (key, value) => {
        localStorage.setItem(key, JSON.stringify(value));
    }

    const getPersistedData = (key) => {
        const data = localStorage.getItem(key);
        if(data) {
            return JSON.parse(data);
        }
        return null;
    }

    const clearAll = () => {
        localStorage.clear();
    }


    return {
        persistData,
        getPersistedData,
        clearAll
    }

}

export default StorageService();