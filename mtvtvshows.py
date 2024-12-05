#!/usr/bin/env python3

import hashlib
import os
from pprint import pprint
import re
import sqlite3
import logging

logging.basicConfig(filename='/usr/share/MTV/media_info.log', level=logging.INFO, format='%(asctime)s - %(message)s')

class ProcessTVShows:
    def __init__(self, tvshows, conn, cursor):
        self.conn = conn
        self.cursor = cursor
        self.tvlist = tvshows
        self.alteredcarbon = re.compile("AlteredCarbon")
        self.forallmankind = re.compile("ForAllManKind")
        self.foundation = re.compile("Foundation")
        self.fuubar = re.compile("FuuBar")
        self.hford1923 = re.compile("HFord1923")
        self.halo = re.compile("Halo")
        self.houseofthedragon = re.compile("HouseOfTheDragon")
        self.lostinspace = re.compile("LostInSpace")
        self.mastersoftheuniverse = re.compile("MastersOfTheUniverse")
        self.monarchlegacyofmonsters = re.compile("MonarchLegacyOfMonsters")
        self.nightsky = re.compile("NightSky")
        self.orville = re.compile("Orville")
        self.prehistoricplanet = re.compile("PrehistoricPlanet")
        self.raisedbywolves = re.compile("RaisedByWolves")
        self.shogun = re.compile("Shogun")
        self.silo = re.compile("Silo")
        self.columbia = re.compile("Columbia")
        self.cowboybebop = re.compile("CowboyBebop")
        self.fallout = re.compile("Fallout")
        self.thecontinental = re.compile("TheContinental")
        self.thelastofus = re.compile("TheLastOfUs")
        self.thelordoftheringstheringsofpower = re.compile("TheLordOfTheRingsTheRingsOfPower")
        self.wheeloftime = re.compile("WheelOfTime")
        self.discovery = re.compile("Discovery")
        self.enterprise = re.compile("Enterprise")
        self.lowerdecks = re.compile("LowerDecks")
        self.picard = re.compile("Picard")
        self.prodigy = re.compile("Prodigy")
        self.sttv = re.compile("STTV")
        self.strangenewworlds = re.compile("StrangeNewWorlds")
        self.tng = re.compile("TNG")
        self.voyager = re.compile("Voyager")
        self.acolyte = re.compile("Acolyte")
        self.andor = re.compile("Andor")
        self.mandalorian = re.compile("Mandalorian")
        self.talesoftheempire = re.compile("TalesOfTheEmpire")
        self.thebadbatch = re.compile("TheBadBatch")
        self.ahsoka = re.compile("Ahsoka")
        self.bookofbobafett = re.compile("BookOfBobaFett")
        self.obiwankenobi = re.compile("ObiWanKenobi")
        self.talesofthejedi = re.compile("TalesOfTheJedi")
        self.visions = re.compile("Visions")
        self.falconwintersoldier = re.compile("FalconWinterSoldier")
        self.iamgroot = re.compile("IAmGroot")
        self.moonknight = re.compile("MoonKnight")
        self.shehulk = re.compile("SheHulk")
        self.hawkeye = re.compile("Hawkeye")
        self.loki = re.compile("Loki")
        self.secretinvasion = re.compile("SecretInvasion")
        self.wandavision = re.compile("WandaVision")
        self.skeletoncrew = re.compile("SkeletonCrew")
        self.episea = re.compile("\sS\d{2}E\d{2}\s")

    def get_tvid(self, tv):
        encoded_string = tv.encode('utf-8')
        md5_hash = hashlib.md5()
        md5_hash.update(encoded_string)
        return md5_hash.hexdigest()

    def get_catagory(self, tv):
        catagory = ""
        if re.search(self.alteredcarbon, tv):
            catagory = "AlteredCarbon"
        elif re.search(self.forallmankind, tv):
            catagory = "ForAllManKind"
        elif re.search(self.foundation, tv):
            catagory = "Foundation"
        elif re.search(self.fuubar, tv):
            catagory = "FuuBar"
        elif re.search(self.hford1923, tv):
            catagory = "HFord1923"
        elif re.search(self.halo, tv):
            catagory = "Halo"
        elif re.search(self.houseofthedragon, tv):
            catagory = "HouseOfTheDragon"
        elif re.search(self.lostinspace, tv):
            catagory = "LostInSpace"
        elif re.search(self.mastersoftheuniverse, tv):
            catagory = "MastersOfTheUniverse"
        elif re.search(self.monarchlegacyofmonsters, tv):
            catagory = "MonarchLegacyOfMonsters"
        elif re.search(self.nightsky, tv):
            catagory = "NightSky"
        elif re.search(self.orville, tv):
            catagory = "Orville"
        elif re.search(self.prehistoricplanet, tv):
            catagory = "PrehistoricPlanet"
        elif re.search(self.raisedbywolves, tv):
            catagory = "RaisedByWolves"
        elif re.search(self.shogun, tv):
            catagory = "Shogun"
        elif re.search(self.silo, tv):
            catagory = "Silo"
        elif re.search(self.columbia, tv):
            catagory = "Columbia"
        elif re.search(self.cowboybebop, tv):
            catagory = "CowboyBebop"
        elif re.search(self.fallout, tv):
            catagory = "Fallout"
        elif re.search(self.thecontinental, tv):
            catagory = "TheContinental"
        elif re.search(self.thelastofus, tv):
            catagory = "TheLastOfUs"
        elif re.search(self.thelordoftheringstheringsofpower, tv):
            catagory = "TheLordOfTheRingsTheRingsOfPower"
        elif re.search(self.wheeloftime, tv):
            catagory = "WheelOfTime"
        elif re.search(self.discovery, tv):
            catagory = "Discovery"
        elif re.search(self.enterprise, tv):
            catagory = "Enterprise"
        elif re.search(self.lowerdecks, tv):
            catagory = "LowerDecks"
        elif re.search(self.picard, tv):
            catagory = "Picard"
        elif re.search(self.prodigy, tv):
            catagory = "Prodigy"
        elif re.search(self.sttv, tv):
            catagory = "STTV"
        elif re.search(self.strangenewworlds, tv):
            catagory = "StrangeNewWorlds"
        elif re.search(self.tng, tv):
            catagory = "TNG"
        elif re.search(self.voyager, tv):
            catagory = "Voyager"
        elif re.search(self.acolyte, tv):
            catagory = "Acolyte"
        elif re.search(self.andor, tv):
            catagory = "Andor"
        elif re.search(self.mandalorian, tv):
            catagory = "Mandalorian"
        elif re.search(self.talesoftheempire, tv):
            catagory = "TalesOfTheEmpire"
        elif re.search(self.thebadbatch, tv):
            catagory = "TheBadBatch"
        elif re.search(self.ahsoka, tv):
            catagory = "Ahsoka"
        elif re.search(self.bookofbobafett, tv):
            catagory = "BookOfBobaFett"
        elif re.search(self.obiwankenobi, tv):
            catagory = "ObiWanKenobi"
        elif re.search(self.talesofthejedi, tv):
            catagory = "TalesOfTheJedi"
        elif re.search(self.visions, tv):
            catagory = "Visions"
        elif re.search(self.obiwankenobi, tv):
            catagory = "ObiWanKenobi"
        elif re.search(self.talesofthejedi, tv):
            catagory = "TalesOfTheJedi"
        elif re.search(self.visions, tv):
            catagory = "Visions"
        elif re.search(self.skeletoncrew, tv):
            catagory = "SkeletonCrew"
        elif re.search(self.falconwintersoldier, tv):
            catagory = "FalconWinterSoldier"
        elif re.search(self.iamgroot, tv):
            catagory = "IAmGroot"
        elif re.search(self.moonknight, tv):
            catagory = "MoonKnight"
        elif re.search(self.shehulk, tv):
            catagory = "SheHulk"
        elif re.search(self.hawkeye, tv):
            catagory = "Hawkeye"
        elif re.search(self.loki, tv):
            catagory = "Loki"
        elif re.search(self.secretinvasion, tv):
            catagory = "SecretInvasion"
        elif re.search(self.wandavision, tv):
            catagory = "WandaVision"
        return catagory

    def get_name(self, tv):
        tv = os.path.split(tv)[1]
        tvu = tv.upper()
        match = re.search(self.episea, tvu)
        if match:
            start = match.start()
            new_start = start + 1
            return tv[:new_start]
        else:
            print("No match")
            print(tv)

    def get_season(self, tv):
        tvu = tv.upper()
        match = re.search(self.episea, tvu)
        if match:
            start = match.start()
            end = match.end()
            SE = tv[start:end]
            season = SE[2:4]
            return season
        
    def get_episode(self, tv):
        tvu = tv.upper()
        match = re.search(self.episea, tvu)
        if match:
            start = match.start()
            end = match.end()
            SE = tv[start:end]
            episode = SE[5:7]
            return episode

    def process(self):
        for idx, tv in enumerate(self.tvlist):
            media_info = {
                "TvId": self.get_tvid(tv),
                "Size": os.stat(tv).st_size,
                "Catagory": self.get_catagory(tv),
                "Name": self.get_name(tv),
                "Season": self.get_season(tv),
                "Episode": self.get_episode(tv),
                "Path": tv,
                "Idx": idx+1,
            }
            
            logging.info(media_info)

            try:
                self.cursor.execute('''INSERT INTO tvshows (TvId, Size, Catagory, Name, Season, Episode, Path, Idx)
                    VALUES (?, ?, ?, ?, ?, ?, ?, ?)
                ''', (
                    media_info["TvId"],
                    media_info["Size"],
                    media_info["Catagory"],
                    media_info["Name"],
                    media_info["Season"],
                    media_info["Episode"],
                    media_info["Path"],
                    media_info["Idx"]
                ))
                self.conn.commit()
                
            except sqlite3.IntegrityError as e:
                print(f"Error: {e}")
                continue
            except sqlite3.OperationalError as e:
                print(f"Error: {e}")
                continue