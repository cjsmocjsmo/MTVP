import vlc
import time
import asyncio
import websockets
import json
import logging
import mtvserverutils
from dotenv import load_dotenv
import sqlite3
import os


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
                    media_path = await get_media_path_from_media_id()
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
                player.set_time(current_time + 30000)
                await websocket.send(json.dumps({"status": "next"}))


            elif command == "test":
                await websocket.send(json.dumps({"status": "Fuck it worked"}))

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

            elif command == "harrypotter":
                harrypotter_data = MTVMEDIA.harrypotter()
                await websocket.send(json.dumps(harrypotter_data))

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
                kingsmen_data = MTVMEDIA.kingsman()
                await websocket.send(json.dumps(kingsmen_data))

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

            elif command == "xmen":
                xmen_data = MTVMEDIA.xmen()
                await websocket.send(json.dumps(xmen_data))

            elif command == "alteredcarbon":
                mediainfo = MTVMEDIA.alteredcarbon()
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

            elif command == "forallmankind":
                mediainfo = MTVMEDIA.forallmankind()
                await websocket.send(json.dumps(mediainfo))

            elif command == "foundation":
                mediainfo = MTVMEDIA.foundation()
                await websocket.send(json.dumps(mediainfo))

            elif command == "fuubar":
                mediainfo = MTVMEDIA.fuubar()
                await websocket.send(json.dumps(mediainfo))

            elif command == "hford1923":
                mediainfo = MTVMEDIA.hford1923()
                await websocket.send(json.dumps(mediainfo))

            elif command == "halo":
                mediainfo = MTVMEDIA.halo()
                await websocket.send(json.dumps(mediainfo))

            elif command == "houseofthedragon":
                mediainfo = MTVMEDIA.houseofthedragon()
                await websocket.send(json.dumps(mediainfo))

            elif command == "lostinspace":
                mediainfo = MTVMEDIA.lostinspace()
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

            elif command == "orville":
                mediainfo = MTVMEDIA.orville()
                await websocket.send(json.dumps(mediainfo))

            elif command == "prehistoricplanet":
                mediainfo = MTVMEDIA.prehistoricplanet()
                await websocket.send(json.dumps(mediainfo))

            elif command == "raisedbywolves":
                mediainfo = MTVMEDIA.raisedbywolves()
                await websocket.send(json.dumps(mediainfo))

            elif command == "shogun":
                mediainfo = MTVMEDIA.shogun()
                await websocket.send(json.dumps(mediainfo))

            elif command == "silo":
                mediainfo = MTVMEDIA.silo()
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

            elif command == "wheeloftime":
                mediainfo = MTVMEDIA.wheeloftime()
                await websocket.send(json.dumps(mediainfo))

            elif command == "discovery":
                mediainfo = MTVMEDIA.discovery()
                await websocket.send(json.dumps(mediainfo))

            elif command == "enterprise":
                mediainfo = MTVMEDIA.enterprise()
                await websocket.send(json.dumps(mediainfo))

            elif command == "lowerdecks":
                mediainfo = MTVMEDIA.lowerdecks()
                await websocket.send(json.dumps(mediainfo))

            elif command == "picard":
                mediainfo = MTVMEDIA.picard()
                await websocket.send(json.dumps(mediainfo))

            elif command == "prodigy":
                mediainfo = MTVMEDIA.prodigy()
                await websocket.send(json.dumps(mediainfo))

            elif command == "sttv":
                mediainfo = MTVMEDIA.sttv()
                await websocket.send(json.dumps(mediainfo))

            elif command == "strangenewworlds":
                mediainfo = MTVMEDIA.strangenewworlds()
                await websocket.send(json.dumps(mediainfo))

            elif command == "tng":
                mediainfo = MTVMEDIA.tng()
                await websocket.send(json.dumps(mediainfo))

            elif command == "voyager":
                mediainfo = MTVMEDIA.voyager()
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
                
            elif command == "mandalorian":
                mediainfo = MTVMEDIA.mandalorian()
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

            elif command == "thebadbatch":
                mediainfo = MTVMEDIA.thebadbatch()
                await websocket.send(json.dumps(mediainfo))

            elif command == "visions":
                mediainfo = MTVMEDIA.visions()
                await websocket.send(json.dumps(mediainfo))

            elif command == "falconwintersoldier":
                mediainfo = MTVMEDIA.falconwintersoldier()
                await websocket.send(json.dumps(mediainfo))

            elif command == "hawkeye":
                mediainfo = MTVMEDIA.hawkeye()
                await websocket.send(json.dumps(mediainfo))

            elif command == "iamgroot":
                mediainfo = MTVMEDIA.iamgroot()
                await websocket.send(json.dumps(mediainfo))

            elif command == "loki":
                mediainfo = MTVMEDIA.loki()
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

    except Exception as e:
        logging.error(f"Exception in handle_message: {e}")
    finally:
        logging.info("WebSocket connection closed")
async def servermain():
    async with websockets.serve(handle_message, "10.0.4.41", 8765):
        await asyncio.Future()  # Run forever

if __name__ == "__main__":
    asyncio.run(servermain())