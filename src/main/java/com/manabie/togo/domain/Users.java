package com.manabie.togo.domain;

import lombok.Getter;

import javax.persistence.*;

@Entity
@Getter
public class Users {

    @Id
    private String userId;

    @Column(nullable = false)
    private String name;

    @Column(nullable = false)
    private int dailyLimit;
}
