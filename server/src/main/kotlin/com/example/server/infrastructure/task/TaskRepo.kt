package com.example.server.infrastructure.task

import org.springframework.data.r2dbc.repository.Query
import org.springframework.data.repository.kotlin.CoroutineCrudRepository
import org.springframework.stereotype.Repository

@Repository
interface TaskRepo : CoroutineCrudRepository<Task, Long> {

    @Query(
        """
        SELECT count(id)
        FROM tasks
        WHERE DATE(created_at) = DATE(CURRENT_TIMESTAMP)
        AND user_id = :userId
    """
    )
    suspend fun getTaskInDayByUserId(userId: String): Long
}
