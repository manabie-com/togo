package com.manabie.assignment.repositories;

import com.manabie.assignment.models.User;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.CrudRepository;

import com.manabie.assignment.models.TasksPerDay;
import org.springframework.data.repository.query.Param;

import java.sql.Date;

public interface TasksPerDayRepository extends CrudRepository<TasksPerDay, Integer> {

    @Query("SELECT t FROM TasksPerDay t WHERE t.user = :user and t.taskDate = :taskDate")
    TasksPerDay findByUserIdAndDate(@Param("user") User user, @Param("taskDate") Date taskDate);
}
