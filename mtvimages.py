#!/usr/bin/env python3

import hashlib
import os
from PIL import Image
import re
import subprocess
import sqlite3
from pprint import pprint
import logging
from dotenv import load_dotenv

# Load environment variables
load_dotenv()

log_file = os.getenv('MTV_SERVER_LOG')
logging.basicConfig(filename=log_file, level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

class ProcessImages:
    def __init__(self, imgs, conn, cursor):
        self.conn = conn
        self.cursor = cursor
        self.imglist = imgs
        self.search = re.compile("\s\(")

    def thumb_dir_check(self):
        img_dir = os.getenv("MTV_THUMBNAIL_PATH")
        if not os.path.exists(img_dir):
            subprocess.run(["mkdir", img_dir])
            logging.info(f"Created directory: {img_dir}")

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
        return md5_hash.hexdigest()

    def get_name(self, img):
        img = os.path.split(img)[1]
        match = re.search(self.search, img)
        if match:
            start = match.start()
            return img[:start]
        else:
            logging.warning(f"No regex match found for image: {img}")
            return img

    def get_thumb_path(self, img):
        new_dir = os.getenv("MTV_THUMBNAIL_PATH")
        fname = os.path.split(img)[1]
        return os.path.join(new_dir, fname)

    def get_http_thumb_path(self, img):
        fname = os.path.split(img)[1]
        server_addr = os.getenv("MTV_SERVER_ADDR")
        server_port = "9999"
        return f"{server_addr}:{server_port}/{fname}"
    
    def process(self):
        self.thumb_dir_check()
        for idx, img in enumerate(self.imglist):
            thumb = self.create_thumbnail(img)
            media_info = {
                "ImgId": self.get_img_id(thumb),
                "Size": os.stat(thumb).st_size,
                "Name": self.get_name(img),
                "ThumbPath": self.get_thumb_path(thumb),
                "Path": thumb,
                "Idx": idx+1,
                "HttpThumbPath": self.get_http_thumb_path(thumb),
            }
            # logging.info(f"Processing image: {media_info}")
            
            try:
                self.cursor.execute('''INSERT INTO images (ImgId, Path, ImgPath, Size, Name, ThumbPath, Idx, HttpThumbPath)
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
                self.conn.commit()
                
            except sqlite3.IntegrityError as e:
                logging.error(f'SQLite IntegrityError in ProcessImages: {e}')
            except sqlite3.OperationalError as e:
                logging.error(f"SQLite OperationalError in ProcessImages: {e}")
        
class ProcessTVShowImages:
    def __init__(self, imgs):
        self.imglist = imgs

    def tv_thumb_dir_check(self):
        img_dir = os.getenv("MTV_TV_THUMBNAIL_PATH")
        if not os.path.exists(img_dir):
            subprocess.run(["mkdir", img_dir])
            logging.info(f"Created TV thumbnail directory: {img_dir}")

    def create_thumbnail(self, img_path):
        thumb_dir = os.getenv("MTV_TV_THUMBNAIL_PATH")
        img = os.path.splitext(img_path)[0]
        fname = os.path.split(img)[1]
        fname = ".".join((fname, "jpg"))
        save_path = os.path.join(thumb_dir, fname)
        logging.info(f"Saving thumbnail to: {save_path}")

        thumb = Image.open(img_path)
        thumb.thumbnail((300, 300))
        thumb.save(save_path)
        return save_path
    
    def process_tv_thumbs(self):
        self.tv_thumb_dir_check()
        logging.info(f"Processing {len(self.imglist)} TV show images")
        for img in self.imglist:
            self.create_thumbnail(img)
