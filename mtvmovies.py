#!/usr/bin/env python3

import hashlib
import os
import re
from pprint import pprint
import sqlite3

class ProcessMovies:
    def __init__(self, movs, conn, cursor):
        self.conn = conn
        self.cursor = cursor
        self.movlist = movs
        self.action = re.compile("Action")
        self.chuck_norris = re.compile("ChuckNorris")
        self.arnold = re.compile("Arnold")
        self.bruce_lee = re.compile("BruceLee")
        self.bruce_willis = re.compile("BruceWillis")
        self.buzz = re.compile("Buzz")
        self.cartoons = re.compile("Cartoons")
        self.charlie_brown = re.compile("CharlieBrown")
        self.chuck_norris = re.compile("ChuckNorris")
        self.comedy = re.compile("Comedy")
        self.documentary = re.compile("Documentary")
        self.drama = re.compile("Drama")
        self.fantasy = re.compile("Fantasy")
        self.ghost_busters = re.compile("GhostBusters")
        self.godzilla = re.compile("Godzilla")
        self.harrison_ford = re.compile("HarrisonFord")
        self.harry_potter = re.compile("HarryPotter")
        self.hellboy = re.compile("HellBoy")
        self.indiana_jones = re.compile("IndianaJones")
        self.james_bond = re.compile("JamesBond")
        self.john_wayne = re.compile("JohnWayne")
        self.john_wick = re.compile("JohnWick")
        self.jurassic_park = re.compile("JurassicPark")
        self.kevin_costner = re.compile("KevinCostner")
        self.kingsman = re.compile("Kingsman")
        self.lego = re.compile("Lego")
        self.meninblack = re.compile("MenInBlack")
        self.minions = re.compile("Minions")
        self.misc = re.compile("Misc")
        self.nicolas_cage = re.compile("NicolasCage")
        self.oldies = re.compile("Oldies")
        self.panda = re.compile("Panda")
        self.pirates = re.compile("Pirates")
        self.riddick = re.compile("Riddick")
        self.sci_fi = re.compile("SciFi")
        self.stalone = re.compile("Stalone")
        self.startrek = re.compile("StarTrek")
        self.starwars = re.compile("StarWars")
        self.super_heros = re.compile("SuperHeros")
        self.the_rock = re.compile("TheRock")
        self.tinker_bell = re.compile("TinkerBell")
        self.tom_cruize = re.compile("TomCruize")
        self.transformers = re.compile("Transformers")
        self.tremors = re.compile("Tremors")
        self.vandam = re.compile("VanDam")
        self.xmen = re.compile("XMen")
        self.crap = re.compile("\s\(")
        

    def get_year(self, mov):
        searchstr1 = re.compile("\(")
        searchstr2 = re.compile("\)")
        start = 0
        end = 0
        match1 = re.search(searchstr1, mov)
        match2 = re.search(searchstr2, mov)
        if match1:
            start = match1.start() + 1
        if match2:
            end = match2.start()
        return mov[start:end]

    def get_poster(self, mov):
        # return os.path.splitext(mov)[0] + ".jpg"
        return ".".join((os.path.splitext(mov)[0], "jpg"))

    def get_mov_id(self, mov):
        encoded_string = mov.encode('utf-8')
        md5_hash = hashlib.md5()
        md5_hash.update(encoded_string)
        return md5_hash.hexdigest()

    def get_catagory(self, mov1):
        mov = os.path.split(mov1)[0]
        catagory = ""
        if re.search(self.action, mov):
            catagory = "Action"
        elif re.search(self.arnold, mov):
            catagory = "Arnold"
        elif re.search(self.bruce_lee, mov):
            catagory = "BruceLee"
        elif re.search(self.bruce_willis, mov):
            catagory = "BruceWillis"
        elif re.search(self.buzz, mov):
            catagory = "Buzz"
        elif re.search(self.cartoons, mov):
            catagory = "Cartoons"
        elif re.search(self.charlie_brown, mov):
            catagory = "CharlieBrown"
        elif re.search(self.chuck_norris, mov):
            catagory = "ChuckNorris"
        elif re.search(self.comedy, mov):
            catagory = "Comedy"
        elif re.search(self.documentary, mov):
            catagory = "Documentary"
        elif re.search(self.drama, mov):
            catagory = "Drama"
        elif re.search(self.fantasy, mov):
            catagory = "Fantasy"
        elif re.search(self.ghost_busters, mov):
            catagory = "GhostBusters"
        elif re.search(self.godzilla, mov):
            catagory = "Godzilla"
        elif re.search(self.harrison_ford, mov):
            catagory = "HarrisonFord"
        elif re.search(self.harry_potter, mov):
            catagory = "HarryPotter"
        elif re.search(self.hellboy, mov):
            catagory = "HellBoy"
        elif re.search(self.indiana_jones, mov):
            catagory = "IndianaJones"
        elif re.search(self.james_bond, mov):
            catagory = "JamesBond"
        elif re.search(self.john_wayne, mov):
            catagory = "JohnWayne"
        elif re.search(self.john_wick, mov):
            catagory = "JohnWick"
        elif re.search(self.jurassic_park, mov):
            catagory = "JurassicPark"
        elif re.search(self.kevin_costner, mov):
            catagory = "KevinCostner"
        elif re.search(self.kingsman, mov):
            catagory = "Kingsman"
        elif re.search(self.lego, mov):
            catagory = "Lego"
        elif re.search(self.meninblack, mov):
            catagory = "MenInBlack"
        elif re.search(self.minions, mov):
            catagory = "Minions"
        elif re.search(self.misc, mov):
            catagory = "Misc"
        elif re.search(self.nicolas_cage, mov):
            catagory = "NicolasCage"
        elif re.search(self.oldies, mov):
            catagory = "Oldies"
        elif re.search(self.panda, mov):
            catagory = "Panda"
        elif re.search(self.pirates, mov):
            catagory = "Pirates"
        elif re.search(self.riddick, mov):
            catagory = "Riddick"
        elif re.search(self.sci_fi, mov):
            catagory = "SciFi"
        elif re.search(self.stalone, mov):
            catagory = "Stalone"
        elif re.search(self.startrek, mov):
            catagory = "StarTrek"
        elif re.search(self.starwars, mov):
            catagory = "StarWars"
        elif re.search(self.super_heros, mov):
            catagory = "SuperHeros"
        elif re.search(self.the_rock, mov):
            catagory = "TheRock"
        elif re.search(self.tinker_bell, mov):
            catagory = "TinkerBell"
        elif re.search(self.tom_cruize, mov):
            catagory = "TomCruize"
        elif re.search(self.transformers, mov):
            catagory = "Transformers"
        elif re.search(self.tremors, mov):
            catagory = "Tremors"
        elif re.search(self.vandam, mov):
            catagory = "VanDam"
        elif re.search(self.xmen, mov):
            catagory = "XMen"           
        return catagory

    def get_http_thumb_path(self, mov):
        fname = os.path.split(mov)[1]
        server_addr = os.getenv("MTV_SERVER_ADDR")
        server_port = "9999"
        return f"{server_addr}:{server_port}/{fname}"
    
    def get_name(self, mov):
        fname = os.path.split(mov)[1]
        match = re.search(self.crap, fname)
        if match:
            start = match.start()
            return fname[:start]

    def process(self):
        for idx, mov in enumerate(self.movlist):
            poster = self.get_poster(mov)
            media_info = {
                "Name": self.get_name(mov),
                "Year": self.get_year(mov),
                "PosterAddr": poster,
                "Size": os.stat(mov).st_size,
                "Path": mov,
                "Idx": idx+1,
                "MovId": self.get_mov_id(mov),
                "Catagory": self.get_catagory(mov),
                "HttpThumbPath": self.get_http_thumb_path(poster),
            }
            # pprint(media_info)
            
            # Insert media_info into the movies table
            try:
                self.cursor.execute('''
                    INSERT INTO movies (Name, Year, PosterAddr, Size, Path, Idx, MovId, Catagory, HttpThumbPath)
                    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
                ''', (
                    media_info["Name"],
                    media_info["Year"],
                    media_info["PosterAddr"],
                    media_info["Size"],
                    media_info["Path"],
                    media_info["Idx"],
                    media_info["MovId"],
                    media_info["Catagory"],
                    media_info["HttpThumbPath"]
                ))

                # Commit the changes and close the connection
                self.conn.commit()
                
            except sqlite3.IntegrityError as e:
                print(f"Error: {e}")
            except sqlite3.OperationalError as e:
                print(f"Error: {e}")
