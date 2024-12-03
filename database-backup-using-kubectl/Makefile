# Define variables
IMAGE_REGISTRY=ttl.sh
IMAGE_NAME=postgres-backup
IMAGE_TAG=1h
BACKUP_SCRIPT=backup.sh
BACKUP_CRONJOB=postgres-backup.yml
CLEANUP_CRONJOB=cleanup-cronjob.yml

.PHONY: build push

build:
	@echo "Building container image"
	podman build -t $(IMAGE_REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG) .
	podman push $(IMAGE_REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG)

# Kubernetes apply
.PHONY: apply-cronjob

apply-cronjob:
	@echo "Applying CronJob"
	kubectl apply -f $(BACKUP_CRONJOB)
	kubectl apply -f $(CLEANUP_CRONJOB)

# Full deployment
.PHONY: deploy

deploy: build apply-cronjob
	@echo "Deployment complete."

