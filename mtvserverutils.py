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
    
    def alteredcarbon(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def cowboybebop(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def fallout(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def forallmankind(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def foundation(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def fubar(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def hfor1923(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def halo(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def houseofthedragon(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def lostinspace(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def mastersoftheuniverse(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def monarchlegacyofmonsters(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def nightsky(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def orville(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def prehistoricplanet(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def raisedbywolves(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def shogun(self):
        command = f"""SELECT * FROM tvshows WHERE catagory='Shogun';"""
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def silo(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def thecontinental(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def thelastofus(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def thelordoftheringstheringsofpower(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def wheeloftime(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def discovery(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def enterprise(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def lowerdecks(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def picard(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def prodigy(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def sttv(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def strangenewworlds(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def tng(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def voyager(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def acolyte(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def ahsoka(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def andor(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def bookofbobafett(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def mandalorian(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def obiwankenobi(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def talesoftheempire(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def talesofthejedi(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def thebadbatch(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def visions(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def falconandthewintersoldier(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def hawkeye(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def iamgroot(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def loki(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def moonknight(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()
    
    def secretinvasion(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def shehulk(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

    def wandavision(self, mediaid): 
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self._fetch_all_as_dict()

