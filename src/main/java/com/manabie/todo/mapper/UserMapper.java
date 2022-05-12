package com.manabie.todo.mapper;

import com.manabie.todo.domain.User;
import com.manabie.todo.entity.UserEntity;
import org.mapstruct.Mapper;
import org.mapstruct.factory.Mappers;

@Mapper
public interface UserMapper {
  UserMapper INSTANCE = Mappers.getMapper(UserMapper.class);

  User toDto(UserEntity userEntity);
}
