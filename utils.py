#!/usr/bin/env python3

import os
import subprocess
    
def mtv_walk_dirs(directory):
    medialist = []
    for root, dirs, files in os.walk(directory):
        for file in files:
            fname = os.path.join(root, file)
            ext = os.path.splitext(fname)[1]
            if ext == ".mp4":
                medialist.append(fname)
    return medialist

def img_walk_dirs(dir):
    jpglist = []
    for root, dirs, files in os.walk(dir):
        for file in files:
            fname = os.path.join(root, file)
            ext = os.path.splitext(fname)[1]
            if ext == ".jpg":
                jpglist.append(fname)
    return jpglist

def sqlite3_check():
    subprocess.run(["dpkg", "-l", "|", "grep", "sqlite3"])
    if subprocess.run(["$?"], shell=True) == 0:
        return True
    else:
        return False

def vlc_check():
    subprocess.run(["dpkg", "-l", "|", "grep", "vlc"])
    if subprocess.run(["$?"], shell=True) == 0:
        return True
    else:
        return False
    
def python3_vlc_check():
    subprocess.run(["dpkg", "-l", "|", "grep", "python3-vlc"])
    if subprocess.run(["$?"], shell=True) == 0:
        return True
    else:
        return False
    
def python3_pil_check():
    subprocess.run(["dpkg", "-l", "|", "grep", "python3-pil"])
    if subprocess.run(["$?"], shell=True) == 0:
        return True
    else:
        return False
    
def python3_dotenv_check():
    subprocess.run(["dpkg", "-l", "|", "grep", "python3-dotenv"])
    if subprocess.run(["$?"], shell=True) == 0:
        return True
    else:
        return False

def python3_websockets_check():
    subprocess.run(["dpkg", "-l", "|", "grep", "python3-websockets"])
    if subprocess.run(["$?"], shell=True) == 0:
        return True
    else:
        return False