#!/usr/bin/env python3

import hashlib
import os
import re

class ProcessTVShows:
    def __init__(self, tvshows):
        self.tvlist = tvshows
        self.alteredcarbon = re.compile(" AlteredCarbon")
        self.forallmankind = re.compile(" ForAllManKind")
        self.foundation = re.compile(" Foundation")
        self.fuubar = re.compile(" FuuBar")
        self.hford1923 = re.compile(" HFord1923")
        self.halo = re.compile(" Halo")
        self.houseofthedragon = re.compile(" HouseOfTheDragon")
        self.lostinspace = re.compile(" LostInSpace")
        self.mastersoftheuniverse = re.compile(" MastersOfTheUniverse")
        self.monarchlegacyofmonsters = re.compile(" MonarchLegacyOfMonsters")
        self.nightsky = re.compile(" NightSky")
        self.orville = re.compile(" Orville")
        self.prehistoricplanet = re.compile(" PrehistoricPlanet")
        self.raisedbywolves = re.compile(" RaisedByWolves")
        self.shogun = re.compile(" Shogun")
        self.silo = re.compile(" Silo")
        self.columbia = re.compile(" Columbia")
        self.cowboybebop = re.compile(" CowboyBebop")
        self.fallout = re.compile(" Fallout")
        self.thecontinental = re.compile(" TheContinental")
        self.thelastofus = re.compile(" TheLastOfUs")
        self.thelordoftheringstheringsofpower = re.compile(" TheLordOfTheRingsTheRingsOfPower")
        self.wheeloftime = re.compile(" WheelOfTime")
        self.discovery = re.compile(" Discovery")
        self.enterprise = re.compile(" Enterprise")
        self.lowerdecks = re.compile(" LowerDecks")
        self.picard = re.compile(" Picard")
        self.prodigy = re.compile(" Prodigy")
        self.sttv = re.compile(" STTV")
        self.strangenewworlds = re.compile(" StrangeNewWorlds")
        self.tng = re.compile(" TNG")
        self.voyager = re.compile(" Voyager")
        self.acolyte = re.compile(" Acolyte")
        self.andor = re.compile(" Andor")
        self.mandalorian = re.compile(" Mandalorian")
        self.talesoftheempire = re.compile(" TalesOfTheEmpire")
        self.thebadbatch = re.compile(" TheBadBatch")
        self.ahsoka = re.compile(" Ahsoka")
        self.bookofbobafett = re.compile(" BookOfBobaFett")
        self.obiwankenobi = re.compile(" ObiWanKenobi")
        self.talesofthejedi = re.compile(" TalesOfTheJedi")
        self.visions = re.compile(" Visions")
        self.falconwintersoldier = re.compile(" FalconWinterSoldier")
        self.iamgroot = re.compile(" IAmGroot")
        self.moonknight = re.compile(" MoonKnight")
        self.shehulk = re.compile(" SheHulk")
        self.hawkeye = re.compile(" Hawkeye")
        self.loki = re.compile(" Loki")
        self.secretinvasion = re.compile(" SecretInvasion")
        self.wandavision = re.compile(" WandaVision")
        self.episea = re.compile("^S\d{2}E\d{2}$")

    def get_tvid(self, tv):
        encoded_string = tv.encode('utf-8')
        md5_hash = hashlib.md5()
        md5_hash.update(encoded_string)
        hash_hex = md5_hash.hexdigest()
        return hash_hex

    def get_catagory(self, tv):
        catagory = ""
        if self.alteredcarbon.search(tv):
            catagory = "AlteredCarbon"
        elif self.forallmankind.search(tv):
            catagory = "ForAllManKind"
        elif self.foundation.search(tv):
            catagory = "Foundation"
        elif self.fuubar.search(tv):
            catagory = "FuuBar"
        elif self.hford1923.search(tv):
            catagory = "HFord1923"
        elif self.halo.search(tv):
            catagory = "Halo"
        elif self.houseofthedragon.search(tv):
            catagory = "HouseOfTheDragon"
        elif self.lostinspace.search(tv):
            catagory = "LostInSpace"
        elif self.mastersoftheuniverse.search(tv):
            catagory = "MastersOfTheUniverse"
        elif self.monarchlegacyofmonsters.search(tv):
            catagory = "MonarchLegacyOfMonsters"
        elif self.nightsky.search(tv):
            catagory = "NightSky"
        elif self.orville.search(tv):
            catagory = "Orville"
        elif self.prehistoricplanet.search(tv):
            catagory = "PrehistoricPlanet"
        elif self.raisedbywolves.search(tv):
            catagory = "RaisedByWolves"
        elif self.shogun.search(tv):
            catagory = "Shogun"
        elif self.silo.search(tv):
            catagory = "Silo"
        elif self.columbia.search(tv):
            catagory = "Columbia"
        elif self.cowboybebop.search(tv):
            catagory = "CowboyBebop"
        elif self.fallout.search(tv):
            catagory = "Fallout"
        elif self.thecontinental.search(tv):
            catagory = "TheContinental"
        elif self.thelastofus.search(tv):
            catagory = "TheLastOfUs"
        elif self.thelordoftheringstheringsofpower.search(tv):
            catagory = "TheLordOfTheRingsTheRingsOfPower"
        elif self.wheeloftime.search(tv):
            catagory = "WheelOfTime" 
        elif self.discovery.search(tv):
            catagory = "Discovery"
        elif self.enterprise.search(tv):
            catagory = "Enterprise"
        elif self.lowerdecks.search(tv):
            catagory = "LowerDecks"
        elif self.picard.search(tv):
            catagory = "Picard"
        elif self.prodigy.search(tv):
            catagory = "Prodigy"
        elif self.sttv.search(tv):
            catagory = "STTV"
        elif self.strangenewworlds.search(tv):
            catagory = "StrangeNewWorlds"
        elif self.tng.search(tv):
            catagory = "TNG"
        elif self.voyager.search(tv):
            catagory = "Voyager"
        elif self.acolyte.search(tv):
            catagory = "Acolyte"
        elif self.andor.search(tv):
            catagory = "Andor" 
        elif self.mandalorian.search(tv):
            catagory = "Mandalorian"
        elif self.talesoftheempire.search(tv):
            catagory = "TalesOfTheEmpire"
        elif self.thebadbatch.search(tv):
            catagory = "TheBadBatch"
        elif self.ahsoka.search(tv):
            catagory = "Ahsoka"
        elif self.bookofbobafett.search(tv):
            catagory = "BookOfBobaFett"
        elif self.obiwankenobi.search(tv):
            catagory = "ObiWanKenobi"
        elif self.talesofthejedi.search(tv):
            catagory = "TalesOfTheJedi"
        elif self.visions.search(tv):
            catagory = "Visions"
        elif self.falconwintersoldier.search(tv):
            catagory = "FalconWinterSoldier"
        elif self.iamgroot.search(tv):
            catagory = "IAmGroot"  
        elif self.moonknight.search(tv):
            catagory = "MoonKnight"
        elif self.shehulk.search(tv):
            catagory = "SheHulk"
        elif self.hawkeye.search(tv):
            catagory = "Hawkeye"
        elif self.loki.search(tv):
            catagory = "Loki"
        elif self.secretinvasion.search(tv):
            catagory = "SecretInvasion"
        elif self.wandavision.search(tv):
            catagory = "WandaVision"

        return catagory

    def get_name(self, tv):
        searchstr = re.compile("^S\d{2}E\d{2}$")
        match = re.search(searchstr, tv)
        if match:
            start = match.start()
            new_start = start - 1
            return tv[:new_start]

    def get_season(self, tv):
        tvu = tv.upper()
        searchstr = re.compile("^S\d{2}E\d{2}$")
        match = re.search(searchstr, tvu)
        if match:
            start = match.start()
            end = match.end()
            SE = tv[start:end]
            season = SE[1:3]
            print(f"Season: {season}")
            return season
        
    def get_episode(self, tv):
        tvu = tv.upper()
        searchstr = re.compile("^S\d{2}E\d{2}$")
        match = re.search(searchstr, tvu)
        if match:
            start = match.start()
            end = match.end()
            SE = tv[start:end]
            episode = SE[4:6]
            print(f"Episode: {episode}")
            return episode
    
    def get_size(self, tv):
        file_stat = os.stat(tv)
        return file_stat.st_size
    

    def process(self):
        idx = 0
        for tv in self.tvlist:
            idx += 1
            
            media_info = {
                "TvId": self.get_tvid(tv),
                "Size": self.get_size(tv),
                "Catagory": self.get_catagory(tv),
                "Name": self.get_name(tv),
                "Season": self.get_season(tv),
                "Episode": self.get_episode(tv),
                "Path": tv,
                "Idx": idx,
            }
            print(media_info)
