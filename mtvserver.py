import tornado.ioloop
import tornado.web
import tornado.websocket

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
        print("WebSocket opened")

    def on_message(self, message):
        if message == "pause":
            # Pause playback
            print("Playback paused")
        elif message == "stop":
            # Stop playback
            print("Playback stopped")

    def on_close(self):
        self.clients.remove(self)
        print("WebSocket closed")

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
    app.listen(8888)
    tornado.ioloop.IOLoop.current().start()