package com.backend.campus.repository;

import com.backend.campus.model.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Repository
public interface UserRepository extends JpaRepository<User, UUID> {

    Optional<User> findByIdAndDeletedAtIsNull(UUID id);

    Optional<User> findByEMailAndDeletedAtIsNull(String eMail);

    List<User> findAllByDeletedAtIsNull();
}