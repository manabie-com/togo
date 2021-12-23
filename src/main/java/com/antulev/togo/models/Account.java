package com.antulev.togo.models;

import org.springframework.ldap.odm.annotations.Attribute;
import org.springframework.ldap.odm.annotations.DnAttribute;
import org.springframework.ldap.odm.annotations.Entry;
import javax.naming.Name;

@Entry(base="ou=Users,dc=antulev,dc=com", objectClasses = { "person", "inetOrgPerson", "organizationalPerson", "top" })
public class Account {
	
	@org.springframework.ldap.odm.annotations.Id
	private Name name;

	@DnAttribute(index= 3, value = "uid") @Attribute(name="uid") String uid;
	
	@Attribute(name="cn") private String firstName;
	
	@Attribute(name="sn") private String lastName;
	
	@Attribute(name="userPassword") String password;

	public String getUid() {
		return uid;
	}

	public void setUid(String uid) {
		this.uid = uid;
	}

	public String getFirstName() {
		return firstName;
	}

	public void setFirstName(String firstName) {
		this.firstName = firstName;
	}

	public String getLastName() {
		return lastName;
	}

	public void setLastName(String lastName) {
		this.lastName = lastName;
	}

	public String getPassword() {
		return password;
	}

	public void setPassword(String password) {
		this.password = password;
	}

	public Name getName() {
		return name;
	}

	public void setName(Name name) {
		this.name = name;
	}
	
}
