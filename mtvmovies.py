#!/usr/bin/env python3

import hashlib
import os
import re

class ProcessMovies:
    def __init__(self, movs):
        self.self.movlist = movs
        self.action = re.compile("Action")
        self.chuck_norris = re.compile("ChuckNorris")
        self.arnold = re.compile("Arnold")
        self.bruce_lee = re.compile("BruceLee")
        self.bruce_willis = re.compile("BruceWillis")
        self.buzz = re.compile("Buzz")
        self.cartoons = re.compile("Cartoons")
        self.charlie_brown = re.compile("CharlieBrown")
        self.comedy = re.compile("Comedy")
        self.documentary = re.compile("Documentary")
        self.drama = re.compile("Drama")
        self.fantasy = re.compile("Fantasy")
        self.ghost_busters = re.compile("GhostBusters")
        self.godzilla = re.compile("Godzilla")
        self.harry_potter = re.compile("HarryPotter")
        self.indiana_jones = re.compile("IndianaJones")
        self.james_bond = re.compile("JamesBond")
        self.john_wayne = re.compile("JohnWayne")
        self.john_wick = re.compile("JohnWick")
        self.jurassic_park = re.compile("JurassicPark")
        self.kevin_costner = re.compile("KevinCostner")
        self.kingsman = re.compile("Kingsman")
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
        self.xmen = re.compile("XMen")

    def get_year(self, mov):
        pass

    def get_poster(self, mov):
        return os.path.splitext(mov)[0] + ".jpg"

    def get_mov_id(self, mov):
        encoded_string = mov.encode('utf-8')
        md5_hash = hashlib.md5()
        md5_hash.update(encoded_string)
        hash_hex = md5_hash.hexdigest()
        return hash_hex

    def get_catagory(self, mov):
        catagory = ""

        if self.action.search(mov):
            catagory = "Action"
        elif self.arnold.search(mov):
            catagory = "Arnold"
        elif self.bruce_lee.search(mov):
            catagory = "BruceLee"
        elif self.bruce_willis.search(mov):
            catagory = "BruceWillis"
        elif self.buzz.search(mov):
            catagory = "Buzz"
        elif self.cartoons.search(mov):
            catagory = "Cartoons"
        elif self.charlie_brown.search(mov):
            catagory = "CharlieBrown"
        elif self.comedy.search(mov):
            catagory = "Comedy"
        elif self.documentary.search(mov):
            catagory = "Documentary"
        elif self.drama.search(mov):
            catagory = "Drama"
        elif self.fantasy.search(mov):
            catagory = "Fantasy"
        elif self.ghost_busters.search(mov):
            catagory = "GhostBusters"
        elif self.godzilla.search(mov):
            catagory = "Godzilla"
        elif self.harry_potter.search(mov):
            catagory = "HarryPotter"
        elif self.indiana_jones.search(mov):
            catagory = "IndianaJones"
        elif self.james_bond.search(mov):
            catagory = "JamesBond"
        elif self.john_wayne.search(mov):
            catagory = "JohnWayne"
        elif self.john_wick.search(mov):
            catagory = "JohnWick"
        elif self.jurassic_park.search(mov):
            catagory = "JurassicPark"
        elif self.kevin_costner.search(mov):
            catagory = "KevinCostner"
        elif self.kingsman.search(mov):
            catagory = "Kingsman"
        elif self.meninblack.search(mov):
            catagory = "MenInBlack"
        elif self.minions.search(mov):
            catagory = "Minions"
        elif self.misc.search(mov):
            catagory = "Misc"
        elif self.nicolas_cage.search(mov):
            catagory = "NicolasCage"
        elif self.oldies.search(mov):
            catagory = "Oldies"
        elif self.panda.search(mov):
            catagory = "Panda"
        elif self.pirates.search(mov):
            catagory = "Pirates"
        elif self.riddick.search(mov):
            catagory = "Riddick"
        elif self.sci_fi.search(mov):
            catagory = "SciFi"
        elif self.stalone.search(mov):
            catagory = "Stalone"
        elif self.startrek.search(mov):
            catagory = "StarTrek"
        elif self.starwars.search(mov):
            catagory = "StarWars"
        elif self.super_heros.search(mov):
            catagory = "SuperHeros"
        elif self.the_rock.search(mov):
            catagory = "TheRock"
        elif self.tinker_bell.search(mov):
            catagory = "TinkerBell"
        elif self.tom_cruize.search(mov):
            catagory = "TomCruize"
        elif self.transformers.search(mov):
            catagory = "Transformers"
        elif self.tremors.search(mov):
            catagory = "Tremors"
        elif self.xmen.search(mov):
            catagory = "XMen"
                                 
        return catagory

    def get_http_thumb_path(self, mov):
        fname = os.path.split(mov)[1]
        server_addr = os.getenv("MTV_SERVER_ADDR")
        server_port = os.getenv("MTV_SERVER_PORT")
        return f"http://{server_addr}:{server_port}/thumbnails/{fname}"

    def process(self):
        idx = 0
        for mov in self.movlist:
            idx += 1
            media_info = {
                "Year": ,
                "PosterAddr": self.get_poster(mov),
                "Size": os.getsize(mov),
                "Path": mov,
                "Idx": idx,
                "MovId": self.get_mov_id(mov),
                "Catagory": self.get_catagory(mov),
                "HttpThumbPath": self.get_http_thumb_path(mov),
            }
            print(mov)

