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

        log_file_path = os.getenv("MTV_SERVER_LOG")
        if not os.path.exists(log_file_path):
            os.makedirs(os.path.dirname(log_file_path), exist_ok=True)
            with open(log_file_path, 'w') as log_file:
                log_file.write("")
        

    def main(self):
        try:
            mtvtables.CreateTables().create_tables()

            tvshows = utils.mtv_walk_dirs(os.getenv("MTV_TV_PATH"))
            mtvtvshows.ProcessTVShows(tvshows, self.conn, self.cursor).process()

            images = utils.img_walk_dirs(os.getenv("MTV_POSTER_PATH"))
            mtvimages.ProcessImages(images, self.conn, self.cursor).process()

            movs = utils.mtv_walk_dirs(os.getenv("MTV_MOVIES_PATH"))
            mtvmovies.ProcessMovies(movs, self.conn, self.cursor).process()

            tvimages = utils.tv_img_walk_dirs(os.getenv("MTV_TV_POSTER_PATH"))
            mtvimages.ProcessTVShowImages(tvimages)

        except sqlite3.OperationalError as e:
            print(e)
        finally:
            self.conn.close()

if __name__ == "__main__":
    m = Main()
    m.main()
