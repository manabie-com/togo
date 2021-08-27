package com.manabie.togo.services;

import com.manabie.togo.model.CustomUserDetails;
import com.manabie.togo.model.Task;
import com.manabie.togo.model.User;
import com.manabie.togo.repository.TaskRepository;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.List;
import java.util.UUID;
import javax.persistence.EntityManager;
import javax.persistence.PersistenceContext;
import javax.persistence.Query;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;

/**
 * Service layer for Task
 * @author mupmup
 */
@Service
public class TaskService {
    
    /**
     * Data layer of task
     */
    @Autowired
    private TaskRepository taskRepository;
    
    /**
     * For query
     */
    @PersistenceContext
    private EntityManager em;
    
    private static final SimpleDateFormat formatter = new SimpleDateFormat("yyyy-MM-dd");
    
    /**
     * List all tasks in db
     * @return 
     */
    public List<Task> listTasks(){
        return taskRepository.findAll();
    }
    
    /**
     * Add task to db
     * @param task 
     */
    public void addTask(Task task){
        taskRepository.save(task);
    }
    
    /**
     * List all tasks that created at createdDate
     * @param createdDate
     * @return List of Tasks
     */
    public List<Task> listTaskInDay(String createdDate){
        String hql = "FROM Task WHERE created_date = :createdDate";
        Query query = em.createQuery(hql, Task.class);
        query.setParameter("createdDate", createdDate);
        return query.getResultList();
    }
    
    /**
     * List tasks that created by user, at createdDate
     * @param userId
     * @param createdDate
     * @return List of Tasks
     */
    public List<Task> listTaskPerDayUser(String userId, String createdDate){
        String hql = "FROM Task WHERE created_date = :createdDate and user_id = :userId";
        Query query = em.createQuery(hql, Task.class);
        query.setParameter("createdDate", createdDate);
        query.setParameter("userId", userId);
        return query.getResultList();
    }
    
    /**
     * Create a new task
     * @param content
     * @return 
     */
    public Task createNewTask(String content){
        try{
            Task task = new Task();

            Date today = new Date();

            User user = null; 
            String userId;
            long max_todo = 0;
            long now_todo = 0;

            String date = formatter.format(today);

            //From authentication --> get user --> max_todo --> list task in day of this user
            Authentication authen = SecurityContextHolder.getContext().getAuthentication();       
            if (authen != null){
                CustomUserDetails details = (CustomUserDetails)authen.getPrincipal();
                if (details != null){
                    user = details.getUser();
                    max_todo = user.getMax_todo();
                    userId = user.getUsername();
                    List<Task> listTask = listTaskPerDayUser(userId, date);
                    if (listTask != null)
                        now_todo = listTask.size();
                }
            }

            //compare now and max
            if (now_todo < max_todo){
                String id = UUID.randomUUID().toString();
                task.setId(id);
                task.setContent(content);        
                task.setCreated_date(date);
                task.setUser(user);
                task.setUser_id(user.getUsername());

                addTask(task);
                return task;
            }else{
                return null;
            }    
        }catch(Exception ex){
            return null;
        }
    
    }
}
