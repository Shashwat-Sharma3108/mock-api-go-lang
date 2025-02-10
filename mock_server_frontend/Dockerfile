# Dockerfile for React app using Node 18
FROM node:18

WORKDIR /app

# Copy package.json files first and install dependencies
COPY package*.json ./

# Install esbuild explicitly
RUN npm install esbuild@latest

# Force the rebuild of esbuild for the correct platform
RUN npm install

# Copy the rest of the application code
COPY . .

# Start the app
CMD ["npm", "run", "dev"]