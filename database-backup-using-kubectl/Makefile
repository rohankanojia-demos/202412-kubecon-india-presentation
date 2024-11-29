# Define variables
IMAGE_REGISTRY=ttl.sh
IMAGE_NAME=postgres-backup
IMAGE_TAG=1h
BACKUP_SCRIPT=backup.sh
CRONJOB_FILE=postgres-backup.yml

.PHONY: build push

build:
	@echo "Building container image"
	podman --remote build -t $(IMAGE_REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG) .
	podman push $(IMAGE_REGISTRY)/$(IMAGE_NAME):$(IMAGE_TAG)

# Kubernetes apply
.PHONY: apply-cronjob

apply-cronjob:
	@echo "Applying CronJob"
	kubectl apply -f $(CRONJOB_FILE)

# Full deployment
.PHONY: deploy

deploy: build apply-cronjob
	@echo "Deployment complete."

