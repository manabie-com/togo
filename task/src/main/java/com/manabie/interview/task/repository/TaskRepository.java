package com.manabie.interview.task.repository;

import com.manabie.interview.task.model.Task;
import com.manabie.interview.task.model.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface TaskRepository extends JpaRepository<Task, String> {
    @Query("SELECT t FROM Task t WHERE t.userUid = ?1 AND t.createdDate = ?2")
    List<Task> findDailyTaskByUserId(String uid, String createDate);
}
