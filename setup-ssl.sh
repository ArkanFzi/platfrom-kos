#!/bin/bash

# ==========================================
# SSL Certificate Setup using Certbot
# ==========================================

echo "======================================"
echo "  SSL Certificate Setup"
echo "======================================"
echo ""

# Check if running as root
if [ "$EUID" -ne 0 ]; then 
    echo "Please run as root (sudo)"
    exit 1
fi

# Variables
DOMAIN="$1"
EMAIL="$2"

if [ -z "$DOMAIN" ] || [ -z "$EMAIL" ]; then
    echo "Usage: sudo ./setup-ssl.sh yourdomain.com your@email.com"
    exit 1
fi

echo "Domain: $DOMAIN"
echo "Email: $EMAIL"
echo ""

# Install certbot
echo "[1/4] Installing Certbot..."
if command -v apt-get &> /dev/null; then
    apt-get update
    apt-get install -y certbot python3-certbot-nginx
elif command -v yum &> /dev/null; then
    yum install -y certbot python3-certbot-nginx
else
    echo "Unsupported package manager. Please install certbot manually."
    exit 1
fi

# Stop nginx if running
echo "[2/4] Stopping nginx..."
docker-compose stop nginx || systemctl stop nginx || true

# Get certificate
echo "[3/4] Obtaining SSL certificate..."
certbot certonly --standalone \
    -d $DOMAIN \
    -d www.$DOMAIN \
    --email $EMAIL \
    --agree-tos \
    --no-eff-email \
    --non-interactive

# Copy certificates to nginx directory
echo "[4/4] Copying certificates..."
mkdir -p nginx/ssl
cp /etc/letsencrypt/live/$DOMAIN/fullchain.pem nginx/ssl/
cp /etc/letsencrypt/live/$DOMAIN/privkey.pem nginx/ssl/
chmod 644 nginx/ssl/*.pem

echo ""
echo "======================================"
echo "  SSL Setup Complete!"
echo "======================================"
echo ""
echo "Certificates located at: nginx/ssl/"
echo "Next steps:"
echo "1. Update nginx/conf.d/koskosan.conf with your domain"
echo "2. Restart nginx: docker-compose restart nginx"
echo ""
echo "Certificate will auto-renew via certbot"
echo "Test renewal: certbot renew --dry-run"
