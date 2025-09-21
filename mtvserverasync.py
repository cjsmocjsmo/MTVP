import vlc
# import time
import asyncio
import websockets
import requests
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

logging.basicConfig(
    level=logging.INFO,
    filename='/usr/share/MTV/mtv.log',
    filemode='a',  # Append mode
    format='%(asctime)s - %(levelname)s - %(message)s'
)

MTVMEDIA = mtvserverutils.Media()

async def get_media_path_from_media_id(media_id):
    """
    Retrieves the media_path from the db using the media_id.
    """
    try:
        # Use 'with' to manage the database connection
        with sqlite3.connect(os.getenv("MTV_DB_PATH")) as conn:
            cursor = conn.cursor()
            cursor.execute("SELECT Path FROM movies WHERE MovId = ?", (media_id,))
            result = cursor.fetchone()
            if result:
                media_path = result[0]
                return media_path
            else:
                logging.error(f"No media path found for media_id: {media_id}")
                return None
    except sqlite3.Error as e:
        logging.error(f"SQLite error: {e}")
        return None
    except Exception as e:
        logging.error(f"Error fetching media path: {e}")
        return None

async def get_media_path_from_media_tv_id(media_tv_id):
    """
    Retrieves the media_path from the db using the media_tv_id.
    """
    try:
        with sqlite3.connect(os.getenv("MTV_DB_PATH")) as conn:
            cursor = conn.cursor()
            cursor.execute("SELECT Path FROM tvshows WHERE TvId = ?", (media_tv_id,))
            result = cursor.fetchone()
            if result:
                media_path = result[0]
                return media_path
            else:
                logging.error(f"No media path found for media_tv_id: {media_tv_id}")
                return None
    except sqlite3.Error as e:
        logging.error(f"SQLite error: {e}")
        return None
    except Exception as e:
        logging.error(f"Error fetching media path: {e}")
        return None      

async def get_weather_for_belfair_wa():
    """
    Retrieves and prints the current weather conditions for Belfair, WA
    from the National Weather Service.
    """
    try:
        latitude = 47.4281
        longitude = -122.8189
        point_url = f"https://api.weather.gov/points/{latitude},{longitude}"
        point_response = requests.get(point_url)
        point_response.raise_for_status()  # Raise an exception for bad status codes
        point_data = point_response.json()
        forecast_url = point_data['properties']['forecastHourly']
        weather_response = requests.get(forecast_url)
        weather_response.raise_for_status()
        weather_data = weather_response.json()
        current_forecast = weather_data['properties']['periods'][0]

        weather_data = {
            "location": "Belfair, WA",
            "temperature": current_forecast['temperature'],
            "temperature_unit": current_forecast['temperatureUnit'],
            "conditions": current_forecast['shortForecast'],
            "winddirection": current_forecast['windDirection'],
            "windspeed": current_forecast['windSpeed']
        }

        return weather_data

    except requests.exceptions.RequestException as e:
        # print(f"Error fetching weather data: {e}")
        logging.error(f"Error fetching weather data: {e}")
        return None
    except (KeyError, IndexError) as e:
        # print(f"Error parsing weather data: {e}")
        logging.error(f"Error parsing weather data: {e}")
        return None
    except Exception as e:
        logging.error(f"Error fetching media path: {e}")
        return None      


