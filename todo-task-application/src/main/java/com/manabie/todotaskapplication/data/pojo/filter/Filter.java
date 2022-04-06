package com.manabie.todotaskapplication.data.pojo.filter;

import lombok.Data;

import java.io.Serializable;
import java.util.Set;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
@Data
public class Filter {
    private int pageSize;
    private int pageIndex;
}
