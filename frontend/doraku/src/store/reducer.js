import { combineReducers } from 'redux';
import { connectRouter } from 'connected-react-router';

const initErrorState = {
  errorMsg: "",
}

const initRecommendState = {
  hobby_id: 0,
  hobb_name: "",
}

const errorReducer = (state = initErrorState, action) => {
  switch (action.type) {
    case 'SET_INTERNAL_SERVER_ERROR':
      return Object.assign({}, state, {
        errorMsg: action.errorMsg,
      })
    default:
      return state
  }
}

const recommendReducer = (state = initRecommendState, action) => {
  switch (action.type) {
    case 'SET_RECOMMEND_HOBBY':
      return Object.assign({}, state, {
        hobby_id: action.hobby_id,
        hobby_name: action.hobby_name,
      })
    default:
      return state
  }
}

const rootReducer = (history) => combineReducers({
  error: errorReducer,
  recommend: recommendReducer,
  router: connectRouter(history)
})

export default rootReducer