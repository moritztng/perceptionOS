#pragma once

#ifdef __cplusplus
extern "C" {
#endif
typedef void* VideoCapture;
VideoCapture camera(char* url);
void save_image(VideoCapture camera, const char* filename);
#ifdef __cplusplus
}
#endif
