#!/usr/bin/env python3
import argparse
import main
import subprocess
from pprint import pprint
from dotenv import load_dotenv
import utils

def setup():
    parser = argparse.ArgumentParser(description="CLI for Rusic music server.")
    parser.add_argument("-i", "--install", action="store_true", help="Install the program")
    parser.add_argument("-u", "--update", action="store_true", help="Update the program")
    parser.add_argument("-d", "--delete", action="store_true", help="Delete the program")

    args = parser.parse_args()

    arch = utils.get_arch()
    

    if args.install:
        main.Main().main()
        if arch == 32:
            subprocess.run([
                'docker', 
                'run', 
                '--name mtv-thumbs-nginx', 
                '-v', 
                '/usr/share/MTV/thumbnails:/usr/share/nginx/html:ro', 
                '-d',
                "-p 9999:80",
                'arm32v7/nginx:bookworm'
            ])
        elif arch == '64':
            subprocess.run([
                'docker', 
                'run', 
                '--name mtv-thumbs-nginx', 
                '-v', 
                '/usr/share/MTV/thumbnails:/usr/share/nginx/html:ro', 
                '-d',
                "-p 9999:80",
                'nginx:bookworm'
            ])
        print(type(arch))
        

    elif args.update:
        pass
    elif args.delete:   
        pass


    

if __name__ == "__main__":
    load_dotenv()
    setup()