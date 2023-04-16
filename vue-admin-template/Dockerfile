FROM node:14.9.0 as base

WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the source code
COPY . .

# Expose the port
EXPOSE 9528

# Start the development server
#CMD ["npm", "run", "dev"]
