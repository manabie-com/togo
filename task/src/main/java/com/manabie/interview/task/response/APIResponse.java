package com.manabie.interview.task.response;

import lombok.*;
import org.springframework.http.HttpStatus;

@Getter
@Setter
@AllArgsConstructor
@EqualsAndHashCode
@ToString
public class APIResponse {
    private String errorCode;
    private HttpStatus status;
}