# async def handle_message(websocket, path):
async def handle_message(websocket):
    try:
        async for message in websocket:
            data = json.loads(message)
            command = data.get("command")

            if command == "set_media":
                try:
                    media_id = data.get("media_id")
                    if media_id:
                        media_path = await get_media_path_from_media_id(media_id)
                        player.set_media(vlc.Media(media_path))
                        player.set_fullscreen(True)
                        print(f"Starting mediaplayer with the path:\n{media_path}")
                        logging.info(f"Starting mediaplayer with the path:\n{media_path}")
                        await websocket.send(json.dumps({"status": "media_set"}))
                except Exception as e:
                    logging.error(f"Error fetching media path: {e}")
                    return None
             
            elif command == "set_tv_media":
                try:
                    media_tv_id = data.get("media_tv_id")
                    if media_tv_id:
                        media_path = await get_media_path_from_media_tv_id(media_tv_id)
                        player.set_media(vlc.Media(media_path))
                        player.set_fullscreen(True)
                        print(f"Starting mediaplayer with the path:\n{media_path}")
                        logging.info(f"Starting mediaplayer with the path:\n{media_path}")
                        await websocket.send(json.dumps({"status": "media_set"}))
                except Exception as e:
                    logging.error(f"Error setting player path with mediapath: {e}")
                    return None

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

            elif command == "weather":
                weather_data = await get_weather_for_belfair_wa()
                await websocket.send(json.dumps(weather_data))

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

            elif command == "cheechandchong":
                cheechandchong_data = MTVMEDIA.cheechandchong()
                await websocket.send(json.dumps(cheechandchong_data))

            elif command == "chucknorris":
                chucknorris_data = MTVMEDIA.chucknorris()
                await websocket.send(json.dumps(chucknorris_data)) 

            elif command == "comedy":
                comedy_data = MTVMEDIA.comedy()
                await websocket.send(json.dumps(comedy_data))

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

            elif command == "musicvids":
                musicvids_data = MTVMEDIA.musicvids()
                await websocket.send(json.dumps(musicvids_data))

            elif command == "nature":
                nature_data = MTVMEDIA.nature()
                await websocket.send(json.dumps(nature_data))

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

            elif command == "science":
                science_data = MTVMEDIA.science()
                await websocket.send(json.dumps(science_data))

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

            elif command == "stooges":
                stooges_data = MTVMEDIA.stooges()
                await websocket.send(json.dumps(stooges_data))

            elif command == "superheros":
                superheros_data = MTVMEDIA.superheros()
                await websocket.send(json.dumps(superheros_data))

            elif command == "superman":
                superman_data = MTVMEDIA.superman()
                await websocket.send(json.dumps(superman_data))

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

            elif command == "trolls":
                trolls_data = MTVMEDIA.trolls()
                await websocket.send(json.dumps(trolls_data))

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

            elif command == "columbias1":
                mediainfo = MTVMEDIA.columbia()
                await websocket.send(json.dumps(mediainfo))

            elif command == "cowboybebops1":
                mediainfo = MTVMEDIA.cowboybebop()
                await websocket.send(json.dumps(mediainfo))

            elif command == "fallouts1":
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

            elif command == "foundations3":
                mediainfo = MTVMEDIA.foundations3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "fuubars1":
                mediainfo = MTVMEDIA.fuubar1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "fuubars2":
                mediainfo = MTVMEDIA.fuubar2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "hford1923s1":
                mediainfo = MTVMEDIA.hford19231()
                await websocket.send(json.dumps(mediainfo))

            elif command == "hford1923s2":
                mediainfo = MTVMEDIA.hford19232()
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

            # elif command == "mastersoftheuniverse":
            #     mediainfo = MTVMEDIA.mastersoftheuniverse()
            #     await websocket.send(json.dumps(mediainfo))

            elif command == "monarchlegacyofmonsterss1":
                mediainfo = MTVMEDIA.monarchlegacyofmonsters()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nightskys1":
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

            elif command == "shoguns1":
                mediainfo = MTVMEDIA.shogun()
                await websocket.send(json.dumps(mediainfo))

            elif command == "silos1":
                mediainfo = MTVMEDIA.silo1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "silos2":
                mediainfo = MTVMEDIA.silo2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "thecontinentals1":
                mediainfo = MTVMEDIA.thecontinental()
                await websocket.send(json.dumps(mediainfo))

            elif command == "thelastofus1":
                mediainfo = MTVMEDIA.thelastofus1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "thelastofus2":
                mediainfo = MTVMEDIA.thelastofus2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "thelordoftheringstheringsofpowers1":
                mediainfo = MTVMEDIA.thelordoftheringstheringsofpower()
                await websocket.send(json.dumps(mediainfo))

            elif command == "wheeloftimes1":
                mediainfo = MTVMEDIA.wheeloftime1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "wheeloftimes2":
                mediainfo = MTVMEDIA.wheeloftime2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "wheeloftimes3":
                mediainfo = MTVMEDIA.wheeloftime3()
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

            elif command == "prodigys1":
                mediainfo = MTVMEDIA.prodigy1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "prodigys2":
                mediainfo = MTVMEDIA.prodigy2()
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

            elif command == "strangenewworldss3":
                mediainfo = MTVMEDIA.strangenewworldss3()
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

            elif command == "deepspacenines1":
                mediainfo = MTVMEDIA.deepspacenines1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "deepspacenines2":
                mediainfo = MTVMEDIA.deepspacenines2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "deepspacenines3":
                mediainfo = MTVMEDIA.deepspacenines3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "deepspacenines4":
                mediainfo = MTVMEDIA.deepspacenines4()
                await websocket.send(json.dumps(mediainfo))

            elif command == "deepspacenines5":
                mediainfo = MTVMEDIA.deepspacenines5()
                await websocket.send(json.dumps(mediainfo))

            elif command == "deepspacenines6":
                mediainfo = MTVMEDIA.deepspacenines6()
                await websocket.send(json.dumps(mediainfo))

            elif command == "deepspacenines7":
                mediainfo = MTVMEDIA.deepspacenines7()
                await websocket.send(json.dumps(mediainfo))

            elif command == "acolytes1":
                mediainfo = MTVMEDIA.acolyte()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ahsokas1":
                mediainfo = MTVMEDIA.ahsoka()
                await websocket.send(json.dumps(mediainfo))
            
            elif command == "andors1":
                mediainfo = MTVMEDIA.andor1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "andors2":
                mediainfo = MTVMEDIA.andor2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "bookofbobafetts1":
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

            elif command == "obiwankenobis1":
                mediainfo = MTVMEDIA.obiwankenobi()
                await websocket.send(json.dumps(mediainfo))

            elif command == "talesoftheempires1":
                mediainfo = MTVMEDIA.talesoftheempire()
                await websocket.send(json.dumps(mediainfo))

            elif command == "talesofthejedis1":
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

            elif command == "falconwintersoldiers1":
                mediainfo = MTVMEDIA.falconwintersoldier()
                await websocket.send(json.dumps(mediainfo))

            elif command == "hawkeyes1":
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

            elif command == "moonknights1":
                mediainfo = MTVMEDIA.moonknight()
                await websocket.send(json.dumps(mediainfo))

            elif command == "secretinvasions1":
                mediainfo = MTVMEDIA.secretinvasion()
                await websocket.send(json.dumps(mediainfo))

            elif command == "shehulks1":
                mediainfo = MTVMEDIA.shehulk()
                await websocket.send(json.dumps(mediainfo))

            elif command == "wandavisions1":
                mediainfo = MTVMEDIA.wandavision()
                await websocket.send(json.dumps(mediainfo))

            elif command == "skeletoncrews1":
                mediainfo = MTVMEDIA.skeletoncrew()
                await websocket.send(json.dumps(mediainfo))
            
            elif command == "moblands1":
                mediainfo = MTVMEDIA.mobland()
                await websocket.send(json.dumps(mediainfo))

            elif command == "talesofthejedis1":
                mediainfo = MTVMEDIA.talesofthejedi()
                await websocket.send(json.dumps(mediainfo))

            elif command == "talesoftheunderworlds1":
                mediainfo = MTVMEDIA.talesoftheunderworld()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ironhearts1":
                mediainfo = MTVMEDIA.ironheart()
                await websocket.send(json.dumps(mediainfo))

            elif command == "wednesdays1":
                mediainfo = MTVMEDIA.wednesday1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "wednesdays2":
                mediainfo = MTVMEDIA.wednesday2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "tonyandzivas1":
                mediainfo = MTVMEDIA.tonyandzivas1()
                await websocket.send(json.dumps(mediainfo))
            
            elif command == "ncishawaiis1":
                mediainfo = MTVMEDIA.ncishawaiis1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncishawaiis2":
                mediainfo = MTVMEDIA.ncishawaiis2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncishawaiis3":
                mediainfo = MTVMEDIA.ncishawaiis3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncisneworleanss1":
                mediainfo = MTVMEDIA.ncisneworleanss1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncisneworleanss2":
                mediainfo = MTVMEDIA.ncisneworleanss2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncisneworleanss3":
                mediainfo = MTVMEDIA.ncisneworleanss3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncisneworleanss4":
                mediainfo = MTVMEDIA.ncisneworleanss4()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncisneworleanss5":
                mediainfo = MTVMEDIA.ncisneworleanss5()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncisneworleanss6":
                mediainfo = MTVMEDIA.ncisneworleanss6()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncisneworleanss7":
                mediainfo = MTVMEDIA.ncisneworleanss7()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncissydneys1":
                mediainfo = MTVMEDIA.ncissydneys1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncissydneys2":
                mediainfo = MTVMEDIA.ncissydneys2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncissydneys3":
                mediainfo = MTVMEDIA.ncissydneys3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncisoriginss1":
                mediainfo = MTVMEDIA.ncisoriginss1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss1":
                mediainfo = MTVMEDIA.ncis()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss2":
                mediainfo = MTVMEDIA.nciss2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss3":
                mediainfo = MTVMEDIA.nciss3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss4":
                mediainfo = MTVMEDIA.nciss4()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss5":
                mediainfo = MTVMEDIA.nciss5()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss6":
                mediainfo = MTVMEDIA.nciss6()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss7":
                mediainfo = MTVMEDIA.nciss7()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss8":
                mediainfo = MTVMEDIA.nciss8()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss9":
                mediainfo = MTVMEDIA.nciss9()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss10":
                mediainfo = MTVMEDIA.nciss10()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss11":
                mediainfo = MTVMEDIA.nciss11()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss12":
                mediainfo = MTVMEDIA.nciss12()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss13":
                mediainfo = MTVMEDIA.nciss13()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss14":
                mediainfo = MTVMEDIA.nciss14()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss15":
                mediainfo = MTVMEDIA.nciss15()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss16":
                mediainfo = MTVMEDIA.nciss16()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss17":
                mediainfo = MTVMEDIA.nciss17()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss18":
                mediainfo = MTVMEDIA.nciss18()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss19":
                mediainfo = MTVMEDIA.nciss19()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss20":
                mediainfo = MTVMEDIA.nciss20()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss21":
                mediainfo = MTVMEDIA.nciss21()
                await websocket.send(json.dumps(mediainfo))

            elif command == "nciss22":
                mediainfo = MTVMEDIA.nciss22()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncislas1":
                mediainfo = MTVMEDIA.ncislas1()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncislas2":
                mediainfo = MTVMEDIA.ncislas2()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncislas3":
                mediainfo = MTVMEDIA.ncislas3()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncislas4":
                mediainfo = MTVMEDIA.ncislas4()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncislas5":
                mediainfo = MTVMEDIA.ncislas5()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncislas6":
                mediainfo = MTVMEDIA.ncislas6()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncislas7":
                mediainfo = MTVMEDIA.ncislas7()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncislas8":
                mediainfo = MTVMEDIA.ncislas8()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncislas9":
                mediainfo = MTVMEDIA.ncislas9()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncislas10":
                mediainfo = MTVMEDIA.ncislas10()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncislas11":
                mediainfo = MTVMEDIA.ncislas11()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncislas12":
                mediainfo = MTVMEDIA.ncislas12()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncislas13":
                mediainfo = MTVMEDIA.ncislas13()
                await websocket.send(json.dumps(mediainfo))

            elif command == "ncislas14":
                mediainfo = MTVMEDIA.ncislas14()
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