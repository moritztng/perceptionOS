import cv2 as cv

class FaceDetection:
    def __init__(self):
        self.classifier = cv.CascadeClassifier('haarcascade_frontalface_default.xml')
    def __call__(self, filename):
        img = cv.imread(filename)
        img = cv.cvtColor(img, cv.COLOR_BGR2GRAY)
        faces = self.classifier.detectMultiScale(img)
        detected = len(faces) > 0
        return detected
