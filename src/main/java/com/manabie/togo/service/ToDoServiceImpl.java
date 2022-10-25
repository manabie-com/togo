package com.manabie.togo.service;

import com.manabie.togo.domain.ToDoTask;
import com.manabie.togo.domain.Users;
import com.manabie.togo.dto.ToDoRequest;
import com.manabie.togo.exception.DailyLimitException;
import com.manabie.togo.exception.UserNotFoundException;
import com.manabie.togo.repository.ToDoTaskRepository;
import com.manabie.togo.repository.UsersRepository;
import org.apache.commons.lang3.RandomStringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.List;
import java.util.Optional;

@Component
public class ToDoServiceImpl implements ToDoService{

    @Autowired
    private UsersRepository usersRepository;

    @Autowired
    private ToDoTaskRepository toDoTaskRepository;

    /**
     * Add To Do Task for the user
     * @param toDoRequest
     * @return
     */
    @Override
    public String addToDo(ToDoRequest toDoRequest) {

        // Get user by user id
        Optional<Users> user = usersRepository.findById(toDoRequest.getUserId());

        // Validate if user is existing
        if (!user.isPresent()){
            throw new UserNotFoundException(toDoRequest.getUserId() + " user id not found");
        }

        // Get to do list by user id and to do date
        List<ToDoTask> toDoList = toDoTaskRepository.findByUserIdAndToDoDate(toDoRequest.getUserId(), toDoRequest.getToDoDate());

        // Validate if to do list reaches daily limit
        if (toDoList.size() + 1 > user.get().getDailyLimit()){
            throw new DailyLimitException("Daily Limit of " + user.get().getDailyLimit()+ " To Do Tasks reached");
        }

        // Save to do task
        toDoTaskRepository.save(new ToDoTask(
                RandomStringUtils.randomNumeric(4),
                toDoRequest.getUserId(),
                toDoRequest.getTitle(),
                toDoRequest.getDescription(),
                toDoRequest.getToDoDate()));

        return "To Do Task Added";
    }
}
