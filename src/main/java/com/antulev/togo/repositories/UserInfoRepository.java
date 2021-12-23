package com.antulev.togo.repositories;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.rest.core.annotation.RestResource;

import com.antulev.togo.models.UserInfo;

@RestResource(exported=false)
public interface UserInfoRepository extends JpaRepository<UserInfo, Long>{
	UserInfo findByUid(String uid);
}
