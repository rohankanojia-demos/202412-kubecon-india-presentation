from kubernetes import client, config
from kubernetes.config.kube_config import KUBE_CONFIG_DEFAULT_LOCATION, list_kube_config_contexts

# Load kubeconfig
config.load_kube_config()

# Get current context and namespace
contexts, current_context = list_kube_config_contexts(config_file=KUBE_CONFIG_DEFAULT_LOCATION)
namespace = current_context.get('context', {}).get('namespace', 'default')

print(f"Using namespace: {namespace}")
print("Creating Deployment:")

deployment = client.V1Deployment(
    api_version="apps/v1",
    kind="Deployment",
    metadata=client.V1ObjectMeta(name="client-python-deploy"),
    spec=client.V1DeploymentSpec(
        replicas=1,
        selector=client.V1LabelSelector(
            match_labels={"app": "client-python-demo-app"}
        ),
        template=client.V1PodTemplateSpec(
            metadata=client.V1ObjectMeta(labels={"app": "client-python-demo-app"}),
            spec=client.V1PodSpec(
                containers=[
                    client.V1Container(
                        name="client-python-demo-container",
                        image="busybox:latest",
                        ports=[client.V1ContainerPort(container_port=8080)]
                    )
                ]
            )
        )
    )
)

appsV1 = client.AppsV1Api()

# Create the Deployment
try:
    appsV1.create_namespaced_deployment(namespace=namespace, body=deployment)
    print("Deployment 'client-python-deploy' created successfully.")
except client.exceptions.ApiException as e:
    print(f"Exception when calling AppsV1Api->create_namespaced_deployment: {e}")