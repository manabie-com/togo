export class Constants {
    static SUCCESS_CODE = 0;
    static SUCCESS_MESSAGE = "Successful";

    static FAIL_CODE = 1;
    static FAIL_MESSAGE = "Failed";

    static SALT_OR_ROUNDS = 10;

    static FAIL_CHECK = {
        isSuccess: false,
        data: null
    }

    static SUCCESS_CHECK = {
        isSuccess: true,
        data: null
    }

    static USER_ROLE = "USER";

    static ADMIN_ROLE = "ADMIN";

    static MAX_TASK_PER_DAY_DEFAULT = 3;
}
