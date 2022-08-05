package com.example.server.infrastructure.task

import com.example.domain.model.TaskEntity
import org.springframework.data.annotation.Id
import org.springframework.data.relational.core.mapping.Column
import org.springframework.data.relational.core.mapping.Table
import org.springframework.data.annotation.Transient
import org.springframework.data.domain.Persistable
import java.time.LocalDateTime
import java.util.*

@Table("tasks")
data class Task(
    @Id
    @Column("id")
    var uid: String = UUID.randomUUID().toString(),
    @Column("title")
    val title: String,
    @Column("description")
    val description : String,
    @Column("user_id")
    val userId: String,
    @Column("created_at")
    var createdAt: LocalDateTime = LocalDateTime.now(),
    @Column("updated_at")
    var updatedAt: LocalDateTime = LocalDateTime.now(),
) : Persistable<String> {

    @Transient
    internal var isNewTask: Boolean = true

    companion object {
        fun fromDomainModel(entity: TaskEntity): Task = Task(
            uid = entity.id,
            title = entity.title,
            description =  entity.description,
            userId = entity.userId
        )
    }

    override fun getId(): String = uid

    override fun isNew(): Boolean = isNewTask
}

fun Task.toDomainModel() : TaskEntity = TaskEntity(
    id = uid,
    title = title,
    description = description,
    userId = userId
)