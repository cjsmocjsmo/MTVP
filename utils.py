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
            if ext == ".mp4":
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

def propane_gallons_total():
    conn = sqlite3.connect(os.getenv("PROPANE_DB_PATH"))  
    cursor = conn.cursor()

    cursor.execute("SELECT Gallons FROM gallons")
    gallons = cursor.fetchall()
    
    gallons_list = [float(gallon[0]) for gallon in gallons]
    total_gallons = sum(gallons_list)
    
    conn.close()
    
    return total_gallons

def propane_amount_total():
    conn = sqlite3.connect(os.getenv("PROPANE_DB_PATH"))  
    cursor = conn.cursor()

    cursor.execute("SELECT Amount FROM amount")
    amounts = cursor.fetchall()
    
    amounts_list = [float(amount[0]) for amount in amounts]
    total_amount = sum(amounts_list)
    
    conn.close()
    
    return total_amount

def insert_amount(amount):
    today_date = datetime.now().strftime("%m/%d/%Y")
    year = datetime.now().strftime("%Y")
    month = datetime.now().strftime("%m")
    day = datetime.now().strftime("%d")
    conn = sqlite3.connect(os.getenv("PROPANE_DB_PATH"))
    cursor = conn.cursor()
    cursor.execute("INSERT INTO amount (Year, Month, Day, Date, Amount) VALUES (?, ?, ?, ?, ?)", (year, month, day, today_date, amount,))
    conn.commit()
    conn.close()

def insert_gallons(gallons):
    today_date = datetime.now().strftime("%m/%d/%Y")
    year = datetime.now().strftime("%Y")
    month = datetime.now().strftime("%m")
    day = datetime.now().strftime("%d")
    conn = sqlite3.connect(os.getenv("PROPANE_DB_PATH"))
    cursor = conn.cursor()
    cursor.execute("INSERT INTO gallons (Year, Month, Day, Date, Gallons) VALUES (?, ?, ?, ?, ?)", (year, month, day, today_date, gallons,))
    conn.commit()
    conn.close()

def convert_month(month_int):
    if month_int == 1:
        return "January"
    elif month_int == 2:
        return "February"
    elif month_int == 3:
        return "March"
    elif month_int == 4:
        return "April"
    elif month_int == 5:
        return "May"
    elif month_int == 6:
        return "June"
    elif month_int == 7:
        return "July"
    elif month_int == 8:
        return "August"
    elif month_int == 9:
        return "September"
    elif month_int == 10:
        return "October"
    elif month_int == 11:
        return "November"
    elif month_int == 12:
        return "December"
    else:
        return "Invalid month"

def monthly_amount_total(month):
    conn = sqlite3.connect(os.getenv("PROPANE_DB_PATH"))  
    cursor = conn.cursor()

    cursor.execute("SELECT Amount FROM amount WHERE Month = ?", (month))
    amounts = cursor.fetchall()
    
    amounts_list = [float(amount[0]) for amount in amounts]
    monthly_total_amount = sum(amounts_list)
 
    month_txt = convert_month(int(month))
    
    conn.close()
    
    return total_amount

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