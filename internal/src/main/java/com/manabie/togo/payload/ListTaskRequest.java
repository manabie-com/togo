package com.manabie.togo.payload;

import javax.validation.constraints.NotBlank;
import lombok.Data;

/**
 * The request of "ListTask" handler
 * @author mupmup
 */
@Data
public class ListTaskRequest {
    @NotBlank
    private String created_date;
}
