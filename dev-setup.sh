#!/bin/bash

cd client
nvm use
npm install
npm run build
cd ../server
cp -Rpv ../client/dist .

echo "All set up for local dev work!"
