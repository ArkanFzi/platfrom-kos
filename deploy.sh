#!/bin/bash

# ==========================================
# Production Deployment Script
# ==========================================

set -e  # Exit on error

echo "======================================"
echo "  Production Deployment"
echo "======================================"
echo ""

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Check if .env.production exists
if [ ! -f ".env.production" ]; then
    echo -e "${RED}Error: .env.production file not found!${NC}"
    echo "Please create .env.production with your production configuration"
    exit 1
fi

echo -e "${GREEN}[1/7]${NC} Loading environment variables..."
source .env.production

echo -e "${GREEN}[2/7]${NC} Pulling latest code..."
git pull origin main

echo -e "${GREEN}[3/7]${NC} Stopping existing containers..."
docker-compose down

echo -e "${GREEN}[4/7]${NC} Building Docker images..."
docker-compose build --no-cache

echo -e "${GREEN}[5/7]${NC} Starting services..."
docker-compose up -d

echo -e "${GREEN}[6/7]${NC} Waiting for services to be healthy..."
sleep 10

# Check if backend is healthy
echo "Checking backend health..."
for i in {1..30}; do
    if curl -f http://localhost:8080/api/health > /dev/null 2>&1; then
        echo -e "${GREEN}✓ Backend is healthy${NC}"
        break
    fi
    if [ $i -eq 30 ]; then
        echo -e "${RED}✗ Backend health check failed${NC}"
        docker-compose logs backend
        exit 1
    fi
    sleep 2
done

echo -e "${GREEN}[7/7]${NC} Running database migrations..."
docker-compose exec -T backend sh -c "cd /app && ./run-migration.sh" || true

echo ""
echo -e "${GREEN}======================================"
echo "  Deployment Complete!"
echo "======================================${NC}"
echo ""
echo "Services running:"
echo "  - Frontend: http://localhost:3000"
echo "  - Backend:  http://localhost:8080"
echo "  - Nginx:    http://localhost:80"
echo ""
echo "View logs:"
echo "  docker-compose logs -f [service_name]"
echo ""
echo "Stop services:"
echo "  docker-compose down"
