import tornado.ioloop
import tornado.web
import tornado.websocket

class IndexHandler(tornado.web.RequestHandler):
    def get(self):
        self.render("index.html")

class WebSocketHandler(tornado.websocket.WebSocketHandler):
    clients = set()

    def open(self):
        self.clients.add(self)  


    def on_message(self, message):
        if message == "add":
            # Add item to playlist or perform other actions
            print("Item added")
        elif message == "play":
            # Start playback
            print("Playback started")
        elif message == "pause":
            # Pause playback
            print("Playback paused")
        elif message == "stop":
            # Stop playback
            print("Playback stopped")

    def on_close(self):
        self.clients.remove(self)

app = tornado.web.Application([
    (r"/", IndexHandler),
    (r"/websocket", WebSocketHandler),
])

if __name__ == "__main__":
    app.listen(8888)
    tornado.ioloop.IOLoop.current().start()