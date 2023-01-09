#include <opencv2/videoio.hpp>
#include <opencv2/imgcodecs.hpp>
#include "camera.h"

VideoCapture camera(char* url) {
    return new cv::VideoCapture();
}

void save_image(VideoCapture camera, const char* filename)
{
    cv::Mat image;
    *((cv::VideoCapture*)camera) >> image;
    cv::imwrite(filename, image);
}
