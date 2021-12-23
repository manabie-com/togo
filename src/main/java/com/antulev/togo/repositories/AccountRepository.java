package com.antulev.togo.repositories;

import java.util.Optional;

import org.springframework.data.ldap.repository.LdapRepository;
import org.springframework.data.rest.core.annotation.RestResource;

import com.antulev.togo.models.Account;

@RestResource(exported=false)
public interface AccountRepository extends LdapRepository<Account>{
	Optional<Account> findByUid(String uid);
}
