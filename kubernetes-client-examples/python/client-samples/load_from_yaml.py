from kubernetes import client, config
import yaml

# Configs can be set in Configuration class directly or using helper utility
config.load_kube_config()

yaml_file = '../../../artifacts/deployment.yaml'
with open(yaml_file, 'r') as f:
    try:
        yaml_content = yaml.safe_load(f)
    except yaml.YAMLError as exc:
        print(f"Error reading YAML file: {exc}")

print("Deployment loaded successfully")

appsV1Api = client.AppsV1Api()
appsV1Api.create_namespaced_deployment("default", yaml_content)
print("Deployment created successfully")