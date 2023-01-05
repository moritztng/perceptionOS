import os, sqlite3
from flask import Flask
from perception.sensors.camera import Camera
from datetime import datetime

app = Flask(__name__)
camera = Camera(os.environ['URL'])

@app.route('/')
def save_image():
    images_db = sqlite3.connect(os.environ['DB_FILENAME'])
    db_cursor = images_db.cursor()
    filename = os.path.join(os.environ['IMAGE_DIRECTORY'], datetime.now().strftime('%Y%m%d-%H%M%S.%f') + '.jpg')
    camera.save_image(filename)
    db_cursor.execute(f"INSERT INTO images (filename) VALUES ('{filename}')")
    images_db.commit()
    return filename
