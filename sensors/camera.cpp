#include <opencv2/videoio.hpp>
#include <opencv2/imgcodecs.hpp>
#include "camera.h"

VideoCapture camera(char* url) { 
    return new cv::VideoCapture(url);
}

void set_resolution(VideoCapture camera, int width, int height) {
    ((cv::VideoCapture*)camera)->set(cv::CAP_PROP_FRAME_WIDTH, width);
    ((cv::VideoCapture*)camera)->set(cv::CAP_PROP_FRAME_HEIGHT, height);
}

void save_image(VideoCapture camera, const char* filename)
{
    cv::Mat image;
    *((cv::VideoCapture*)camera) >> image;
    cv::imwrite(filename, image);
}
