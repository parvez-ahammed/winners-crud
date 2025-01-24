# Stage 1: Build Stage for Dependencies
FROM node:22-alpine AS builder

# Set working directory to /app
WORKDIR /app

# Copy package files from client
COPY client/package.json /app/client/

# Install dependencies using npm workspaces
WORKDIR /app/client
RUN npm install

# Stage 2: Build Client
# Assumes client dirs contains the .env file
COPY client/ /app/client/

# # Build shared libraries
RUN npm run build

# Stage 3: Runtime Stage
CMD ["npm", "run", "start"]
