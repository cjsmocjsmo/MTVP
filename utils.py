#!/usr/bin/env python3

import os
import subprocess
import sqlite3
from datetime import datetime

def get_arch():
    arch =  os.uname().machine
    if arch == "armv7l":
        return "32"
    elif arch == "arm64" or arch == "x86_64":
        return "64"
    
def mtv_walk_dirs(directory):
    medialist = []
    for root, dirs, files in os.walk(directory):
        for file in files:
            fname = os.path.join(root, file)
            ext = os.path.splitext(fname)[1]
            if ext == ".mp4" or ext == ".mkv":
                medialist.append(fname)
    return medialist

def movie_count():
    conn = sqlite3.connect(os.getenv("MTV_DB_PATH"))
    cursor = conn.cursor()
    cursor.execute("SELECT COUNT(*) FROM movies")
    count = cursor.fetchone()[0]
    return count

def tvshow_count():
    conn = sqlite3.connect(os.getenv("MTV_DB_PATH"))  
    cursor = conn.cursor()
    cursor.execute("SELECT COUNT(*) FROM tvshows")
    count = cursor.fetchone()[0]
    return count

def movies_size_on_disk():
    conn = sqlite3.connect(os.getenv("MTV_DB_PATH"))  
    cursor = conn.cursor()

    cursor.execute("SELECT Size FROM movies")
    sizes = cursor.fetchall()
    
    size_list = [int(size[0]) for size in sizes]
    total_movie_size = sum(size_list)
    
    conn.close()
    
    total_movie_size_gb = total_movie_size / (1024 ** 3)  # Convert bytes to gigabytes
    total_movie_size_gb = round(total_movie_size_gb, 1)
    return total_movie_size_gb

def tvshows_size_on_disk():
    conn = sqlite3.connect(os.getenv("MTV_DB_PATH"))  
    cursor = conn.cursor()

    cursor.execute("SELECT Size FROM tvshows")
    sizes = cursor.fetchall()
    
    size_list = [int(size[0]) for size in sizes]
    total_tvshow_size = sum(size_list)
    
    conn.close()
    
    total_tvshow_size_gb = total_tvshow_size / (1024 ** 3)  # Convert bytes to gigabytes
    total_tvshow_size_gb = round(total_tvshow_size_gb, 1)
    return total_tvshow_size_gb

def img_walk_dirs(dir):
    jpglist = []
    for root, dirs, files in os.walk(dir):
        for file in files:
            fname = os.path.join(root, file)
            ext = os.path.splitext(fname)[1]
            if ext == ".jpg":
                jpglist.append(fname)
    return jpglist

def tv_img_walk_dirs(dir):
    imglist = []
    for root, dirs, files in os.walk(dir):
        for file in files:
            fname = os.path.join(root, file)
            ext = os.path.splitext(fname)[1]
            if ext == ".jpg":
                imglist.append(fname)
    return imglist    

def sqlite3_check():
    sqlite3 = subprocess.run(["apt-cache", "policy", "sqlite3"])
    if sqlite3.returncode == 0:
        return True
    else:
        return False
    
def vlc_check():
    vlc = subprocess.run(["apt-cache", "policy", "vlc"])
    if vlc.returncode == 0:
        return True
    else:
        return False
    
def python3_vlc_check():
    pvlc = subprocess.run(["apt-cache", "policy", "python3-vlc"])
    if pvlc.returncode == 0:
        return True
    else:
        return False
    
def python3_pil_check():
    pil = subprocess.run(["apt-cache", "policy", "python3-pil"])
    if pil.returncode == 0:
        return True
    else:
        return False
    
def python3_dotenv_check():
    dot = subprocess.run(["apt-cache", "policy", "python3-dotenv"])
    if dot.returncode == 0:
        return True
    else:
        return False

def python3_websockets_check():
    ws = subprocess.run(["apt-cache", "policy", "python3-websockets"])
    if ws.returncode == 0:
        return True
    else:
        return False