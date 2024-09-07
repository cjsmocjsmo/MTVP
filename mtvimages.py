#!/usr/bin/env python3

import hashlib
import os
from PIL import Image
import re
import subprocess
import sqlite3

class ProcessImages:
    def __init__(self, imgs):
        self.imglist = imgs
        self.search = re.compile("\s\(")

    def thumb_dir_check(self):
        img_dir = os.getenv("MTV_THUMBNAIL_PATH")
        if not os.path.exists(img_dir):
            subprocess.run(["mkdir", img_dir])
            print(f"Created directory")

    def create_thumbnail(self, img):
        thumb_dir = os.getenv("MTV_THUMBNAIL_PATH")
        fname = os.path.split(img)[1]
        save_path = os.path.join(thumb_dir, fname)

        thumb = Image.open(img)
        thumb.thumbnail((300, 300))
        thumb.save(save_path)
        return save_path


    def get_img_id(self, imgstr):
        encoded_string = imgstr.encode('utf-8')
        md5_hash = hashlib.md5()
        md5_hash.update(encoded_string)
        hash_hex = md5_hash.hexdigest()
        return hash_hex

    def get_name(self, img):
        img = os.path.split(img)[1]
        match = re.search(self.search, img)
        if match:
            start = match.start()
            return img[:start]
        else:
            print("No match")

    def get_thumb_path(self, img):
        new_dir = os.getenv("MTV_THUMBNAIL_PATH")
        fname = os.path.split(img)[1]
        return os.path.join(new_dir, fname)

    def get_http_thumb_path(self, img):
        fname = os.path.split(img)[1]
        server_addr = os.getenv("MTV_SERVER_ADDR")
        server_port = os.getenv("MTV_SERVER_PORT")
        return f"http://{server_addr}:{server_port}/thumbnails/{fname}"
    
    def get_size(self, img):
        file_stat = os.stat(img)
        return file_stat.st_size

    def process(self):
        idx = 0
        self.thumb_dir_check()
        for img in self.imglist:
            thumb = self.create_thumbnail(img)
            idx += 1
            media_info = {
                "ImgId": self.get_img_id(thumb),
                "Size": self.get_size(thumb),
                "Name": self.get_name(img),
                "ThumbPath": self.get_thumb_path(thumb),
                "Path": thumb,
                "Idx": idx,
                "HttpThumbPath": self.get_http_thumb_path(thumb),
            }
            
            db_path = os.getenv("MTV_DB_PATH")
            conn = sqlite3.connect(db_path, timeout=30)
            conn.execute("PRAGMA journal_mode=WAL")
            c = conn.cursor()
            try:
                c.execute('''INSERT INTO images (ImgId, Path, ImgPath, Size, Name, ThumbPath, Idx, HttpThumbPath)
                            VALUES (?, ?, ?, ?, ?, ?, ?, ?)
                ''', (
                    media_info["ImgId"], 
                    media_info["Path"], 
                    media_info["Path"], 
                    media_info["Size"], 
                    media_info["Name"], 
                    media_info["ThumbPath"], 
                    media_info["Idx"], 
                    media_info["HttpThumbPath"]
                ))
                conn.commit()
                
            except sqlite3.IntegrityError:
                print(f'this is bad{img}')

            conn.close()

            




















# ImgId TEXT NOT NULL UNIQUE,
#             Path TEXT NOT NULL,
#             ImgPath TEXT NOT NULL,
#             Size TEXT NOT NULL,
#             Name TEXT NOT NULL,
#             ThumbPath TEXT NOT NULL,
#             Idx INTEGER NOT NULL,
#             HttpThumbPath TEXT NOT NULL