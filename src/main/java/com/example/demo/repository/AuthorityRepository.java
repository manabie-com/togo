package com.example.demo.repository;

import com.example.demo.model.MyAuthorities;
import org.springframework.data.jpa.repository.JpaRepository;

public interface AuthorityRepository extends JpaRepository<MyAuthorities, Long> {
}
