import * as actionTypes from '../actions/actionTypes';

const initialState = {
  tasks: [],
  error: null,
  loading: false
};

const fetchTasksStart = ( state, action ) => ({...state, loading: true });

const fetchTasksSuccess = ( state, action ) => ({...state, tasks: action.tasks, loading: false});

const fetchTasksFail = ( state, action ) => ({...state, loading: false})

const addTaskStart = ( state, action ) => ({state, loading: true});

const addTaskSuccess = ( state, action ) => (
  {state, tasks: [...state.tasks, action.task], loading: false});

const addTaskFail = ( state, action ) => ({...state, loading: false});

const reducer = ( state = initialState, action ) => {
  switch ( action.type ) {
      case actionTypes.FETCH_TASKS_START: return fetchTasksStart(state, action);
      case actionTypes.FETCH_TASKS_SUCCESS: return fetchTasksSuccess(state, action);
      case actionTypes.FETCH_TASKS_FAIL: return fetchTasksFail(state, action);
      case actionTypes.ADD_TASK_START: return addTaskStart(state, action);
      case actionTypes.ADD_TASK_SUCCESS: return addTaskSuccess(state, action);
      case actionTypes.ADD_TASK_FAIL: return addTaskFail(state, action);
      default:
          return state;
  }
};

export default reducer;