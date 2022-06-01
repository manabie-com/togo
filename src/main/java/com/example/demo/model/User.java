package com.example.demo.model;

import com.fasterxml.jackson.annotation.JsonIgnore;
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

    @JsonIgnore
    private String password;
    @JsonIgnore
    private boolean enabled = true;
    @OneToMany(fetch = FetchType.EAGER)
    @JsonIgnore
    private List<MyAuthorities> authorities;
    @JsonIgnore
    private boolean accountNonExpired = true;
    @JsonIgnore
    private boolean accountNonLocked = true;
    @JsonIgnore
    private boolean credentialsNonExpired = true;

    //
    //private int limit;
    public User(String username) {
        this.username = username;
    }

    public User(String username, String password) {
        this.username = username;
        this.password = password;
    }

    public void addAuthority(MyAuthorities auth) {
        if (authorities == null) {
            authorities = new ArrayList<>();
        }
        authorities.add(auth);
    }


}
