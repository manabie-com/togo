package com.manabie.todo.mapper;

import com.manabie.todo.domain.User;
import com.manabie.todo.domain.User.UserBuilder;
import com.manabie.todo.entity.UserEntity;
import javax.annotation.processing.Generated;

@Generated(
    value = "org.mapstruct.ap.MappingProcessor",
    date = "2022-05-13T02:36:02+0700",
    comments = "version: 1.4.2.Final, compiler: javac, environment: Java 17.0.2 (Oracle Corporation)"
)
public class UserMapperImpl implements UserMapper {

    @Override
    public User toDto(UserEntity userEntity) {
        if ( userEntity == null ) {
            return null;
        }

        UserBuilder user = User.builder();

        user.id( userEntity.getId() );
        user.username( userEntity.getUsername() );
        user.password( userEntity.getPassword() );
        user.taskQuote( userEntity.getTaskQuote() );
        user.isDelete( userEntity.getIsDelete() );

        return user.build();
    }
}
