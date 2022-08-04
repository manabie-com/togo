package com.example.server.infrastructure.user

import com.example.domain.model.UserEntity
import com.example.domain.usecase.user.UserDateSource
import org.springframework.stereotype.Component
import org.springframework.transaction.annotation.Transactional

@Component
@Transactional
class UserDS(
    private val userRepo: UserRepo,
) : UserDateSource {
    override suspend fun getUser(userId: String): UserEntity? {
        return userRepo.findById(userId)?.toDomainModel()
    }
}
