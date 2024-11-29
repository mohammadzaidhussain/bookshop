export const GetApiClient = (client) => {
    const getInstance = () => {
        return {
            get,
            post,
            patch,
            put,
            remove
        }
    }

    const get = async (url, params = {}, headers = {}, appOptions = {}) => {
        return client.get(url, {
            params,
            headers,
            appOptions
        })
    }

    const post = async (url, data, params = {}, headers = {}, appOptions = {}) => {
        return client.post(url, data, {
            params,
            headers,
            appOptions
        })
    }

    const patch = async (url, data, params = {}, headers = {}, appOptions = {}) => {
        return client.patch(url, data, {
            params,
            headers,
            appOptions
        })
    }

    const put = async (url, data, params = {}, headers = {}, appOptions = {}) => {
        return client.patch(url, data, {
            params,
            headers,
            appOptions
        })
    }

    const remove = async (url, data, params = {}, headers = {}, appOptions = {}) => {
        return client.delete(url, data, {
            params,
            headers,
            appOptions
        })
    }

    return {
        getInstance
    }

}
