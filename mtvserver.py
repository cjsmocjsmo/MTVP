import tornado.ioloop
import tornado.web
import tornado.websocket
import logging
import os
from dotenv import load_dotenv

# Load environment variables
load_dotenv()

log_file = os.getenv('MTV_SERVER_LOG')
log_dir = os.path.dirname(log_file)
if not os.path.exists(log_dir):
    os.makedirs(log_dir, exist_ok=True)
if not os.path.exists(log_file):
    open(log_file, 'a').close()

logging.basicConfig(filename=log_file, level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

class IndexHandler(tornado.web.RequestHandler):
    def get(self):
        self.write("Index Page")

class HelloWorldHandler(tornado.web.RequestHandler):
    def get(self):
        self.write("Hello, world")

class WebSocketHandler(tornado.websocket.WebSocketHandler):
    clients = set()

    def open(self):
        self.clients.add(self)
        logging.info("WebSocket opened")

    def on_message(self, message):
        if message == "add":
            # Add a song to the playlist
            logging.info("Song added")
        elif message == "play":
            # Start playback
            logging.info("Playback started")
        elif message == "pause":
            # Pause playback
            logging.info("Playback paused")
        elif message == "stop":
            # Stop playback
            logging.info("Playback stopped")

    def on_close(self):
        self.clients.remove(self)
        logging.info("WebSocket closed")

    def check_origin(self, origin):
        return True  # Allow all origins

class CORSMiddleware(tornado.web.RequestHandler):
    def set_default_headers(self):
        self.set_header("Access-Control-Allow-Origin", "*")
        self.set_header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        self.set_header("Access-Control-Allow-Headers", "Content-Type, Authorization")

    def options(self):
        self.set_status(204)
        self.finish()

class Application(tornado.web.Application):
    def __init__(self):
        handlers = [
            (r"/", IndexHandler),
            (r"/hello", HelloWorldHandler),
            (r"/websocket", WebSocketHandler),
        ]
        settings = {
            "default_handler_class": CORSMiddleware,
        }
        super(Application, self).__init__(handlers, **settings)

if __name__ == "__main__":
    app = Application()
    app.listen(7777)
    tornado.ioloop.IOLoop.current().start()