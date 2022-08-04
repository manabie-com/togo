package com.example.server.infrastructure.user

import com.example.domain.model.UserEntity
import org.springframework.data.annotation.Id
import org.springframework.data.annotation.Transient
import org.springframework.data.domain.Persistable
import org.springframework.data.relational.core.mapping.Column
import org.springframework.data.relational.core.mapping.Table
import java.util.*

@Table("users")
data class User(
    @Id
    @Column("id")
    var uid: String = UUID.randomUUID().toString(),
    @Column("limit_task")
    var limitTask: Long,
) : Persistable<String> {

    @Transient
    internal var isNewUser: Boolean = true

    override fun getId(): String = uid

    override fun isNew(): Boolean = isNewUser

}

fun User.toDomainModel(): UserEntity = UserEntity(
    id = uid,
    limitTask = limitTask,
)