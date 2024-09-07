#!/usr/bin/env python3

import os
import re

def get_arch():
    if os.uname().machine == "armv7l":
        return "32"
    elif os.uname().machine == "aarch64":
        return "64"
    elif os.uname().machine == "x86_64":
        return "64"
    
def mtv_walk_dirs(directory):
    re_mov = re.compile(r"Movies", re.IGNORECASE)
    re_tv = re.compile(r"TVShows", re.IGNORECASE)
    medialist = []
    for root, dirs, files in os.walk(directory):
        for file in files:
            fname = os.path.join(root, file)
            medialist.append(fname)
    return medialist

def img_walk_dirs(dir):
    medialist = []
    for root, dirs, files in os.walk(dir):
        for file in files:
            fname = os.path.join(root, file)
            ext = os.path.splitext(fname)[1]
            if ext == ".jpg":
                medialist.append(fname)
    return medialist
