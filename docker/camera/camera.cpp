#include <iostream>
#include <string>
#include <opencv2/core.hpp>
#include <opencv2/videoio.hpp>
#include <opencv2/imgcodecs.hpp>
#include "camera.h"
using namespace std;
using namespace cv;

void image()
{
    VideoCapture captRefrnc("http://192.168.0.99:8080/shot.jpg");
    if (!captRefrnc.isOpened())
    {
        cout  << "Could not open reference" << endl;
    }
    Mat frameReference;
    captRefrnc >> frameReference;
    if (frameReference.empty())
    {
        cout << " < < <  Game over!  > > > ";
    }
    imwrite("test.jpg", frameReference);
}
