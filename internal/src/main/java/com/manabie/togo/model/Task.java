package com.manabie.togo.model;

import com.fasterxml.jackson.annotation.JsonIgnore;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.OneToOne;
import javax.persistence.Table;
import lombok.Data;

/**
 * Task entity map to "tasks" table
 * @author mupmup
 */
@Entity
@Table(name = "tasks")
@Data
public class Task {

    @Id
    @Column(name = "id", nullable = false, unique = true)
    private String id;

    @Column(name = "content", nullable = false)
    private String content;

    @Column(name = "created_date", nullable = false)
    private String created_date;
    
    @JsonIgnore
    @OneToOne
    @JoinColumn(name = "user_id", nullable = true)
    private User user;
    
    @Column(name = "user_id", insertable = false, updatable = false)
    private String user_id;
    
}
