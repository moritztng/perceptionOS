cmake_minimum_required(VERSION 3.13)

project(app C)
set(CMAKE_C_STANDARD 11)

find_package(onnxruntime REQUIRED)
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS}")

add_executable(app app.c)
target_link_libraries(app "${onnxruntime_LIBRARIES}")
set_property(TARGET app PROPERTY C_STANDARD 11)
