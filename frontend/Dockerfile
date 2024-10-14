# Stage 1: Build
FROM node:20.17-alpine3.19 AS builder

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json to install dependencies
COPY package*.json .

# Install dependencies
RUN npm install

# Copy the rest of the code
COPY . .

# # Make the entrypoint script executable
# RUN chmod +x /app/entrypoint.sh

# # Replace environment variables in the Angular environment.ts file before building
# RUN /app/entrypoint.sh

# Build the Angular app for production
RUN npm run build --c=production

# Stage 2: Serve the Angular app using NGINX
FROM nginx:alpine3.19 AS prod

# Copy the build output from the build stage to the Nginx html directory
COPY --from=builder /app/dist/frontend/browser /usr/share/nginx/html

# Copy the set-env.js script
COPY src/assets/set-env.js /usr/share/nginx/html/assets/set-env.js

# Copy the entrypoint script into the container
COPY entrypoint.sh /entrypoint.sh

# Make the entrypoint script executable
RUN chmod +x /entrypoint.sh

# Expose port 80
EXPOSE 80
EXPOSE 443

# # Start Nginx server
# CMD ["nginx", "-g", "daemon off;"]

# Use the custom entrypoint script
ENTRYPOINT ["/entrypoint.sh"]
