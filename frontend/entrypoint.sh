#!/bin/sh

echo "Replacing environment variables in set-env.js with envsubst..."
envsubst < /usr/share/nginx/html/assets/set-env.js > /usr/share/nginx/html/assets/set-env.tmp.js && \
mv /usr/share/nginx/html/assets/set-env.tmp.js /usr/share/nginx/html/assets/set-env.js

# Start Nginx
nginx -g 'daemon off;'