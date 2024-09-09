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

    arch = utils.get_arch()
    
    if args.install:
        main.Main().main()
        crap = "/".join((CWD, "Dockerfile"))
        if os.path.exists(crap):
            subprocess.run(["rm", crap])
        if arch == '32':
            subprocess.run([
                'docker', 
                'run', 
                '-v', 
                '/usr/share/MTV/thumbnails:/usr/share/nginx/html:ro', 
                '-d',
                "-p",
                "9999:80",
                'arm32v7/nginx:bookworm'
            ])
            old = "/".join((CWD, "mtv_docker_32_file"))
            new = "/".join((CWD, "/Dockerfile"))
            subprocess.run([
                "cp",
                "-pvr", 
                old, 
                new
            ])
            subprocess.run([
                "docker", 
                "build", 
                "-t", 
                "mtv32:0.0.1", "."
            ])
            subprocess.run([
                'docker',
                'run',
                '-v',
                '/usr/share/MTV/mtv.db:/usr/share/MTV/mtv.db:ro',
                '-d',
                '-p',
                '8080:8080',
                "mtv32:0.0.1"
            ])

        elif arch == '64':
            subprocess.run([
                'docker', 
                'run', 
                '-v', 
                '/usr/share/MTV/thumbnails:/usr/share/nginx/html:ro', 
                '-d',
                "-p",
                "9999:80",
                'nginx:bookworm'
            ])
            old = "/".join((CWD, "mtv_docker_64_file"))
            new = "/".join((CWD, "/Dockerfile"))
            subprocess.run(["cp", "-pvr", old, new])
            subprocess.run([
                'docker', 
                'build',
                '-t',
                'mtv64:0.0.1',
                '.'
            ])
            subprocess.run([
                'docker',
                'run',
                '-v',
                '/usr/share/MTV/mtv.db:/usr/share/MTV/mtv.db:ro',
                '-d',
                '-p',
                '8080:8080',
                "."
            ])

        print(type(arch))
        

    elif args.update:
        pass
    elif args.delete:   
        pass


    

if __name__ == "__main__":
    load_dotenv()
    setup()