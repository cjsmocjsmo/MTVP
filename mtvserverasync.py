import vlc
# import time
import asyncio
import websockets
import json
import logging
import mtvserverutils
from dotenv import load_dotenv
import sqlite3
import os
import utils as UTILS


# Initialize VLC player
instance = vlc.Instance()
player = instance.media_player_new()

load_dotenv()

logging.basicConfig(level=logging.INFO)

MTVMEDIA = mtvserverutils.Media()

async def get_media_path_from_media_id(media_id):
    conn = sqlite3.connect(os.getenv("MTV_DB_PATH"))
    cursor = conn.cursor()
    cursor.execute("SELECT Path FROM movies WHERE MovId = ?", (media_id,))
    media_path = cursor.fetchone()[0]
    print(f"Media path:\n{media_path}")
    conn.close()
    return media_path

async def get_media_path_from_media_tv_id(media_tv_id):
    conn = sqlite3.connect(os.getenv("MTV_DB_PATH"))
    cursor = conn.cursor()
    cursor.execute("SELECT Path FROM tvshows WHERE TvId = ?", (media_tv_id,))
    media_path = cursor.fetchone()[0]
    print(f"Media path:\n{media_path}")
    conn.close()
    return media_path


# async def handle_message(websocket, path):
async def handle_message(websocket):
    try:
        async for message in websocket:
            data = json.loads(message)
            command = data.get("command")

            if command == "set_media":
                media_id = data.get("media_id")
                if media_id:
                    media_path = await get_media_path_from_media_id(media_id)
                    print(f"Starting mediaplayer with the path:\n{media_path}")
                    player.set_media(vlc.Media(media_path))
                    player.set_fullscreen(True)
                    await websocket.send(json.dumps({"status": "media_set"}))

            elif command == "set_tv_media":
                media_tv_id = data.get("media_tv_id")
                if media_tv_id:
                    media_path = await get_media_path_from_media_tv_id(media_tv_id)
                    print(f"Starting mediaplayer with the path:\n{media_path}")
                    player.set_media(vlc.Media(media_path))
                    player.set_fullscreen(True)
                    await websocket.send(json.dumps({"status": "media_set"}))

            elif command == "search":
                phrase = data.get("phrase")
                if phrase:
                    search_results = MTVMEDIA.search(phrase)
                    await websocket.send(json.dumps(search_results))

            elif command == "play":
                player.play()
                await websocket.send(json.dumps({"status": "playing"}))
            
            elif command == "pause":
                player.pause()
                await websocket.send(json.dumps({"status": "paused"}))

            elif command == "stop":
                player.stop()
                await websocket.send(json.dumps({"status": "stopped"}))

            elif command == "next":
                current_time = player.get_time()
                player.set_time(current_time + 35000)
                await websocket.send(json.dumps({"status": "next"}))

            elif command == "previous":
                current_time = player.get_time()
                player.set_time(current_time - 35000)
                await websocket.send(json.dumps({"status": "previous"}))

            elif command == "test":
                await websocket.send(json.dumps({"status": "Fuck it worked"}))

            elif command == "movcount":
                mov_count = UTILS.movie_count()
                await websocket.send(json.dumps(mov_count))

            elif command == "tvcount":
                tv_count = UTILS.tvshow_count()
                await websocket.send(json.dumps(tv_count))

            elif command == "movsizeondisk":
                movsizeondisk = UTILS.movies_size_on_disk()
                await websocket.send(json.dumps(movsizeondisk))

            elif command == "tvsizeondisk":
                tvsizeondisk = UTILS.tvshows_size_on_disk()
                await websocket.send(json.dumps(tvsizeondisk))




            elif command == "gallonstotal":
                gallons_total = UTILS.propane_gallons_total()
                await websocket.send(json.dumps(gallons_total))

            elif command == "amounttotal":
                amount_total = UTILS.propane_amount_total()
                await websocket.send(json.dumps(amount_total))

            elif command == "insertgallons":
                gallons = data.get("gallons")
                if gallons:
                    UTILS.insert_gallons(gallons)
                    await websocket.send(json.dumps({"status": "gallons_inserted"}))

            elif command == "insertamount":
                amount = data.get("amount")
                if amount:
                    UTILS.insert_amount(amount)
                    await websocket.send(json.dumps({"status": "amount_inserted"}))






            elif command == "action":
                action_data = MTVMEDIA.action()
                await websocket.send(json.dumps(action_data))

            elif command == "arnold":
                arnold_data = MTVMEDIA.arnold()
                await websocket.send(json.dumps(arnold_data))

            elif command == "brucelee":
                brucelee_data = MTVMEDIA.brucelee()
                await websocket.send(json.dumps(brucelee_data))

            elif command == "brucewillis":
                brucewillis_data = MTVMEDIA.brucewillis()
                await websocket.send(json.dumps(brucewillis_data))

            elif command == "buzz":
                buzz_data = MTVMEDIA.buzz()
                await websocket.send(json.dumps(buzz_data))

            elif command == "cartoons":
                cartoons_data = MTVMEDIA.cartoons()
                await websocket.send(json.dumps(cartoons_data))

            elif command == "charliebrown":
                charliebrown_data = MTVMEDIA.charliebrown()
                await websocket.send(json.dumps(charliebrown_data))

            elif command == "comedy":
                comedy_data = MTVMEDIA.comedy()
                await websocket.send(json.dumps(comedy_data))

            elif command == "chucknorris":
                chucknorris_data = MTVMEDIA.chucknorris()
                await websocket.send(json.dumps(chucknorris_data)) 

            elif command == "documentary":
                documentary_data = MTVMEDIA.documentary()
                await websocket.send(json.dumps(documentary_data))

            elif command == "drama":
                drama_data = MTVMEDIA.drama()
                await websocket.send(json.dumps(drama_data))

            elif command == "fantasy":
                fantasy_data = MTVMEDIA.fantasy()
                await websocket.send(json.dumps(fantasy_data))  

            elif command == "ghostbusters":
                ghostbusters_data = MTVMEDIA.ghostbusters()
                await websocket.send(json.dumps(ghostbusters_data))

            elif command == "godzilla":
                godzilla_data = MTVMEDIA.godzilla()
                await websocket.send(json.dumps(godzilla_data))

            elif command == "harrisonford":
                harrisonford_data = MTVMEDIA.harrisonford()
                await websocket.send(json.dumps(harrisonford_data))

            elif command == "harrypotter":
                harrypotter_data = MTVMEDIA.harrypotter()
                await websocket.send(json.dumps(harrypotter_data))

            elif command == "hellboy":
                hellboy_data = MTVMEDIA.hellboy()
                await websocket.send(json.dumps(hellboy_data))

            elif command == "indianajones":
                indianajones_data = MTVMEDIA.indianajones()
                await websocket.send(json.dumps(indianajones_data))

            elif command == "jamesbond":
                jamesbond_data = MTVMEDIA.jamesbond()
                await websocket.send(json.dumps(jamesbond_data))

            elif command == "johnwayne":
                johnwayne_data = MTVMEDIA.johnwayne()
                await websocket.send(json.dumps(johnwayne_data))

            elif command == "johnwick":
                johnwick_data = MTVMEDIA.johnwick()
                await websocket.send(json.dumps(johnwick_data))

            elif command == "jurassicpark":
                jurassicpark_data = MTVMEDIA.jurassicpark()
                await websocket.send(json.dumps(jurassicpark_data))

            elif command == "kevincostner":
                kevincostner_data = MTVMEDIA.kevincostner()
                await websocket.send(json.dumps(kevincostner_data))

            elif command == "kingsman":
                kingsman_data = MTVMEDIA.kingsman()
                await websocket.send(json.dumps(kingsman_data))

            elif command == "lego":
                lego_data = MTVMEDIA.lego()
                await websocket.send(json.dumps(lego_data))

            elif command == "meninblack":
                meninblack_data = MTVMEDIA.meninblack()
                await websocket.send(json.dumps(meninblack_data))

            elif command == "minions":
                minions_data = MTVMEDIA.minions()
                await websocket.send(json.dumps(minions_data))

            elif command == "misc":
                misc_data = MTVMEDIA.misc()
                await websocket.send(json.dumps(misc_data))

            elif command == "nicolascage":
                nicolascage_data = MTVMEDIA.nicolascage()
                await websocket.send(json.dumps(nicolascage_data))

            elif command == "oldies":
                oldies_data = MTVMEDIA.oldies()
                await websocket.send(json.dumps(oldies_data))

            elif command == "panda":
                panda_data = MTVMEDIA.panda()
                await websocket.send(json.dumps(panda_data))

            elif command == "pirates":
                pirates_data = MTVMEDIA.pirates()
                await websocket.send(json.dumps(pirates_data))

            elif command == "riddick":
                riddick_data = MTVMEDIA.riddick()
                await websocket.send(json.dumps(riddick_data))

            elif command == "scifi":
                scifi_data = MTVMEDIA.scifi()
                await websocket.send(json.dumps(scifi_data))

            elif command == "stalone":
                stalone_data = MTVMEDIA.stalone()
                await websocket.send(json.dumps(stalone_data))

            elif command == "startrek":
                startrek_data = MTVMEDIA.startrek()
                await websocket.send(json.dumps(startrek_data))

            elif command == "starwars":
                starwars_data = MTVMEDIA.starwars()
                await websocket.send(json.dumps(starwars_data))

            elif command == "superheros":
                superheros_data = MTVMEDIA.superheros()
                await websocket.send(json.dumps(superheros_data))

            elif command == "therock":
                therock_data = MTVMEDIA.therock()
                await websocket.send(json.dumps(therock_data))

            elif command == "tinkerbell":
                tinkerbell_data = MTVMEDIA.tinkerbell()
                await websocket.send(json.dumps(tinkerbell_data))

            elif command == "tomcruize":
                tomcruize_data = MTVMEDIA.tomcruize()
                await websocket.send(json.dumps(tomcruize_data))

            elif command == "transformers":
                transformers_data = MTVMEDIA.transformers()
                await websocket.send(json.dumps(transformers_data))

            elif command == "tremors":
                tremors_data = MTVMEDIA.tremors()
                await websocket.send(json.dumps(tremors_data))

            elif command == "vandam":
                vandam_data = MTVMEDIA.vandam()
                await websocket.send(json.dumps(vandam_data))

            elif command == "xmen":
                xmen_data = MTVMEDIA.xmen()
                await websocket.send(json.dumps(xmen_data))

            elif command == "alteredcarbons1":
                mediainfo = MTVMEDIA.alteredcarbons1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "alteredcarbons2":
                mediainfo = MTVMEDIA.alteredcarbons2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "columbia":
                mediainfo = MTVMEDIA.columbia()
                await websocket.send(json.dumps(mediainfo))

            elif command == "cowboybebop":
                mediainfo = MTVMEDIA.cowboybebop()
                await websocket.send(json.dumps(mediainfo))

            elif command == "fallout":
                mediainfo = MTVMEDIA.fallout()
                await websocket.send(json.dumps(mediainfo))

            elif command == "forallmankinds1":
                mediainfo = MTVMEDIA.forallmankinds1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "forallmankinds2":
                mediainfo = MTVMEDIA.forallmankinds2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "forallmankinds3":
                mediainfo = MTVMEDIA.forallmankinds3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "forallmankinds4":
                mediainfo = MTVMEDIA.forallmankinds4()
                await websocket.send(json.dumps(mediainfo))

            elif command == "foundations1":
                mediainfo = MTVMEDIA.foundations1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "foundations2":
                mediainfo = MTVMEDIA.foundations2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "fuubar":
                mediainfo = MTVMEDIA.fuubar()
                await websocket.send(json.dumps(mediainfo))

            elif command == "hford1923":
                mediainfo = MTVMEDIA.hford1923()
                await websocket.send(json.dumps(mediainfo))

            elif command == "halos1":
                mediainfo = MTVMEDIA.halos1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "halos2":
                mediainfo = MTVMEDIA.halos2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "houseofthedragons1":
                mediainfo = MTVMEDIA.houseofthedragon_s1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "houseofthedragons2":
                mediainfo = MTVMEDIA.houseofthedragon_s2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "lostinspaces1":
                mediainfo = MTVMEDIA.lostinspaces1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "lostinspaces2":
                mediainfo = MTVMEDIA.lostinspaces2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "lostinspaces3":
                mediainfo = MTVMEDIA.lostinspaces3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "mastersoftheuniverse":
                mediainfo = MTVMEDIA.mastersoftheuniverse()
                await websocket.send(json.dumps(mediainfo))

            elif command == "monarchlegacyofmonsters":
                mediainfo = MTVMEDIA.monarchlegacyofmonsters()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nightsky":
                mediainfo = MTVMEDIA.nightsky()
                await websocket.send(json.dumps(mediainfo))

            elif command == "orvilles1":
                mediainfo = MTVMEDIA.orvilles1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "orvilles2":
                mediainfo = MTVMEDIA.orvilles2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "orvilles3":
                mediainfo = MTVMEDIA.orvilles3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "prehistoricplanets1":
                mediainfo = MTVMEDIA.prehistoricplanets1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "prehistoricplanets2":
                mediainfo = MTVMEDIA.prehistoricplanets2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "raisedbywolvess1":
                mediainfo = MTVMEDIA.raisedbywolvess1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "raisedbywolvess2":
                mediainfo = MTVMEDIA.raisedbywolvess2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "shogun":
                mediainfo = MTVMEDIA.shogun()
                await websocket.send(json.dumps(mediainfo))

            elif command == "silo1":
                mediainfo = MTVMEDIA.silo1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "silo2":
                mediainfo = MTVMEDIA.silo2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "thecontinental":
                mediainfo = MTVMEDIA.thecontinental()
                await websocket.send(json.dumps(mediainfo))

            elif command == "thelastofus":
                mediainfo = MTVMEDIA.thelastofus()
                await websocket.send(json.dumps(mediainfo))

            elif command == "thelordoftheringstheringsofpower":
                mediainfo = MTVMEDIA.thelordoftheringstheringsofpower()
                await websocket.send(json.dumps(mediainfo))

            elif command == "wheeloftimes1":
                mediainfo = MTVMEDIA.wheeloftimes1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "wheeloftimes2":
                mediainfo = MTVMEDIA.wheeloftimes2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "discoverys1":
                mediainfo = MTVMEDIA.discoverys1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "discoverys2":
                mediainfo = MTVMEDIA.discoverys2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "discoverys3":
                mediainfo = MTVMEDIA.discoverys3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "discoverys4":
                mediainfo = MTVMEDIA.discoverys4()
                await websocket.send(json.dumps(mediainfo))

            elif command == 'discoverys5':
                mediainfo = MTVMEDIA.discoverys5()
                await websocket.send(json.dumps(mediainfo))

            elif command == "enterprises1":
                mediainfo = MTVMEDIA.enterprises1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "enterprises2":
                mediainfo = MTVMEDIA.enterprises2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "enterprises3":
                mediainfo = MTVMEDIA.enterprises3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "enterprises4":
                mediainfo = MTVMEDIA.enterprises4()
                await websocket.send(json.dumps(mediainfo))

            elif command == "enterprises5":
                mediainfo = MTVMEDIA.enterprises5()
                await websocket.send(json.dumps(mediainfo))

            elif command == "lowerdeckss1":
                mediainfo = MTVMEDIA.lowerdeckss1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "lowerdeckss2":
                mediainfo = MTVMEDIA.lowerdeckss2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "lowerdeckss3":
                mediainfo = MTVMEDIA.lowerdeckss3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "lowerdeckss4":
                mediainfo = MTVMEDIA.lowerdeckss4()
                await websocket.send(json.dumps(mediainfo))

            elif command == "lowerdeckss5":
                mediainfo = MTVMEDIA.lowerdeckss5()
                await websocket.send(json.dumps(mediainfo))

            elif command == "picards1":
                mediainfo = MTVMEDIA.picards1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "picards2":
                mediainfo = MTVMEDIA.picards2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "prodigy":
                mediainfo = MTVMEDIA.prodigy()
                await websocket.send(json.dumps(mediainfo))

            elif command == "sttvs1":
                mediainfo = MTVMEDIA.sttvs1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "sttvs2":
                mediainfo = MTVMEDIA.sttvs2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "sttvs3":
                mediainfo = MTVMEDIA.sttvs3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "strangenewworldss1":
                mediainfo = MTVMEDIA.strangenewworldss1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "strangenewworldss2":
                mediainfo = MTVMEDIA.strangenewworldss2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "tngs1":
                mediainfo = MTVMEDIA.tngs1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "tngs2":
                mediainfo = MTVMEDIA.tngs2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "tngs3":
                mediainfo = MTVMEDIA.tngs3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "tngs4":
                mediainfo = MTVMEDIA.tngs4()
                await websocket.send(json.dumps(mediainfo))

            elif command == "tngs5":
                mediainfo = MTVMEDIA.tngs5()
                await websocket.send(json.dumps(mediainfo))

            elif command == "tngs6":
                mediainfo = MTVMEDIA.tngs6()
                await websocket.send(json.dumps(mediainfo))

            elif command == "tngs7":
                mediainfo = MTVMEDIA.tngs7()
                await websocket.send(json.dumps(mediainfo))

            elif command == "voyagers1":
                mediainfo = MTVMEDIA.voyagers1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "voyagers2":
                mediainfo = MTVMEDIA.voyagers2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "voyagers3":
                mediainfo = MTVMEDIA.voyagers3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "voyagers4":
                mediainfo = MTVMEDIA.voyagers4()
                await websocket.send(json.dumps(mediainfo))

            elif command == "voyagers5":
                mediainfo = MTVMEDIA.voyagers5()
                await websocket.send(json.dumps(mediainfo))

            elif command == "voyagers6":
                mediainfo = MTVMEDIA.voyagers6()
                await websocket.send(json.dumps(mediainfo))

            elif command == "voyagers7":
                mediainfo = MTVMEDIA.voyagers7()
                await websocket.send(json.dumps(mediainfo))

            elif command == "acolyte":
                mediainfo = MTVMEDIA.acolyte()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ahsoka":
                mediainfo = MTVMEDIA.ahsoka()
                await websocket.send(json.dumps(mediainfo))
            
            elif command == "andor":
                mediainfo = MTVMEDIA.andor()
                await websocket.send(json.dumps(mediainfo))

            elif command == "bookofbobafett":
                mediainfo = MTVMEDIA.bookofbobafett()
                await websocket.send(json.dumps(mediainfo))
                
            elif command == "mandalorians1":
                mediainfo = MTVMEDIA.mandalorians1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "mandalorians2":
                mediainfo = MTVMEDIA.mandalorians2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "mandalorians3":
                mediainfo = MTVMEDIA.mandalorians3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "obiwankenobi":
                mediainfo = MTVMEDIA.obiwankenobi()
                await websocket.send(json.dumps(mediainfo))

            elif command == "talesoftheempire":
                mediainfo = MTVMEDIA.talesoftheempire()
                await websocket.send(json.dumps(mediainfo))

            elif command == "talesofthejedi":
                mediainfo = MTVMEDIA.talesofthejedi()
                await websocket.send(json.dumps(mediainfo))

            elif command == "thebadbatchs1":
                mediainfo = MTVMEDIA.thebadbatchs1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "thebadbatchs2":
                mediainfo = MTVMEDIA.thebadbatchs2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "thebadbatchs3":
                mediainfo = MTVMEDIA.thebadbatchs3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "visionss1":
                mediainfo = MTVMEDIA.visionss1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "visionss2":
                mediainfo = MTVMEDIA.visionss2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "falconwintersoldier":
                mediainfo = MTVMEDIA.falconwintersoldier()
                await websocket.send(json.dumps(mediainfo))

            elif command == "hawkeye":
                mediainfo = MTVMEDIA.hawkeye()
                await websocket.send(json.dumps(mediainfo))

            elif command == "iamgroots1":
                mediainfo = MTVMEDIA.iamgroots1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "iamgroots2":
                mediainfo = MTVMEDIA.iamgroots2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "lokis1":
                mediainfo = MTVMEDIA.lokis1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "lokis2":
                mediainfo = MTVMEDIA.lokis2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "moonknight":
                mediainfo = MTVMEDIA.moonknight()
                await websocket.send(json.dumps(mediainfo))

            elif command == "secretinvasion":
                mediainfo = MTVMEDIA.secretinvasion()
                await websocket.send(json.dumps(mediainfo))

            elif command == "shehulk":
                mediainfo = MTVMEDIA.shehulk()
                await websocket.send(json.dumps(mediainfo))

            elif command == "wandavision":
                mediainfo = MTVMEDIA.wandavision()
                await websocket.send(json.dumps(mediainfo))

            elif command == "skeletoncrew":
                mediainfo = MTVMEDIA.skeletoncrew()
                await websocket.send(json.dumps(mediainfo))

    except Exception as e:
        logging.error(f"Exception in handle_message: {e}")
    finally:
        logging.info("WebSocket connection closed")
async def servermain():
    async with websockets.serve(handle_message, "10.0.4.41", 8765):
        await asyncio.Future()  # Run forever

if __name__ == "__main__":
    asyncio.run(servermain())