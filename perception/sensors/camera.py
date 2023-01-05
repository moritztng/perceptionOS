import cv2 as cv

class Camera:
    def __init__(self, url):
        self.url = url
    def save_image(self, filename):
        capture = cv.VideoCapture(self.url)
        if not capture.isOpened():
            print("Cannot open camera")
            exit()
        ret, frame = capture.read()
        if not ret:
            print("Can't receive frame (stream end?). Exiting ...")
        cv.imwrite(filename, frame)
        capture.release()
