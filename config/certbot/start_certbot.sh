#!/bin/sh

trap 'exit 0' TERM

export PATH="/usr/local/bin:/usr/local/sbin:/usr/sbin:/usr/bin:/sbin:/bin"

while :; do
  cert_path="/etc/letsencrypt/live/ssoauth.online/fullchain.pem"

  if [ ! -f $cert_path ]; then
      echo "Certificate not found. Requesting new certificate."
      certbot certonly --webroot --webroot-path=/var/www/certbot --email ant.goncharik.development@gmail.com --agree-tos --no-eff-email -d ssoauth.online
  else
      echo "Certificate found. Checking expiry date."
      expiry_date=$(openssl x509 -enddate -noout -in $cert_path | cut -d= -f2)
      expiry_seconds=$(date -d "$expiry_date" +%s)
      current_seconds=$(date +%s)
      diff_seconds=$((expiry_seconds - current_seconds))

      # If certificate expires in less than 30 days (2592000 seconds), renew it
      if [ $diff_seconds -le 2592000 ]; then
          echo "Certificate is expiring in less than 30 days. Renewing."
          certbot renew --deploy-hook "/usr/sbin/nginx -s reload" --disable-hook-validation
      else
          echo "Certificate is valid for more than 30 days. No renewal needed."
      fi
  fi

  sleep 24h & wait $!
done
