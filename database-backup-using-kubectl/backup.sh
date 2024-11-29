#!/bin/bash
PGPASSWORD=$POSTGRES_PASSWORD pg_dump \
  -U $POSTGRES_USERNAME \
  -h $DB_HOST -d \
  $DB_NAME > $BACKUP_DIR/$DB_NAME-$(date +%Y-%m-%d_%H-%M-%S).sql