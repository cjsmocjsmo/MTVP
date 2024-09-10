#!/usr/bin/env python3
import argparse
import main
import os
import subprocess
from pprint import pprint
from dotenv import load_dotenv
import utils

CWD = os.getcwd()

def setup():
    parser = argparse.ArgumentParser(description="CLI for Rusic music server.")
    parser.add_argument("-i", "--install", action="store_true", help="Install the program")
    parser.add_argument("-u", "--update", action="store_true", help="Update the program")
    parser.add_argument("-d", "--delete", action="store_true", help="Delete the program")

    args = parser.parse_args()
    
    if args.install:
        print(f"utils.sqlite3_check(): {utils.sqlite3_check()}")
        print(f"utils.vlc_check(): {utils.vlc_check()}")
        print(f"utils.python3_vlc_check(): {utils.python3_vlc_check()}")
        print(f"utils.python3_pil_check(): {utils.python3_pil_check()}")
        print(f"utils.python3_dotenv_check(): {utils.python3_dotenv_check()}")
        print(f"utils.python3_websockets_check(): {utils.python3_websockets_check()}")

        if not utils.sqlite3_check():
            subprocess.run(["sudo", "apt-get", "-y", "install", "sqlite3"])
        if not utils.vlc_check():
            subprocess.run(["sudo", "apt-get", "-y", "install", "vlc"])
        if not utils.python3_vlc_check():
            subprocess.run(["sudo", "apt-get", "-y", "install", "python3-vlc"])
        if not utils.python3_pil_check():
            subprocess.run(["sudo", "apt-get", "-y", "install", "python3-pil"])
        if not utils.python3_dotenv_check():
            subprocess.run(["sudo", "apt-get", "-y", "install", "python3-dotenv"])
        if not utils.python3_websockets_check():
            subprocess.run(["sudo", "apt-get", "-y", "install", "python3-websockets"])
        
        main.Main().main()
        
        import mtvserverasync
        mtvserverasync.servermain()
        

    elif args.update:
        pass
    elif args.delete:   
        pass


    

if __name__ == "__main__":
    load_dotenv()
    setup()