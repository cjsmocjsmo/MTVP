FROM arm32v7/nginx:mainline-alpine

RUN \
    mkdir /usr/share/nginx/html/_app && \
    mkdir /usr/share/nginx/html/_app/immutable && \
    mkdir /usr/share/nginx/html/_app/immutable/assets && \
    mkdir /usr/share/nginx/html/_app/immutable/chunks && \
    mkdir /usr/share/nginx/html/_app/immutable/entry && \
    mkdir /usr/share/nginx/html/_app/immutable/nodes && \
    mkdir /usr/share/nginx/html/Movies && \
    mkdir /usr/share/nginx/html/TVShows

COPY build/*.html /usr/share/nginx/html/
COPY build/*.png /usr/share/nginx/html/
COPY build/*.txt /usr/share/nginx/html/
COPY build/_app/env.js /usr/share/nginx/html/_app/
COPY build/_app/version.json /usr/share/nginx/html/_app/
COPY build/_app/immutable/assets/*.css /usr/share/nginx/html/_app/immutable/assets/
COPY build/_app/immutable/assets/*.woff /usr/share/nginx/html/_app/immutable/assets/
COPY build/_app/immutable/assets/*.woff2 /usr/share/nginx/html/_app/immutable/assets/
COPY build/_app/immutable/assets/*.jpg /usr/share/nginx/html/_app/immutable/assets/
COPY build/_app/immutable/assets/*.webp /usr/share/nginx/html/_app/immutable/assets/
COPY build/_app/immutable/assets/*.avif /usr/share/nginx/html/_app/immutable/assets/
COPY build/_app/immutable/chunks/*.js /usr/share/nginx/html/_app/immutable/chunks/
COPY build/_app/immutable/entry/*.js /usr/share/nginx/html/_app/immutable/entry/
COPY build/_app/immutable/nodes/*.js /usr/share/nginx/html/_app/immutable/nodes/
COPY build/Movies/*.html /usr/share/nginx/html/Movies/
COPY build/TVShows/*.html /usr/share/nginx/html/TVShows/

# EXPOSE 8090