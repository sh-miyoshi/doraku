import { combineReducers } from 'redux';
import { connectRouter } from 'connected-react-router';

const initState = {
  errorMsg: "",
}

const errorReducer = (state = initState, action) => {
  switch (action.type) {
    case 'SET_INTERNAL_SERVER_ERROR':
      return Object.assign({}, state, {
        errorMsg: action.errorMsg,
      })
    default:
      return state
  }
}

const rootReducer = (history) => combineReducers({
  error: errorReducer,
  router: connectRouter(history)
})

export default rootReducer