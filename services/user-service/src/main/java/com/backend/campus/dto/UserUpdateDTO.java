package com.backend.campus.dto;

import com.backend.campus.model.Role;
import lombok.Data;

@Data
public class UserUpdateDTO {
    private String name;
    private String surName;
    private Role role; // Admin tarafından güncellenirken kullanılır
}