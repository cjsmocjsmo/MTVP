#!/usr/bin/env python3

import hashlib
import os
import re

class ProcessImages:
    def __init__(self, imgs):
        self.imglist = imgs

    def get_img_id(self, imgstr):
        encoded_string = imgstr.encode('utf-8')
        md5_hash = hashlib.md5()
        md5_hash.update(encoded_string)
        hash_hex = md5_hash.hexdigest()
        return hash_hex

    def get_name(self, img):
        searchstr = re.compile("\(")
        match = re.search(searchstr, img)
        start, end = match.span()
        new_start = start - 1
        return img[:new_start]

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

    def process(self, imglist):
        idx = 0
        for img in self.imglist:
            idx += 1
            media_info = {
                "ImgId": self.get_img_id(img),
                "Size": self.get_size(img),
                "Name": self.get_name(img),
                "ThumbPath": self.get_thumb_path(img),
                "Path": img,
                "Idx": idx,
                "HttpThumbPath": self.get_http_thumb_path(img),
            }
            print(media_info)





















# ImgId TEXT NOT NULL UNIQUE,
#             Path TEXT NOT NULL,
#             ImgPath TEXT NOT NULL,
#             Size TEXT NOT NULL,
#             Name TEXT NOT NULL,
#             ThumbPath TEXT NOT NULL,
#             Idx INTEGER NOT NULL,
#             HttpThumbPath TEXT NOT NULL