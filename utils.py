#!/usr/bin/env python3

import os
import subprocess

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
    sqlite3 = subprocess.run(["apt-cache", "policy", "sqlite3"])
    if sqlite3 == 0:
        return True
    else:
        return False
    

def vlc_check():
    vlc = subprocess.run(["apt-cache", "policy", "vlc"])
    if vlc == 0:
        return True
    else:
        return False
    
def python3_vlc_check():
    pvlc = subprocess.run(["apt-cache", "policy", "python3-vlc"])
    if pvlc == 0:
        return True
    else:
        return False
    
def python3_pil_check():
    pil = subprocess.run(["apt-cache", "policy", "python3-pil"])
    if pil == 0:
        return True
    else:
        return False
    
def python3_dotenv_check():
    dot = subprocess.run(["apt-cache", "policy", "python3-dotenv"])
    if dot == 0:
        return True
    else:
        return False

def python3_websockets_check():
    ws = subprocess.run(["apt-cache", "policy", "python3-websockets"])
    if ws == 0:
        return True
    else:
        return False