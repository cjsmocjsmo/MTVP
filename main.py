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

            tvimages = utils.tv_img_walk_dirs(os.getenv("MTV_TV_POSTER_PATH"))
            mtvimages.ProcessTVShowImages(tvimages).process_tv_thumbs()

            images = utils.img_walk_dirs(os.getenv("MTV_POSTER_PATH"))
            mtvimages.ProcessImages(images, self.conn, self.cursor).process()

            movs = utils.mtv_walk_dirs(os.getenv("MTV_MOVIES_PATH"))
            mtvmovies.ProcessMovies(movs, self.conn, self.cursor).process()

        except sqlite3.OperationalError as e:
            print(e)
        finally:
            self.conn.close()
            # self.cursor.close()

    def update(self):
        try:
            dirtvshows = utils.mtv_walk_dirs(os.getenv("MTV_TV_PATH"))
            dbtvshows = mtvtvshows.UpdateTVShows(self.conn, self.cursor).get_all_tvshow_paths()
            new_tvshows = [tvshow for tvshow in dirtvshows if tvshow not in dbtvshows]
            mtvtvshows.ProcessTVShows(new_tvshows, self.conn, self.cursor).process()

            dirtvshowimages = utils.tv_img_walk_dirs(os.getenv("MTV_TV_POSTER_PATH"))
            dbtvshowimages = mtvtvshows.UpdateTVShows(self.conn, self.cursor).get_all_tvshow_images()
            new_tvshowimages = [tvshowimage for tvshowimage in dirtvshowimages if tvshowimage not in dbtvshowimages]
            mtvimages.ProcessTVShowImages(new_tvshowimages).process_tv_thumbs()

            dirmovs = utils.mtv_walk_dirs(os.getenv("MTV_MOVIES_PATH"))
            dbmovs = mtvmovies.UpdateMovies(self.conn, self.cursor).get_all_movie_paths()
            new_movs = [mov for mov in dirmovs if mov not in dbmovs]
            mtvmovies.ProcessMovies(new_movs, self.conn, self.cursor).process()

            dirmovimages = utils.img_walk_dirs(os.getenv("MTV_POSTER_PATH"))
            dbmovimages = mtvmovies.UpdateMovies(self.conn, self.cursor).get_all_movie_images()
            new_movimages = [movimage for movimage in dirmovimages if movimage not in dbmovimages]
            mtvimages.ProcessImages(new_movimages, self.conn, self.cursor).process()
            

        except sqlite3.OperationalError as e:
            print(e)
        finally:
            self.conn.close()
            # self.cursor.close()

# if __name__ == "__main__":
#     m = Main()
#     m.main()
