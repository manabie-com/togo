package com.tasks.todo.item;

import org.springframework.data.map.repository.config.EnableMapRepositories;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.Optional;

@Service
@EnableMapRepositories
public class ItemService {
    private final CrudRepository<Item, String> repository;

    public ItemService(CrudRepository<Item, String> repository) {
        this.repository = repository;
    }

    public List<Item> findAll() {
        List<Item> list = new ArrayList<>();
        Iterable<Item> items = repository.findAll();
        items.forEach(list::add);
        return list;
    }

    public Optional<Item> find(String userId, Date date) {
        return repository.findById(userId + date.getTime());
    }

    public Item create(Item item) {
        // To ensure the item ID remains unique,
        // use the current timestamp.
        Item copy = new Item(
                item.getUserId() + new Date().getTime(),
                item.getTaskName(),
                item.getTaskDescription()
        );
        return repository.save(copy);
    }

    public Optional<Item> update( String userId, Item newItem) {
        // Only update an item if it can be found first.
        return repository.findById(userId)
                .map(oldItem -> {
                    Item updated = oldItem.updateWith(newItem);
                    return repository.save(updated);
                });
    }

    public void delete(String userId) {
        repository.deleteById(userId);
    }
}