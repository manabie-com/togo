package com.todo.service.account;

import com.todo.entity.AppAccount;
import com.todo.repository.AppAccountRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class AccountServiceImpl implements AccountService {
    @Autowired
    private AppAccountRepository appAccountRepository;



    @Override
    public AppAccount findById(Long id) {
        return appAccountRepository.findById(id).orElse(null);
    }

    @Override
    public void changePassword(Long id, String password) {
        AppAccount appAccount = appAccountRepository.findById(id).orElse(null);
        if (appAccount != null) {
            appAccount.setPassword(password);
            appAccountRepository.save(appAccount);
        }
    }

    @Override
    public void changePass(AppAccount  appAccount, Long id) {
        AppAccount appAccountUpdate = this.findByIdAppAccount(id);
        appAccountUpdate.setPassword(appAccount.getPassword());
        this.appAccountRepository.save(appAccountUpdate);
    }

    @Override
    public void saveAppAccount(AppAccount appAccount) {
        this.appAccountRepository.save(appAccount);
    }

    @Override
    public AppAccount findByIdAppAccount(Long id) {
        return this.appAccountRepository.findById(id).orElse(null);
    }

    @Override
    public AppAccount findByUsername(String username) {
        return appAccountRepository.findByUsername(username);
    }
}
