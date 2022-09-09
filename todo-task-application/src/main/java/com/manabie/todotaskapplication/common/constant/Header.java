package com.manabie.todotaskapplication.common.constant;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
public enum Header {
    AUTHORIZATION("Authorization");

    private String value;

    Header(String value) {
        this.value = value;
    }

    public String getValue() {
        return value;
    }
}
