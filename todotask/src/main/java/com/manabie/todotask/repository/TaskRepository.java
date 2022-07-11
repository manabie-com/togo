package com.manabie.todotask.repository;

import com.manabie.todotask.entity.Task;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import java.time.ZonedDateTime;
import java.util.Date;

@Repository
public interface TaskRepository extends CrudRepository<Task, Integer> {
    @Query("SELECT COUNT (t) FROM Task t where t.userId = :userId and t.targetDate >= :startDate and t.targetDate < :endDate")
    int countTaskByUserIdAndTargetDate(Integer userId, ZonedDateTime startDate, ZonedDateTime endDate);
}
