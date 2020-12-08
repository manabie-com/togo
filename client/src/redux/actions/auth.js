import authApi from '../../api/authApi';
import * as actionTypes from './actionTypes';

const EXPIRE_DURATION = 900000

export const authStart = () => {
    return {
        type: actionTypes.AUTH_START
    };
};

export const authSuccess = (token) => {
    return {
        type: actionTypes.AUTH_SUCCESS,
        idToken: token,
    };
};

export const authFail = (error) => {
    return {
        type: actionTypes.AUTH_FAIL,
        error: error
    };
};

export const logout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('receiveTime');
    return {
        type: actionTypes.AUTH_LOGOUT
    };
};  

export const auth = (name, password) => {
    return async (dispatch) => {
        dispatch(authStart());
        const authData = {
          name: name,
          password: password
        }
        try { 
          const authRes = authApi.post(authData)
          localStorage.setItem('token', authRes.data.idToken);
          localStorage.setItem('tokenReceive', new Date().getTime());
          dispatch(authSuccess(authRes.idToken));
        } catch(err) {
          dispatch(authFail(err));
        }
    };
};

export const authCheckState = () => {
    return dispatch => {
        const token = localStorage.getItem('token');
        if (!token) {
            dispatch(logout());
        } else {
            const timeReceive = localStorage.getItem('timeReceive');
            const now = new Date().getTime()
            const isExpired = (now -timeReceive) > EXPIRE_DURATION
            if (isExpired){
                dispatch(logout());
            } else {
                dispatch(authSuccess(token));
                dispatch(checkAuthTimeout(timeReceive + EXPIRE_DURATION - now));
            }   
        }
    };
};

export const checkAuthTimeout = (expirationTime) => {
  return dispatch => {
      setTimeout(() => {
          dispatch(logout());
      }, expirationTime);
  };
};

export const setAuthRedirectPath = (path) => {
  return {
      type: actionTypes.SET_AUTH_REDIRECT_PATH,
      path: path
  };
};