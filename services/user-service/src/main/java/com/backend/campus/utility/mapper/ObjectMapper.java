package com.backend.campus.utility.mapper;

import org.modelmapper.ModelMapper;
import org.springframework.stereotype.Service;

@Service
public class ObjectMapper {

    private final ModelMapper modelMapper;

    public ObjectMapper(ModelMapper modelMapper) {
        this.modelMapper = modelMapper;
    }
}