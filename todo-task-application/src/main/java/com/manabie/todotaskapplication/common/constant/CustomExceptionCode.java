package com.manabie.todotaskapplication.common.constant;

/**
 * @author @quoctrung.phan
 * @created 04/06/2022
 * @project todo-task-application
 */
public enum CustomExceptionCode {
    RATE_LIMIT_EXCEPTION("RATE_LIMIT_EXCEPTION"),
    VALIDATE_CREATE_TASK_EXCEPTION("VALIDATE_CREATE_TASK_EXCEPTION"),
    VALIDATE_UPDATE_TASK_EXCEPTION("VALIDATE_UPDATE_TASK_EXCEPTION"),
    INTERAL_SERVER_ERROR("INTERAL_SERVER_ERROR");
    private String value;

    CustomExceptionCode(String value) {
        this.value = value;
    }

    public String getValue() {
        return value;
    }

}
