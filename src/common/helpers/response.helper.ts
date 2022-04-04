const success = (data: any = null, message = null) => {
  if (message) {
    return { success: true, message: message, data: data };
  }

  return { success: true, data: data, message: '' };
};

const failed = (message = '') => {
  return {
    data: null,
    message,
    success: false,
  };
};

const AUTH_ERROR = {
  success: false,
  message: 'Sai tài khoản hoặc mật khẩu',
  data: null,
};

const SYSTEM_ERROR = {
  data: null,
  success: false,
  message: 'Có lỗi với hệ thống, vui lòng liên hệ ban quản trị.',
};

const Response = {
  success,
  SYSTEM_ERROR,
  AUTH_ERROR,
  failed,
};

export default Response;
