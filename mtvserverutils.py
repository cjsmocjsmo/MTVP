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
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Action'")
        return self._fetch_all_as_dict()

    
    def arnold(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Arnold'")
        return self.cursor.fetchall()
    
    def brucelee(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='BruceLee'")
        return self.cursor.fetchall()
    
    def brucewillis(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='BruceWillis'")
        return self.cursor.fetchall()
    
    def buzz(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Buzz'")
        return self.cursor.fetchall()
    
    def cartoons(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Cartoons'")
        return self.cursor.fetchall()
    
    def charliebrown(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='CharlieBrown'")
        return self.cursor.fetchall()
    
    def comedy(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Comedy'")
        return self.cursor.fetchall()
    
    def chucknorris(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='ChuckNorris'")
        return self.cursor.fetchall()

    def documentary(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Documentary'")
        return self.cursor.fetchall()
    
    def drama(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Drama'")
        return self.cursor.fetchall()
    
    def fantasy(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Fantasy'")
        return self.cursor.fetchall()
    
    def ghostbusters(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='GhostBusters'")
        return self.cursor.fetchall()
    
    def godzilla(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Godzilla'")
        return self.cursor.fetchall()
    
    def harrypotter(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='HarryPotter'")
        return self.cursor.fetchall()
    
    def indianajones(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='IndianaJones'")
        return self.cursor.fetchall()
    
    def jamesbond(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='JamesBond'")
        return self.cursor.fetchall()
    
    def johnwayne(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='JohnWayne'")
        return self.cursor.fetchall()
    
    def johnwick(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='JohnWick'")
        return self.cursor.fetchall()
    
    def jurassicpark(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='JurassicPark'")
        return self.cursor.fetchall()
    
    def kevincostner(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='KevinCostner'")
        return self.cursor.fetchall()
    
    def kingsmen(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Kingsmen'")
        return self.cursor.fetchall()
    
    def meninblack(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='MenInBlack'")
        return self.cursor.fetchall()
    
    def minions(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Minions'")
        return self.cursor.fetchall()
    
    def misc(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Misc'")
        return self.cursor.fetchall()
    
    def nicolascage(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='NicolasCage'")
        return self.cursor.fetchall()

    def oldies(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Oldies'")
        return self.cursor.fetchall()
    
    def panda(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Panda'")
        return self.cursor.fetchall()
    
    def pirates(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Pirates'")
        return self.cursor.fetchall()
    
    def riddick(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Riddick'")
        return self.cursor.fetchall()

    def scifi(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='SciFi'")
        return self.cursor.fetchall()
    
    def stalone(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Stalone'")
        return self.cursor.fetchall()

    def startrek(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='StarTrek'")
        return self.cursor.fetchall()
    
    def starwars(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='StarWars'")
        return self.cursor.fetchall()
    
    def superheros(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='SuperHeros'")
        return self.cursor.fetchall()
    
    def therock(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='TheRock'")
        return self.cursor.fetchall()

    def tinkerbell(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='TinkerBell'")
        return self.cursor.fetchall()
    
    def tremors(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Tremors'")
        return self.cursor.fetchall()

    def tomcruize(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='TomCruize'")
        return self.cursor.fetchall()

    def transformers(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='Transformers'")
        return self.cursor.fetchall()
    
    def xmen(self):
        self.cursor.execute("SELECT * FROM movies WHERE catagory='XMen'")
        return self.cursor.fetchall()
    
    def alteredcarbon(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def cowboybebop(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def fallout(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def forallmankind(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def foundation(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()

    def fubar(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()

    def hfor1923(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def halo(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def houseofthedragon(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def lostinspace(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def mastersoftheuniverse(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def monarchlegacyofmonsters(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def nightsky(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def orville(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def prehistoricplanet(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def raisedbywolves(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def shogun(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()

    def silo(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def thecontinental(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def thelastofus(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def thelordoftheringstheringsofpower(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def wheeloftime(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def discovery(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def enterprise(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def lowerdecks(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def picard(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def prodigy(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()

    def sttv(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def strangenewworlds(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()

    def tng(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def voyager(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()

    def acolyte(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def ahsoka(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def andor(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def bookofbobafett(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def mandalorian(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def obiwankenobi(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def talesoftheempire(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def talesofthejedi(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def thebadbatch(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()

    def visions(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def falconandthewintersoldier(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def hawkeye(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def iamgroot(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def loki(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()

    def moonknight(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()
    
    def secretinvasion(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()

    def shehulk(self, mediaid):
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()

    def wandavision(self, mediaid): 
        command = f"SELECT * FROM tvshows WHERE tvid={mediaid}"
        self.cursor.execute(command)
        return self.cursor.fetchall()

