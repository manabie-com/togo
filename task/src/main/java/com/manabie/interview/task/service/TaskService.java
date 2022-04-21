package com.manabie.interview.task.service;

import com.manabie.interview.task.model.Task;
import com.manabie.interview.task.model.User;
import com.manabie.interview.task.repository.TaskRepository;
import com.manabie.interview.task.repository.UserRepository;
import com.manabie.interview.task.response.APIResponse;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class TaskService {

    private final TaskRepository taskRepository;
    private final UserRepository userRepository;

    @Autowired
    public TaskService(TaskRepository taskRepository, UserRepository userRepository) {
        this.taskRepository = taskRepository;
        this.userRepository = userRepository;
    }

    public List<Task> getTasks(){
        return taskRepository.findAll();
    }

    public APIResponse registerNewTask(Task task) {
        Optional<User> userOptional= userRepository.findUserById(task.getUserUid());
        if(!userOptional.isPresent()){
            return new APIResponse(String.format("User %s doesn't exist. Can't not assign task", task.getUserUid()), HttpStatus.OK);
        }
        List<Task> taskList = taskRepository.findDailyTaskByUserId(task.getUserUid(), task.getCreatedDate());
        if(taskList.size() >= userOptional.get().getMaxTask()){
            return new APIResponse(String.format("User %s reaches limit daily task. Can't not assign task", task.getUserUid()), HttpStatus.OK);
        }
        taskRepository.save(task);
        return new APIResponse("Assigned Successfully", HttpStatus.OK);
    }
}
