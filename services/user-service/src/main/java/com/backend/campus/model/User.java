package com.backend.campus.model;

import jakarta.persistence.*;
import lombok.Data;
import org.hibernate.annotations.CreationTimestamp;
import org.hibernate.annotations.UpdateTimestamp;

import java.util.Date;
import java.util.UUID;

@Entity
@Table(name = "USER")
@Data
public class User {

    @Id
    @GeneratedValue
    @Column(name = "USER_ID", columnDefinition = "uuid")
    private UUID userId;

    @Column(name = "NAME", nullable = false)
    private String name;

    @Column(name = "SURNAME", nullable = false)
    private String surName;

    @Column(name = "EMAIL", nullable = false, unique = true)
    private String eMail;

    @Column(name = "PASSWORD", nullable = false)
    private String password;

    @CreationTimestamp
    @Temporal(TemporalType.TIMESTAMP)
    @Column(name = "CREATED_AT", updatable = false)
    private Date createdAt;

    @UpdateTimestamp
    @Temporal(TemporalType.TIMESTAMP)
    @Column(name = "UPDATED_AT")
    private Date updatedAt;

    @Temporal(TemporalType.TIMESTAMP)
    @Column(name = "DELETED_AT")
    private Date deletedAt;

    @Enumerated(EnumType.STRING)
    @Column(name = "ROLE", nullable = false)
    private Role role;
}
