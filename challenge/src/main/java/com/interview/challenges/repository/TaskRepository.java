package com.interview.challenges.repository;

import java.util.List;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import com.interview.challenges.domain.Task;

@Repository
public interface TaskRepository extends CrudRepository<Task, Long>{
	List<Task> findByUserId(String userId);
}
