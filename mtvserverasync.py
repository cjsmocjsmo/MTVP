import vlc
import time
import asyncio
import websockets
import json
import json
import logging
import mtvserverutils

# Initialize VLC player
instance = vlc.Instance()
player = instance.media_player_new()

# Configure logging
logging.basicConfig(level=logging.INFO)

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

            elif command == "action":
                action_data = mtvserverutils.Media().action()
                await websocket.send(json.dumps(action_data))

            elif command == "arnold":
                arnold_data = mtvserverutils.Media().arnold()
                await websocket.send(json.dumps(arnold_data))

            elif command == "brucelee":
                brucelee_data = mtvserverutils.Media().brucelee()
                await websocket.send(json.dumps(brucelee_data))

            elif command == "brucewillis":
                brucewillis_data = mtvserverutils.Media().brucewillis()
                await websocket.send(json.dumps(brucewillis_data))

            elif command == "buzz":
                buzz_data = mtvserverutils.Media().buzz()
                await websocket.send(json.dumps(buzz_data))

            elif command == "cartoons":
                cartoons_data = mtvserverutils.Media().cartoons()
                await websocket.send(json.dumps(cartoons_data))

            elif command == "charliebrown":
                charliebrown_data = mtvserverutils.Media().charliebrown()
                await websocket.send(json.dumps(charliebrown_data))

            elif command == "comedy":
                comedy_data = mtvserverutils.Media().comedy()
                await websocket.send(json.dumps(comedy_data))

            elif command == "chucknorris":
                chucknorris_data = mtvserverutils.Media().chucknorris()
                await websocket.send(json.dumps(chucknorris_data))         
 


    except Exception as e:
        logging.error(f"Exception in handle_message: {e}")
    finally:
        logging.info("WebSocket connection closed")
async def main():
    async with websockets.serve(handle_message, "192.168.0.113", 8765):
        await asyncio.Future()  # Run forever

if __name__ == "__main__":
    asyncio.run(main())