#!/usr/bin/env python3

class ProcessTVShows:
    def __init__(self, tvshows):
        self.tvlist = tvshows

    def process(self):
        for tv in self.tvlist:
            print(tv)