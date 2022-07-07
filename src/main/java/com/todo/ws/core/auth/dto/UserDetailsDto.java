package com.todo.ws.core.auth.dto;

import java.util.Collection;

import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;

import com.todo.ws.core.auth.model.User;


public class UserDetailsDto implements UserDetails {

	 	private final String username;
	    private final String password;
	    private final Long id;

	    public UserDetailsDto(User user) {
	        this.username = user.getUsername();
	        this.password = user.getPassword();
	        this.id = user.getId();
	    }

	    @Override
	    public Collection<? extends GrantedAuthority> getAuthorities() {
	        return null;
	    }

	    @Override
	    public String getPassword() {
	        return this.password;
	    }

	    @Override
	    public String getUsername() {
	        return this.username;
	    }

	    @Override
	    public boolean isAccountNonExpired() {
	        return true;
	    }

	    @Override
	    public boolean isAccountNonLocked() {
	        return true;
	    }

	    @Override
	    public boolean isCredentialsNonExpired() {
	        return true;
	    }

	    @Override
	    public boolean isEnabled() {
	        return true;
	    }

	    public Long getId() {
	        return id;
	    }
}
