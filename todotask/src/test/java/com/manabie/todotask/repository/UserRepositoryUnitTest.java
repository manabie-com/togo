package com.manabie.todotask.repository;

import com.manabie.todotask.entity.UserDailyLimit;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.annotation.DirtiesContext;

import java.util.Optional;

@SpringBootTest
@DirtiesContext(classMode = DirtiesContext.ClassMode.BEFORE_EACH_TEST_METHOD)
class UserRepositoryUnitTest {
    @Autowired
    UserRepository userRepository;

    @Test
    public void addAndFind(){
        UserDailyLimit userDailyLimit = new UserDailyLimit();
        userDailyLimit.setUserId(1);
        userDailyLimit.setDailyTaskLimit(5);
        UserDailyLimit added = userRepository.save(userDailyLimit);

        Optional<UserDailyLimit> indbOpt = userRepository.findById(userDailyLimit.getUserId());
        assert (indbOpt.isPresent());
        assert (indbOpt.get().getUserId().equals(userDailyLimit.getUserId()));
        assert (indbOpt.get().getDailyTaskLimit().equals(userDailyLimit.getDailyTaskLimit()));
    }
}