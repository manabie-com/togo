
import * as actionTypes from './actionTypes';
import taskApi from '../../api/taskApi';

export const fetchTasksSuccess = ( tasks ) => {
    return {
        type: actionTypes.FETCH_TASKS_SUCCESS,
        tasks: tasks
    };
};

export const fetchTasksFail = ( error ) => {
    return {
        type: actionTypes.FETCH_TASKS_FAIL,
        error: error
    };
};

export const fetchTasksStart = () => {
    return {
        type: actionTypes.FETCH_TASKS_START
    };
};

export const fetchTasks = (token) => {
    return async dispatch => {
        dispatch(fetchTasksStart());
        try {
          const tasks = await taskApi.getAll(token)
          dispatch(fetchTasksSuccess(tasks));
        } catch(err) {
          dispatch(fetchTasksFail(err));
        }
    };
};

export const addTaskSuccess = ( task ) => {
  return {
      type: actionTypes.ADD_TASK_SUCCESS,
      tasks: task
  };
};

export const addTaskFail = ( error ) => {
  return {
      type: actionTypes.ADD_TASK_FAIL,
      error: error
  };
};

export const addTaskStart = () => {
  return {
      type: actionTypes.ADD_TASK_START
  };
};

export const addTask = (token, task) => {
  return async dispatch => {
    dispatch(addTaskStart());
    try {
      const tasks = await taskApi.add(token, task)
      dispatch(addTaskSuccess(tasks));
    } catch(err) {
      dispatch(addTaskFail(err));
    }
  };
};