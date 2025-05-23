# Use the latest version of Node image as a build stage
FROM node:latest AS build

# Set you working directory in the build satge
WORKDIR /app

# Copy the package.json and yarn lock first to leverage Docker cache
COPY ./package.json .

# install dependencies
RUN yarn install

# Cpoy the rest of the app code
COPY . .

# Build the project
RUN yarn build

# Use the latest Nginx image as the runtime stage
FROM nginx:latest

# Copy the build files into the Nginx html directory
COPY --from=build /app/dist /usr/share/nginx/html

# Expose the port that Nginx will run on
EXPOSE 80