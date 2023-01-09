#include "camera.h"

VideoCapture camera(char* url) {
    return new cv::VideoCapture();
}

void save_image(VideoCapture camera, const char* filename)
{
    cv::Mat* image = new cv::Mat();
    ((cv::VideoCapture*)camera)->read(*image);
    cv::imwrite(filename, *image);
}
