package com.manabie.interview.task.model;

import lombok.*;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;

import javax.persistence.*;
import java.util.Collection;
import java.util.Collections;

@EqualsAndHashCode
@Entity
@Getter
@Setter
@ToString
@AllArgsConstructor
@Table(name = "users")
public class User implements UserDetails {

    @Id
    private String uid;
    @ToString.Exclude private String upassword;
    private int maxTask;

    @Enumerated(EnumType.STRING)
    private UserRole userRole;



    @Override
    public Collection<? extends GrantedAuthority> getAuthorities() {
        SimpleGrantedAuthority authority = new SimpleGrantedAuthority(userRole.name());
        return Collections.singletonList(authority);
    }

    public User(){

    }

    @Override
    public String getUsername() {
        return uid;
    }
    @Override
    public String getPassword() {
        return upassword;
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

    public void setPassword(String encode) {
        this.upassword = encode;
    }
}
