package com.backend.campus.model;

import jakarta.persistence.Entity;
import jakarta.persistence.Table;
import lombok.Data;

import java.util.Date;
import java.util.UUID;

@Entity
@Table(name = "USER")
@Data
public class User {

    private UUID userId;

    private String name;

    private String surName;

    private String eMail;

    private String password;

    private Date createdAt;

    private Date updatedAt;

    private Date deletedAt;

    private Role role;
}
