import vlc
import time
import asyncio
import websockets
import json
import json
import logging
import mtvserverutils
from dotenv import load_dotenv

load_dotenv()
# Initialize VLC player
instance = vlc.Instance()
player = instance.media_player_new()

# Configure logging
logging.basicConfig(level=logging.INFO)

MTVMEDIA = mtvserverutils.Media()

# async def handle_message(websocket, path):
async def handle_message(websocket):
    try:
        async for message in websocket:
            data = json.loads(message)
            command = data.get("command")
            
            if command == "set_media":
                media_path = data.get("media_path")
                if media_path:
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
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.alteredcarbon(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "columbia":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.columbia(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "cowboybebop":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.cowboybebop(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "fallout":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.fallout(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "forallmankind":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.forallmankind(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "foundation":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.foundation(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "fubar":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.fubar(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "hford1923":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.hford1923(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "halo":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.halo(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "houseofthedragon":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.houseofthedragon(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "lostinspace":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.lostinspace(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "mastersoftheuniverse":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.mastersoftheuniverse(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "monarchlegacyofmonsters":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.monarchlegacyofmonsters(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "nightsky":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.nightsky(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "orville":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.orville(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "prehistoricplanet":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.prehistoricplanet(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "raisedbywolves":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.raisedbywolves(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "shogun":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.shogun(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "silo":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.silo(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "thecontinental":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.thecontinental(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "thelastofus":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.thelastofus(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "thelordoftheringstheringsofpower":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.thelordoftheringstheringsofpower(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "wheeloftime":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.wheeloftime(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "discovery":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.discovery(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "enterprise":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.enterprise(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "lowerdecks":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.lowerdecks(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "picard":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.picard(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "prodigy":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.prodigy(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "sttv":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.sttv(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "strangenewworlds":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.strangenewworlds(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "tng":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.tng(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "voyager":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.voyager(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "acolyte":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.acolyte(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "ahsoka":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.ahsoka(media_id)
                    await websocket.send(json.dumps(mediainfo))
            
            elif command == "andor":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.andor(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "bookofbobafett":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.bookofbobafett(media_id)
                    await websocket.send(json.dumps(mediainfo))
                
            elif command == "mandalorian":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.mandalorian(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "obiwankenobi":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.obiwankenobi(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "talesoftheempire":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.talesoftheempire(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "talesofthejedi":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.talesofthejedi(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "thebadbatch":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.thebadbatch(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "visions":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.visions(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "falconwintersoldier":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.falconwintersoldier(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "hawkeye":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.hawkeye(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "iamgroot":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.iamgroot(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "loki":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.loki(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "moonknight":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.moonknight(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "secretinvasion":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.secretinvasion(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "shehulk":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.shehulk(media_id)
                    await websocket.send(json.dumps(mediainfo))

            elif command == "wandavision":
                media_id = data.get("media_id")
                if media_id:
                    mediainfo = MTVMEDIA.wandavision(media_id)
                    await websocket.send(json.dumps(mediainfo))

    except Exception as e:
        logging.error(f"Exception in handle_message: {e}")
    finally:
        logging.info("WebSocket connection closed")
async def main():
    async with websockets.serve(handle_message, "192.168.0.113", 8765):
        await asyncio.Future()  # Run forever

if __name__ == "__main__":
    asyncio.run(main())