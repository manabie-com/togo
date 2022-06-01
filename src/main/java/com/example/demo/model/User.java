package com.example.demo.model;

import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;

import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.OneToMany;
import javax.persistence.Table;
import javax.persistence.Transient;
import java.util.ArrayList;
import java.util.Collection;
import java.util.List;

@Entity
@Data
@Table(name="users")
@NoArgsConstructor
public class User implements UserDetails {

    // Sprint Security Default
    @Id
    private String username;
    private String password;
    private boolean enabled = true;
    @OneToMany(fetch = FetchType.EAGER)
    private List<MyAuthorities> authorities;
    private boolean accountNonExpired = true;
    private boolean accountNonLocked = true;
    private boolean credentialsNonExpired = true;

    //
    //private int limit;
    public User(String username) {
        this.username = username;
    }

    public void addAuthority(MyAuthorities auth) {
        if (authorities == null) {
            authorities = new ArrayList<>();
        }
        authorities.add(auth);
    }

    public User(String username, String password) {
        this.username = username;
        this.password = password;
    }
}
