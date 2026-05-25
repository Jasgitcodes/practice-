#!/bin/bash

mkdir 0 A
chmod 401 0 A

ln -s 0 3

touch 1 2 4 5 6 7 8 9 

chmod 402 1 9
chmod 604 2 8
chmod 510 4 7
chmod 560 5 6

TZ=utc touch -t 198601050000 0
TZ=utc touch -t 198611130001 1
TZ=utc touch -t 198803050010 2
TZ=utc touch -t 199002160011 3
TZ=utc touch -t 199010070100 4
TZ=utc touch -t 199011070101 5
TZ=utc touch -t 199102080110 6
TZ=utc touch -t 199103080111 7
TZ=utc touch -t 199405201000 8
TZ=utc touch -t 199406101001 9
TZ=utc touch -t 199504101010 A

TZ=utc ls -l --time-style='+%F %R' | sed 1d | awk '{print $1, $6, $7, $8, $9, $10}'

echo "done ✅"