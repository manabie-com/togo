package com.manabie.todotaskapplication.common.constant;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
public enum TaskActionType {
    CREATE("CREATE"),
    UPDATE("UPDATE");

    private String value;

    TaskActionType(String value) {
        this.value = value;
    }

    public String getValue() {
        return value;
    }
}
