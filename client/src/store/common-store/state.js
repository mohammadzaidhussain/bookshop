
export const getDefaultState = () => {
    return {
        toast: {
            visible: false,
            message: '',
            position: 'top',
            defaultTimeoutInMs: 5000,
            type: 'success'
        }
    }
}

export const state = getDefaultState();