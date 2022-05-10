package com.todo.service.account;

import com.todo.entity.AppAccount;

public interface AccountService {
    AppAccount findById(Long id);
    void changePassword(Long id, String password);
    public void changePass(AppAccount appAccount ,Long id);
    public void saveAppAccount(AppAccount appAccount);
    public AppAccount findByIdAppAccount(Long id);

    AppAccount findByUsername(String username);

}
