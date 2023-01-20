cmake_minimum_required(VERSION 3.13)

project(model C)
set(CMAKE_C_STANDARD 11)

find_package(onnxruntime REQUIRED)
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS}")

add_executable(model model.c)
target_link_libraries(app "${onnxruntime_LIBRARIES}")
set_property(TARGET model PROPERTY C_STANDARD 11)
