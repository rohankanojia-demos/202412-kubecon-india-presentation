from kubernetes import client, config
import yaml

def readDeploymentYAML(yaml_file):
    with open(yaml_file, 'r') as f:
        return yaml.safe_load(f)

# Configs can be set in Configuration class directly or using helper utility
config.load_kube_config()

deployment = readDeploymentYAML('../../../artifacts/deployment.yaml')
appsV1Api = client.AppsV1Api()
appsV1Api.create_namespaced_deployment("default", deployment)