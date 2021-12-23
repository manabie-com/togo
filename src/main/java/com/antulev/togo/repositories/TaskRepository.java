package com.antulev.togo.repositories;

import java.util.Date;
import java.util.List;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.rest.core.annotation.RestResource;

import com.antulev.togo.models.Task;

@RestResource(exported=false)
public interface TaskRepository extends JpaRepository<Task, Long>{
	List<Task>findByCreatedDateBetween(Date before, Date after);
}
