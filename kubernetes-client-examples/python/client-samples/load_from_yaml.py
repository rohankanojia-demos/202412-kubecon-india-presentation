from kubernetes import client, config
from kubernetes.config.kube_config import list_kube_config_contexts, KUBE_CONFIG_DEFAULT_LOCATION
import yaml

def readDeploymentYAML(yaml_file):
    with open(yaml_file, 'r') as f:
        return yaml.safe_load(f)

# Load kubeconfig
config.load_kube_config()

# Get namespace from current context
contexts, current_context = list_kube_config_contexts(config_file=KUBE_CONFIG_DEFAULT_LOCATION)
namespace = current_context.get('context', {}).get('namespace', 'default')

# Load the deployment YAML
deployment = readDeploymentYAML('../../../artifacts/deployment.yaml')

# Create the deployment in the current context's namespace
appsV1Api = client.AppsV1Api()
try:
    appsV1Api.create_namespaced_deployment(namespace=namespace, body=deployment)
    print(f"Deployment created successfully in namespace '{namespace}'")
except client.exceptions.ApiException as e:
    print(f"Exception when creating deployment: {e}")
