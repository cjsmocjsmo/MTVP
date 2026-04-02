#!/usr/bin/env python3

import mtvmovies
import mtvtvshows
import mtvimages
import mtvtables
import mtvvideos
import os
from pprint import pprint
import sqlite3
import utils
import logging
from dotenv import load_dotenv

CWD = os.getcwd()

class Main:
    def __init__(self):
        load_dotenv()
        
        # Set up logging
        log_file = os.getenv('MTV_SERVER_LOG')
        log_dir = os.path.dirname(log_file)
        if not os.path.exists(log_dir):
            os.makedirs(log_dir, exist_ok=True)
        if not os.path.exists(log_file):
            open(log_file, 'a').close()
            
        logging.basicConfig(filename=log_file, level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')
        
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

            tvimages = utils.tv_img_walk_dirs(os.getenv("MTV_TV_POSTER_PATH"))
            mtvimages.ProcessTVShowImages(tvimages).process_tv_thumbs()

            images = utils.img_walk_dirs(os.getenv("MTV_POSTER_PATH"))
            mtvimages.ProcessImages(images, self.conn, self.cursor).process()

            movs = utils.mtv_walk_dirs(os.getenv("MTV_MOVIES_PATH"))
            mtvmovies.ProcessMovies(movs, self.conn, self.cursor).process()

            videos = utils.mtv_walk_dirs(os.getenv("MTV_VIDEOS_PATH"))
            print(videos)
            mtvvideos.ProcessVideos(videos, self.conn, self.cursor).process()

        except sqlite3.OperationalError as e:
            logging.error(f"SQLite OperationalError in main(): {e}")
        finally:
            self.conn.close()
            # self.cursor.close()

    def update(self):
        try:
            Movresults = mtvmovies.UpdateMovies(self.conn, self.cursor).check_for_mov_updates()
            if Movresults:
                mtvmovies.UpdateMovies(self.conn, self.cursor).updateMovs()
            
            TVresults = mtvtvshows.UpdateTVShows(self.conn, self.cursor).check_for_tv_updates()
            if TVresults:
                mtvtvshows.UpdateTVShows(self.conn, self.cursor).updateTV()
            
        except sqlite3.OperationalError as e:
            logging.error(f"SQLite OperationalError in update(): {e}")
        finally:
            self.conn.close()
            # self.cursor.close()

# if __name__ == "__main__":
#     m = Main()
#     m.main()
