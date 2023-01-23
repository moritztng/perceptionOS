#pragma once

#ifdef __cplusplus
extern "C" {
#endif
typedef void* VideoCapture;
VideoCapture camera(char* url);
void set_resolution(VideoCapture camera, int width, int height);
void save_image(VideoCapture camera, const char* filename);
#ifdef __cplusplus
}
#endif
