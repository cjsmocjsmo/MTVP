#!/usr/bin/env python

import os
import sqlite3

class Media:
    def __init__(self):
        self.conn = sqlite3.connect(os.getenv('MTV_DB_PATH'))
        self.cursor = self.conn.cursor()

    def _fetch_all_as_dict(self):
        columns = [column[0] for column in self.cursor.description]
        return [dict(zip(columns, row)) for row in self.cursor.fetchall()]


    def action(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Action' ORDER BY year DESC")
        return self._fetch_all_as_dict()

    
    def arnold(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Arnold' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def brucelee(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='BruceLee' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def brucewillis(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='BruceWillis' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def buzz(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Buzz' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def cartoons(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Cartoons' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def charliebrown(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='CharlieBrown' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def comedy(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Comedy' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def chucknorris(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='ChuckNorris' ORDER BY year DESC")
        return self._fetch_all_as_dict()

    def documentary(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Documentary' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def drama(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Drama' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def fantasy(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Fantasy' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def ghostbusters(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='GhostBusters' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def godzilla(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Godzilla' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def harrypotter(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='HarryPotter' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def indianajones(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='IndianaJones' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def jamesbond(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='JamesBond' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def johnwayne(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='JohnWayne' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def johnwick(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='JohnWick' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def jurassicpark(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='JurassicPark' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def kevincostner(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='KevinCostner' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def kingsmen(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Kingsmen' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def lego(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Lego' ORDER BY year DESC")
        print(f"lego executed")
        return self._fetch_all_as_dict()
    
    def meninblack(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='MenInBlack' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def minions(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Minions' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def misc(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Misc' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def nicolascage(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='NicolasCage' ORDER BY year DESC")
        return self._fetch_all_as_dict()

    def oldies(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Oldies' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def panda(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Panda' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def pirates(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Pirates' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def riddick(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Riddick' ORDER BY year DESC")
        return self._fetch_all_as_dict()

    def scifi(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='SciFi' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def stalone(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Stalone' ORDER BY year DESC")
        return self._fetch_all_as_dict()

    def startrek(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='StarTrek' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def starwars(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='StarWars' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def superheros(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='SuperHeros' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def therock(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='TheRock' ORDER BY year DESC")
        return self._fetch_all_as_dict()

    def tinkerbell(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='TinkerBell' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def tremors(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Tremors' ORDER BY year DESC")
        return self._fetch_all_as_dict()

    def tomcruize(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='TomCruize' ORDER BY year DESC")
        return self._fetch_all_as_dict()

    def transformers(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Transformers' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def xmen(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='XMen' ORDER BY year DESC")
        return self._fetch_all_as_dict()
    
    def alteredcarbon(self):
        command = "SELECT * FROM tvshows WHERE catagory='AlteredCarbon' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def cowboybebop(self):
        command = "SELECT * FROM tvshows WHERE catagory='CowboyBebop' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def fallout(self):
        command = "SELECT * FROM tvshows WHERE catagory='FallOut' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def forallmankind(self):
        command = "SELECT * FROM tvshows WHERE catagory='ForAllMankind' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def foundation(self):
        command = "SELECT * FROM tvshows WHERE catagory='Foundation' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def fuubar(self):
        command = "SELECT * FROM tvshows WHERE catagory='FuuBar' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def hford1923(self):
        command = "SELECT * FROM tvshows WHERE catagory='HFord1923' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def halo(self):
        command = "SELECT * FROM tvshows WHERE catagory='Halo' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def houseofthedragon_s1(self):
        command = "SELECT * FROM tvshows WHERE catagory='HouseOfTheDragon' AND season='01' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def houseofthedragon_s2(self):
        command = "SELECT * FROM tvshows WHERE catagory='HouseOfTheDragon' AND season='02' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def lostinspace(self):
        command = "SELECT * FROM tvshows WHERE catagory='LostInSpace' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def mastersoftheuniverse(self):
        command = "SELECT * FROM tvshows WHERE catagory='MastersOfTheUniverse' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def monarchlegacyofmonsters(self):
        command = "SELECT * FROM tvshows WHERE catagory='MonarchLegacyOfMonsters' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def nightsky(self):
        command = "SELECT * FROM tvshows WHERE catagory='NightSky' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def orville(self):
        command = "SELECT * FROM tvshows WHERE catagory='Orville' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def prehistoricplanet(self):
        command = "SELECT * FROM tvshows WHERE catagory='PrehistoricPlanet' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def raisedbywolves(self):
        command = "SELECT * FROM tvshows WHERE catagory='RaisedByWolves' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def shogun(self):
        command = "SELECT * FROM tvshows WHERE catagory='Shogun' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def silo(self):
        command = "SELECT * FROM tvshows WHERE catagory='Silo' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def thecontinental(self):
        command = "SELECT * FROM tvshows WHERE catagory='TheContinental' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def thelastofus(self):
        command = "SELECT * FROM tvshows WHERE catagory='TheLastOfUs' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def thelordoftheringstheringsofpower(self):
        command = "SELECT * FROM tvshows WHERE catagory='TheLordOfTheRingsTheRingsOfPower' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    
    
    
    def wheeloftimes1(self):
        command = "SELECT * FROM tvshows WHERE catagory='WheelOfTime' AND season='01' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def wheeloftimes2(self):
        command = "SELECT * FROM tvshows WHERE catagory='WheelOfTime' AND season='02' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    

    
    def discovery(self):
        command = "SELECT * FROM tvshows WHERE catagory='Discovery' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def enterprise(self):
        command = "SELECT * FROM tvshows WHERE catagory='Enterprise' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def lowerdecks(self):
        command = "SELECT * FROM tvshows WHERE catagory='LowerDecks' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def picard(self):
        command = "SELECT * FROM tvshows WHERE catagory='Picard' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def prodigy(self):
        command = "SELECT * FROM tvshows WHERE catagory='Prodigy' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def sttv(self):
        command = "SELECT * FROM tvshows WHERE catagory='STTV' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def strangenewworlds(self):
        command = "SELECT * FROM tvshows WHERE catagory='StrangeNewWorlds' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def tng(self):
        command = "SELECT * FROM tvshows WHERE catagory='TNG' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def voyager(self):
        command = "SELECT * FROM tvshows WHERE catagory='Voyager' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def acolyte(self):
        command = "SELECT * FROM tvshows WHERE catagory='Acolyte' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def ahsoka(self):
        command = "SELECT * FROM tvshows WHERE catagory='Ahsoka' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def andor(self):
        command = "SELECT * FROM tvshows WHERE catagory='Andor' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def bookofbobafett(self):
        command = "SELECT * FROM tvshows WHERE catagory='BookOfBobaFett' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def mandalorian(self):
        command = "SELECT * FROM tvshows WHERE catagory='Mandalorian' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def obiwankenobi(self):
        command = "SELECT * FROM tvshows WHERE catagory='ObiWanKenobi' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def talesoftheempire(self):
        command = "SELECT * FROM tvshows WHERE catagory='TalesOfTheEmpire' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def talesofthejedi(self):
        command = "SELECT * FROM tvshows WHERE catagory='TalesOfTheJedi' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def thebadbatch(self):
        command = "SELECT * FROM tvshows WHERE catagory='TheBadBatch' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def visions(self):
        command = "SELECT * FROM tvshows WHERE catagory='Visions' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def falconwintersoldier(self):
        command = "SELECT * FROM tvshows WHERE catagory='FalconWinterSoldier' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def hawkeye(self):
        command = "SELECT * FROM tvshows WHERE catagory='Hawkeye' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def iamgroots1(self):
        command = "SELECT * FROM tvshows WHERE catagory='IAmGroot' AND season='01' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def iamgroots2(self):
        command = "SELECT * FROM tvshows WHERE catagory='IAmGroot' AND season='02' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def loki(self):
        command = "SELECT * FROM tvshows WHERE catagory='Loki' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def moonknight(self):
        command = "SELECT * FROM tvshows WHERE catagory='MoonKnight' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def secretinvasion(self):
        command = "SELECT * FROM tvshows WHERE catagory='SecretInvasion' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def shehulk(self):
        command = "SELECT * FROM tvshows WHERE catagory='SheHulk' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def wandavision(self): 
        command = "SELECT * FROM tvshows WHERE catagory='WandaVision' ORDER BY Episode ASC;"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

