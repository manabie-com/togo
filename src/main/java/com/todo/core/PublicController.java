package com.todo.core;

import com.todo.core.commons.model.GenericResponse;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RequestMapping("/api/v1")
@RestController
public class PublicController {

    @GetMapping("/version")
    public GenericResponse getVersion() {
        return new GenericResponse(null, "1.0.0");
    }

}
