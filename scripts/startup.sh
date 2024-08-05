#!/usr/bin/env bash

yell() { echo "$0: $*" >&2; }
die() { yell "$*"; exit 1; }
try() { "$@" || die "cannot $*"; }

## Esses serviços estão no automático

# sudo systemctl enable transmission-daemon


# try sudo minidlnad &
# try sudo transmission-daemon &
try /home/droidian/scripts/gerencia-atualizacoes-ip-dynv6.sh &
try /home/droidian/go/bin/go run /home/droidian/scripts/fileserver.go &
try /home/droidian/go/bin/go run /home/droidian/p/torrent-downloader/main.go &
try sudo /home/droidian/bin/jellyfin/jellyfin &
