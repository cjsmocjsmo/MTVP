#!/usr/bin/env python3

class ProcessMovies:
    def __init__(self, movs):
        self.movlist = movs

    def process(self):
        for mov in self.movlist:
            print(mov)