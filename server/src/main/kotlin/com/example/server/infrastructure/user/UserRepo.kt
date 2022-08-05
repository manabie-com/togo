package com.example.server.infrastructure.user

import org.springframework.data.repository.kotlin.CoroutineCrudRepository
import org.springframework.stereotype.Repository

@Repository
interface UserRepo: CoroutineCrudRepository<User, String>