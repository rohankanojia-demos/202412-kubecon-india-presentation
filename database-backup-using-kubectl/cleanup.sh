#!/bin/bash
# For testing purposes let's delete file older than 1 minute
find $BACKUP_DIR -type f -mmin +1 -delete
