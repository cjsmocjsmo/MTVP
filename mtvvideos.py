#!/usr/bin/env python3

import hashlib
import os
import re
from pprint import pprint
import sqlite3
import logging
from dotenv import load_dotenv

# Load environment variables
load_dotenv()

log_file = os.getenv('MTV_SERVER_LOG')
logging.basicConfig(filename=log_file, level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

class ProcessVideos:
    def __init__(self, vids, conn, cursor):
        self.conn = conn
        self.cursor = cursor
        self.vidlist = vids
        self.vid = re.compile(r'VIDEO')
        self.avi = re.compile(r'AVI')
        self.mp4 = re.compile(r'MP4')
        self.mpg = re.compile(r'MPG')
        self.dcvid = re.compile(r'dcamVids')
        print(f"Initialized ProcessVideos with {len(vids)} videos.")

    def vid_id(self, file_path):
        sha256_hash = hashlib.sha256()
        try:
            with open(file_path, "rb") as f:
                for byte_block in iter(lambda: f.read(4096), b""):
                    sha256_hash.update(byte_block)
            return sha256_hash.hexdigest()
        except FileNotFoundError:
            return None
    
    def vid_name(self, path):
        name_path = os.path.splitext(os.path.basename(path))[0]
        return name_path
    
    def vid_size(self, path):
        return str(os.path.getsize(path))

    def process(self):
        print(f"Processing {len(self.vidlist)} videos...")
        for idx, v in enumerate(self.vidlist):
            print(v)
            if self.vid.search(v) or self.avi.search(v) or self.mp4.search(v) or self.mpg.search(v) or self.dcvid.search(v):
                try:
                    vid_id = self.vid_id(v)
                    name = self.vid_name(v)
                    size = self.vid_size(v)
                    idx += 1

                    self.cursor.execute("SELECT * FROM videos WHERE VidId = ?", (vid_id,))
                    existing_vid = self.cursor.fetchone()

                    if existing_vid:
                        logging.info(f"Video already exists in database: {name}")
                        continue

                    self.cursor.execute("INSERT INTO videos (VidId, VidPath, Size, Name, Idx) VALUES (?, ?, ?, ?, ?)", (vid_id, v, size, name, idx))
                    self.conn.commit()
                    print(f"Inserted video into database: {name}")
                    logging.info(f"Inserted video into database: {name}")
                except sqlite3.Error as e:
                    logging.error(f"SQLite error while processing video {name}: {e}")