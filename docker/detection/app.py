import os
from flask import Flask
from perception.sensors.camera import Camera
from datetime import datetime

app = Flask(__name__)
camera = Camera(os.environ['URL'])

@app.route('/')
def save_image():
    filename = os.path.join(os.environ['IMAGE_DIRECTORY'], datetime.now().strftime('%Y%m%d-%H%M%S.%f') + '.jpg')
    camera.save_image(filename)    
    return filename
