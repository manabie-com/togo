package com.manabie.assignment.controllers;

import com.manabie.assignment.models.Task;
import com.manabie.assignment.models.TasksPerDay;
import com.manabie.assignment.models.User;
import com.manabie.assignment.repositories.TaskRepository;
import com.manabie.assignment.repositories.TasksPerDayRepository;
import com.manabie.assignment.repositories.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import java.util.Calendar;
import java.sql.Date;
import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("task")
public class TaskController {
    @Autowired
    private TaskRepository taskRepository;
    @Autowired
    private TasksPerDayRepository tasksPerDayRepository;
    @Autowired
    private UserRepository userRepository;

    @PostMapping("/add")
    public String addTask(@RequestParam(name = "userId") Integer userId,
                          @RequestParam(name = "taskLimit", required = false) Integer taskLimit,
                          @RequestBody Task task) {
        Optional<User> user = userRepository.findById(userId);
        if (user.isEmpty()) {
            return "Not found a user with id " + userId;
        }

        Date currentDate = new Date(Calendar.getInstance().getTime().getTime());
        List<Task> currentTasks = taskRepository.findAllByUserIdAndDate(user.get(), currentDate);
        TasksPerDay tasksPerDay = tasksPerDayRepository.findByUserIdAndDate(user.get(), currentDate);
        if (taskLimit == null) {
            if (tasksPerDay == null) {
                return String.format("Need to set the limit of tasks per day for user %s first", userId);
            }
            if (currentTasks.size() >= tasksPerDay.getTaskLimit()) {
                return String.format("Could not add task more for user %s since the limit of tasks is %s",
                        userId, tasksPerDay.getTaskLimit());
            }
        } else {
            if (taskLimit < currentTasks.size()) {
                return String.format("Could not update the limit of tasks to %s for user %s since the number of current tasks is %s",
                        taskLimit, userId, currentTasks.size());
            }
            if (tasksPerDay != null) {
                if (currentTasks.size() >= taskLimit) {
                    return String.format("Could not add task more for user %s since the limit of tasks is %s",
                            userId, tasksPerDay.getTaskLimit());
                }
            } else {
                tasksPerDay = new TasksPerDay();
                tasksPerDay.setTaskDate(currentDate);
                tasksPerDay.setUser(user.get());
            }

            tasksPerDay.setTaskLimit(taskLimit);
            tasksPerDayRepository.save(tasksPerDay);
        }

        task.setUser(user.get());
        task.setTaskDate(currentDate);
        taskRepository.save(task);
        return "Saved";
    }

    @GetMapping("/all")
    public Iterable<Task> getAll() {
        return taskRepository.findAll();
    }
}
