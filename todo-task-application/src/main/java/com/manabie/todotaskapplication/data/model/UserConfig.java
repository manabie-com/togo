package com.manabie.todotaskapplication.data.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.hibernate.annotations.GenericGenerator;
import org.hibernate.annotations.Type;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.Table;
import java.io.Serializable;
import java.util.UUID;

/**
 * @author @quoctrung.phan
 * @created 04/06/2022
 * @project todo-task-application
 */
@Entity
@Table(name = "user_configs")
@Data
@AllArgsConstructor
@NoArgsConstructor
public class UserConfig implements Serializable {
    private static final long serialVersionUID = 3620997384988593541L;
    @Id
    @GeneratedValue(generator = "UUID")
    @GenericGenerator(
            name = "UUID",
            strategy = "org.hibernate.id.UUIDGenerator"
    )
    @Type(type="uuid-char")
    private UUID id;

    private String userId;

    private String configType;

    private String value;
}
