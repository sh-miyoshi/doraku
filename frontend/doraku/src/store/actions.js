export const setInternalServerError = (errorMsg) => ({
    type: 'SET_INTERNAL_SERVER_ERROR',
    errorMsg
})

export const setRecommendHobby = (hobby_id, hobby_name) => ({
    type: 'SET_RECOMMEND_HOBBY',
    hobby_id,
    hobby_name
})
