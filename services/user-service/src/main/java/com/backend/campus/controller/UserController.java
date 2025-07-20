package com.backend.campus.controller;

import com.backend.campus.dto.UserUpdateDTO;
import com.backend.campus.dto.UserViewDTO;
import com.backend.campus.service.UserService;
import com.backend.campus.utility.result.DataResult;
import com.backend.campus.utility.result.Result;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.security.Principal;
import java.util.List;
import java.util.UUID;

@CrossOrigin
@RestController
@RequestMapping("/api/users")
public class UserController {

    private final UserService userService;

    public UserController(UserService userService) {
        this.userService = userService;
    }

    // Herkes erişebilir → Temel kullanıcı bilgileri
    @GetMapping("/public/{id}")
    public ResponseEntity<DataResult<UserViewDTO>> getPublicUserById(@PathVariable UUID id) {
        return ResponseEntity.ok(userService.getPublicUserById(id));
    }

    // Giriş yapan kişi kendi bilgilerini görür
    @GetMapping("/me")
    public ResponseEntity<DataResult<UserViewDTO>> getCurrentUser(Principal principal) {
        return ResponseEntity.ok(userService.getCurrentUser(principal.getName()));
    }

    // Giriş yapan kişi kendi bilgilerini günceller
    @PutMapping("/me")
    public ResponseEntity<DataResult<UserViewDTO>> updateCurrentUser(@RequestBody UserUpdateDTO dto, Principal principal) {
        return ResponseEntity.ok(userService.updateCurrentUser(principal.getName(), dto));
    }

    // Admin: tüm kullanıcıları getir
    @GetMapping
    public ResponseEntity<DataResult<List<UserViewDTO>>> getAllUsers() {
        return ResponseEntity.ok(userService.getAllUsers());
    }

    // Admin: kullanıcıyı detaylı getir
    @GetMapping("/{id}")
    public ResponseEntity<DataResult<UserViewDTO>> getUserById(@PathVariable UUID id) {
        return ResponseEntity.ok(userService.getUserById(id));
    }

    // Admin: kullanıcıyı güncelle
    @PutMapping("/{id}")
    public ResponseEntity<DataResult<UserViewDTO>> updateUserById(@PathVariable UUID id, @RequestBody UserUpdateDTO dto) {
        return ResponseEntity.ok(userService.updateUserById(id, dto));
    }

    // Admin: kullanıcıyı sil
    @DeleteMapping("/{id}")
    public ResponseEntity<Result> deleteUserById(@PathVariable UUID id) {
        return ResponseEntity.ok(userService.deleteUser(id));
    }
}