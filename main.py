#!/usr/bin/env python3

import mtvmovies
import mtvtvshows
import mtvimages
import mtvtables
import os
from pprint import pprint
import sqlite3
import utils
from dotenv import load_dotenv

CWD = os.getcwd()

class Main:
    def __init__(self):
        load_dotenv()
        self.conn = sqlite3.connect(os.getenv("MTV_DB_PATH"))
        self.cursor = self.conn.cursor()
        

    def main(self):
    
        try:
            mtvtables.CreateTables().create_tables()

            tvshows = utils.mtv_walk_dirs(os.getenv("MTV_TV_PATH"))
            mtvtvshows.ProcessTVShows(tvshows, self.conn, self.cursor).process()

            images = utils.img_walk_dirs(os.getenv("MTV_POSTER_PATH"))
            mtvimages.ProcessImages(images, self.conn, self.cursor).process()

            movs = utils.mtv_walk_dirs(os.getenv("MTV_MOVIES_PATH"))
            mtvmovies.ProcessMovies(movs, self.conn, self.cursor).process()

        except sqlite3.OperationalError as e:
            print(e)
        finally:
            self.conn.close()

if __name__ == "__main__":
    m = Main()
    m.main()
