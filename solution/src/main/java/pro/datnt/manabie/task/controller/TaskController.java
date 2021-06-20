package pro.datnt.manabie.task.controller;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.AccessDeniedException;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.security.core.annotation.CurrentSecurityContext;
import org.springframework.security.core.context.SecurityContext;
import org.springframework.web.bind.annotation.*;
import pro.datnt.manabie.task.controller.model.TaskDTO;
import pro.datnt.manabie.task.model.TaskDBO;
import pro.datnt.manabie.task.repository.TaskRepository;
import pro.datnt.manabie.task.service.TaskService;

@RestController
@RequiredArgsConstructor
@RequestMapping("/task")
public class TaskController {
    private final TaskService taskService;

    @PostMapping("/add")
    public ResponseEntity<?> createTask(@RequestBody TaskDTO task, @AuthenticationPrincipal Long userId) {
        try {
            TaskDBO savedTask = taskService.createTask(task.getContent(), userId);
            return ResponseEntity.accepted()
                    .body(savedTask);
        } catch (AccessDeniedException e) {
            return ResponseEntity.unprocessableEntity().build();
        }
    }

    @GetMapping("/list")
    public ResponseEntity<?> listTask(@AuthenticationPrincipal Long userId) {
        return ResponseEntity.ok(taskService.list(userId));
    }
}
