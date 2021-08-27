package com.manabie.togo.model;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.Table;

import lombok.Data;

/**
 * User entity map to "users" table
 * @author mupmup
 */
@Entity
@Table(name = "users")
@Data
public class User {

    @Id
    @Column(name = "id",  nullable = false, unique = true)
    private String username;
    
    @Column(name="password", nullable = false)
    private String password;
    
    @Column(name = "max_todo", nullable = false)
    private Long max_todo;
    

}
