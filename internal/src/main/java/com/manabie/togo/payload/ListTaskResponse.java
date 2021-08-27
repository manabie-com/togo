/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.manabie.togo.payload;

import com.manabie.togo.model.Task;
import java.util.List;
import lombok.Data;

/**
 * The response of "listTask" handler
 * @author mupmup
 */
@Data
public class ListTaskResponse {
    private List<Task> data;
}
