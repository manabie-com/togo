package com.uuhnaut69.app.todo.repository;

import com.uuhnaut69.app.todo.model.Todo;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

/**
 * @author uuhnaut
 */
@Repository
public interface TodoRepository extends JpaRepository<Todo, Long> {

  @Query(value = """
      SELECT COUNT(1) FROM todo
      WHERE user_id = :userId
      AND CAST(created_at AS DATE) = CAST(CURRENT_TIMESTAMP AS DATE)
      """, nativeQuery = true)
  long countNumberOfCreatedTodosTodayByUserId(Long userId);
}
