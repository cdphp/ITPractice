#!/bin/bash
npm run build
scp -r ./dist/* root@120.77.153.236:/var/www/html
scp -r ./static root@120.77.153.236:/var/www/html
