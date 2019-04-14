import { combineReducers } from 'redux'
import { connectRouter } from 'connected-react-router'

const initState = {
  selected_hobby_id: 0,
  selected_hobby_name: "",
  error: "",
}

const hobbyReducer = (state = initState, action) => {
  switch (action.type) {
    case 'SET_HOBBY':
      return Object.assign({}, state, {
        selected_hobby_id: action.selected_hobby_id,
        selected_hobby_name: action.selected_hobby_name,
      })
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