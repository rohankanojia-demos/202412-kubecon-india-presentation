from kubernetes import client, config
from kubernetes.config.kube_config import list_kube_config_contexts, KUBE_CONFIG_DEFAULT_LOCATION

# Load kubeconfig
config.load_kube_config()

# Get namespace from current context
contexts, current_context = list_kube_config_contexts(config_file=KUBE_CONFIG_DEFAULT_LOCATION)
namespace = current_context.get('context', {}).get('namespace', 'default')

# Create CoreV1 API client
v1 = client.CoreV1Api()

# List pods in the specified namespace
pods_list = v1.list_namespaced_pod(namespace=namespace)
for pod in pods_list.items:
    print("%s\t%s\t%s" % (pod.status.pod_ip, pod.metadata.namespace, pod.metadata.name))
