package com.backend.campus.dto;

import com.backend.campus.model.Role;
import lombok.Data;

import java.util.Date;
import java.util.UUID;

@Data
public class UserViewDTO {
    private UUID userId;
    private String name;
    private String surName;
    private String eMail;
    private Role role;
    private Date createdAt;
}
