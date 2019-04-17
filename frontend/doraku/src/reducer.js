import { combineReducers } from 'redux';
import { connectRouter } from 'connected-react-router';

const initState = {
  error: "",
}

const hobbyReducer = (state = initState, action) => {
  switch (action.type) {
    case 'SET_ERROR':
      return Object.assign({}, state, {
        error: action.error,
      })
    default:
      return state
  }
}

const rootReducer = (history) => combineReducers({
  hobby: hobbyReducer,
  router: connectRouter(history)
})

export default rootReducer