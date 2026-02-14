#!/bin/bash

# Exit on any error
set -e

echo "🚀 Starting full deployment sequence..."

# 1. Build Backend
echo "📦 Building Backend..."
cd backend/shop
make build
cd ../..

# 2. Build Frontend Admin
echo "📦 Building Frontend Admin..."
cd frontend/shop-admin
yarn build
cd ../..

# 3. Build Mobile User Web
echo "📦 Building Mobile User Web..."
cd mobile/shop/user
yarn build:web:production
cd ../../../..

# 4. Build Mobile Keeper Web
echo "📦 Building Mobile Keeper Web..."
cd mobile/shop/keeper
yarn build:web:production
cd ../../../..

# 5. Start Docker Services
echo "🐳 Starting Docker services..."
cd opt/projects/shop
./start.sh
cd ../../..

echo "✅ Deployment completed successfully!"
