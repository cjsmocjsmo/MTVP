#!/usr/bin/env python3

import os
import sqlite3

class CreateTables:
    def __init__(self):
        self.conn = sqlite3.connect(os.getenv("MTV_DB_PATH"))
        self.cursor = self.conn.cursor()

        self.conn2 = sqlite3.connect(os.getenv("PROPANE_DB_PATH"))
        self.cursor2 = self.conn2.cursor()

    def create_tables(self):
        self.cursor.execute("""CREATE TABLE IF NOT EXISTS movies (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            Name TEXT NOT NULL,
            Year TEXT NOT NULL,
            PosterAddr TEXT NOT NULL,
            Size TEXT NOT NULL,
            Path TEXT NOT NULL,
            Idx TEXT NOT NULL,
            MovId TEXT NOT NULL UNIQUE,
            Catagory TEXT NOT NULL,
            HttpThumbPath TEXT NOT NULL
        )""")

        self.cursor.execute("""CREATE TABLE IF NOT EXISTS tvshows (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            TvId TEXT NOT NULL UNIQUE,
            Size TEXT NOT NULL,
            Catagory TEXT NOT NULL,
            Name TEXT NOT NULL,
            Season TEXT NOT NULL,
            Episode TEXT NOT NULL,
            Path TEXT NOT NULL,
            Idx TEXT NOT NULL
         )""")
        
        self.cursor.execute("""CREATE TABLE IF NOT EXISTS images (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            ImgId TEXT NOT NULL,
            Path TEXT NOT NULL,
            ImgPath TEXT NOT NULL,
            Size TEXT NOT NULL,
            Name TEXT NOT NULL,
            ThumbPath TEXT NOT NULL,
            Idx INTEGER NOT NULL,
            HttpThumbPath TEXT NOT NULL
         )""")
        
        self.cursor2.execute("""CREATE TABLE IF NOT EXISTS amount (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            Year TEXT NOT NULL,
            Month TEXT NOT NULL,
            Day TEXT NOT NULL,
            Date TEXT NOT NULL,
            Amount TEXT NOT NULL
         )""")
        
        self.cursor2.execute("""CREATE TABLE IF NOT EXISTS gallons (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            Year TEXT NOT NULL,
            Month TEXT NOT NULL,
            Day TEXT NOT NULL,
            Date TEXT NOT NULL,
            Gallons TEXT NOT NULL
         )""")



        

        
