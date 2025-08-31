package com.backend.campus.service;

import com.backend.campus.dto.UserUpdateDTO;
import com.backend.campus.dto.UserViewDTO;
import com.backend.campus.model.User;
import com.backend.campus.repository.UserRepository;
import com.backend.campus.utility.mapper.ObjectMapper;
import com.backend.campus.utility.result.*;
import org.springframework.stereotype.Service;

import java.util.Date;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Service
public class UserService {

    private final UserRepository userRepository;
    private final ObjectMapper objectMapper;

    public UserService(UserRepository userRepository, ObjectMapper objectMapper) {
        this.userRepository = userRepository;
        this.objectMapper = objectMapper;
    }


    public DataResult<UserViewDTO> getPublicUserById(UUID id) {
        Optional<User> optionalUser = userRepository.findByIdAndDeletedAtIsNull(id);
        if (optionalUser.isPresent()) {
            User user = optionalUser.get();
            UserViewDTO dto = objectMapper.map(user, UserViewDTO.class);
            return new SuccessDataResult<UserViewDTO>(dto, "Kullanıcı başarıyla getirildi");
        } else {
            return new ErrorDataResult<>("Kullanıcı bulunamadı");
        }
    }

    public DataResult<UserViewDTO> getCurrentUser(String email) {
        Optional<User> optionalUser = userRepository.findByEMailAndDeletedAtIsNull(email);

        if (optionalUser.isPresent()) {
            User user = optionalUser.get();
            UserViewDTO dto = objectMapper.map(user, UserViewDTO.class);
            return new SuccessDataResult<>(dto, "Giriş yapan kullanıcı getirildi");
        } else {
            return new ErrorDataResult<>("Geçerli kullanıcı bulunamadı");
        }
    }

    public DataResult<UserViewDTO> updateCurrentUser(String email, UserUpdateDTO dto) {
        Optional<User> optionalUser = userRepository.findByEMailAndDeletedAtIsNull(email);

        if (optionalUser.isPresent()) {
            User user = optionalUser.get();

            objectMapper.map(dto, user);
            user.setUpdatedAt(new Date());

            User updated = userRepository.save(user);
            UserViewDTO view = objectMapper.map(updated, UserViewDTO.class);

            return new SuccessDataResult<>(view, "Kullanıcı bilgileri güncellendi");
        } else {
            return new ErrorDataResult<>("Kullanıcı bulunamadı");
        }
    }

    public DataResult<List<UserViewDTO>> getAllUsers() {
        List<User> userList = userRepository.findAllByDeletedAtIsNull();

        if (userList.isEmpty()) {
            return new ErrorDataResult<>("Hiç kullanıcı bulunamadı");
        }

        List<UserViewDTO> dtoList = userList.stream()
                .map(user -> objectMapper.map(user, UserViewDTO.class))
                .toList();

        return new SuccessDataResult<>(dtoList, "Tüm kullanıcılar getirildi");
    }

    public DataResult<UserViewDTO> getUserById(UUID id) {
        Optional<User> optionalUser = userRepository.findByIdAndDeletedAtIsNull(id);

        if (optionalUser.isPresent()) {
            User user = optionalUser.get();
            UserViewDTO dto = objectMapper.map(user, UserViewDTO.class);
            return new SuccessDataResult<>(dto, "Kullanıcı getirildi");
        } else {
            return new ErrorDataResult<>("Kullanıcı bulunamadı");
        }
    }

    public DataResult<UserViewDTO> updateUserById(UUID id, UserUpdateDTO dto) {
        Optional<User> optionalUser = userRepository.findByIdAndDeletedAtIsNull(id);

        if (optionalUser.isPresent()) {
            User user = optionalUser.get();

            objectMapper.map(dto, user);
            user.setUpdatedAt(new Date());

            User updated = userRepository.save(user);
            UserViewDTO dtoView = objectMapper.map(updated, UserViewDTO.class);

            return new SuccessDataResult<>(dtoView, "Kullanıcı güncellendi");
        } else {
            return new ErrorDataResult<>("Kullanıcı bulunamadı");
        }
    }

    public Result deleteUser(UUID id) {
        Optional<User> optionalUser = userRepository.findByIdAndDeletedAtIsNull(id);

        if (optionalUser.isPresent()) {
            User user = optionalUser.get();
            user.setDeletedAt(new Date());
            userRepository.save(user);
            return new SuccessResult("Kullanıcı başarıyla silindi");
        } else {
            return new ErrorResult("Kullanıcı bulunamadı");
        }
    }
}