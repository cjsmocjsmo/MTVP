#!/usr/bin/env python

import os
import sqlite3

class Media:
    def __init__(self):
        self.conn = sqlite3.connect(os.getenv('MTV_DB_PATH'))
        self.cursor = self.conn.cursor()

    def action(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Action'")
        return self.cursor.fetchall()
    
    def arnold(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Arnold'")
        return self.cursor.fetchall()
    
    def brucelee(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='BruceLee'")
        return self.cursor.fetchall()
    
    def brucewillis(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='BruceWillis'")
        return self.cursor.fetchall()
    
    def buzz(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Buzz'")
        return self.cursor.fetchall()
    
    def cartoons(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Cartoons'")
        return self.cursor.fetchall()
    
    def charliebrown(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='CharlieBrown'")
        return self.cursor.fetchall()
    
    def comedy(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Comedy'")
        return self.cursor.fetchall()
    
    def chucknorris(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='ChuckNorris'")
        return self.cursor.fetchall()