package pro.datnt.manabie.task.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;
import pro.datnt.manabie.task.model.TaskDBO;

import java.util.List;


@Repository
public interface TaskRepository extends JpaRepository<TaskDBO, Long>, CrudRepository<TaskDBO, Long> {
    @Query("SELECT COUNT(userId) FROM TaskDBO WHERE userId = ?1 AND createdDate >= CURRENT_DATE")
    Integer countTask(Long userId);

    List<TaskDBO> findAllByUserId(Long userId);
}
