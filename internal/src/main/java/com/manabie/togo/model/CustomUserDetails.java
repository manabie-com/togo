package com.manabie.togo.model;

import java.util.Collection;
import java.util.Collections;

import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;

import lombok.AllArgsConstructor;
import lombok.Data;

/**
 * Custom user details
 * @author mupmup
 */
@Data
@AllArgsConstructor
public class CustomUserDetails implements UserDetails {

    User user;

    /**
     * Get user authorities
     *
     * @return
     */
    @Override
    public Collection<? extends GrantedAuthority> getAuthorities() {
        // for simple. All users have ROLE_USER
        return Collections.singleton(new SimpleGrantedAuthority("ROLE_USER"));
    }

    /**
     * Get password
     *
     * @return
     */
    @Override
    public String getPassword() {
        return user.getPassword();
    }

    /**
     * Get user name
     *
     * @return
     */
    @Override
    public String getUsername() {
        return user.getUsername();
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
}
