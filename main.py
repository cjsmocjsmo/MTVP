#!/usr/bin/env python3

import argparse
import mtvmovies
import os
import subprocess
import utils
from dotenv import load_dotenv

CWD = os.getcwd()



def main():
    parser = argparse.ArgumentParser(description="CLI for Rusic music server.")
    parser.add_argument("-i", "--install", action="store_true", help="Install the program")
    parser.add_argument("-u", "--update", action="store_true", help="Update the program")
    parser.add_argument("-d", "--delete", action="store_true", help="Delete the program")

    args = parser.parse_args()

    arch = utils.get_arch()

    if args.install:

        movs = utils.mtv_walk_dirs(os.getenv("MTV_MOVIES_PATH"))
        mtvmovies.ProcessMovies(movs).process()

        tvshows = utils.mtv_walk_dirs(os.getenv("MTV_TV_PATH"))
        mtvtvshows = mtvtvshows.ProcessTVShows(tvshows).process()
        
    elif args.update:
        pass
    elif args.delete:
        pass
        
           

if __name__ == "__main__":
    load_dotenv()
    main()