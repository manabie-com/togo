package com.todo.ws;


import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.todo.ws.commons.model.ResponseEntity;

@RequestMapping("/api/v1")
@RestController
public class PublicController {

    @GetMapping("/version")
    public ResponseEntity<?> getVersion() {
        return new ResponseEntity(null, "1.0.0");
    }

}
