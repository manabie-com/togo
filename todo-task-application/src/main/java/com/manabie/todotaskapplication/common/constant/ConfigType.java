package com.manabie.todotaskapplication.common.constant;

/**
 * @author @quoctrung.phan
 * @created 04/06/2022
 * @project todo-task-application
 */
public enum ConfigType {
    RATE_LIMIT_ADD_TASK_PER_DAY("RATE_LIMIT_ADD_TASK_PER_DAY");

    private String value;

    ConfigType(String value) {
        this.value = value;
    }

    public String getValue() {
        return value;
    }
}
