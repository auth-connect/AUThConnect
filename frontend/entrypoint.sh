#!/bin/sh
set -e
set -x

echo "Replacing environment variables in set-env.js with envsubst..."
echo "Current value of API_URL: ${API_URL}"

# Make sure set-env.js exists before trying to replace variables
if [ -f /usr/share/nginx/html/assets/set-env.js ]; then
    cat /usr/share/nginx/html/assets/set-env.js  # Print original file
    envsubst < /usr/share/nginx/html/assets/set-env.js > /usr/share/nginx/html/assets/set-env.tmp.js
    mv /usr/share/nginx/html/assets/set-env.tmp.js /usr/share/nginx/html/assets/set-env.js
    echo "Updated set-env.js:"
    cat /usr/share/nginx/html/assets/set-env.js  # Print updated file
else
    echo "Error: /usr/share/nginx/html/assets/set-env.js does not exist"
    exit 1
fi

# Start Nginx
nginx -g 'daemon off;'
