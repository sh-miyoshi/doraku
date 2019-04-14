export const setHobby = (id, name) => ({
    type: 'SET_HOBBY',
    id,
    name
})

export const setError = (error) => ({
    type: 'SET_ERROR',
    error
})