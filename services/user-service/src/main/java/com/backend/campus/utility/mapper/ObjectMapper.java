package com.backend.campus.utility.mapper;

import org.modelmapper.ModelMapper;
import org.springframework.stereotype.Service;

@Service
public class ObjectMapper {

    private final ModelMapper modelMapper;

    public ObjectMapper(ModelMapper modelMapper) {
        this.modelMapper = modelMapper;
    }

    public <T> T map(Object source, Class<T> destinationType) {
        return modelMapper.map(source, destinationType);
    }

    public void map(Object source, Object destination) {
        modelMapper.map(source, destination);
    }
}