package pro.datnt.manabie.task.model;

import lombok.Getter;
import lombok.Setter;
import org.hibernate.annotations.CreationTimestamp;

import javax.persistence.*;
import java.time.LocalDateTime;
import java.util.Calendar;

@Entity
@Table(name = "tasks")
@Getter @Setter
public class TaskDBO {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String content;
    @Column
    private Long userId;
    @CreationTimestamp
    private LocalDateTime createdDate;

}
