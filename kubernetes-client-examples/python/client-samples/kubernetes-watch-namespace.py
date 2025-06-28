from kubernetes import client, config, watch

# Load the kubeconfig file
config.load_kube_config()

# Get the current context and namespace
contexts, active_context = config.list_kube_config_contexts()
namespace = active_context.get('context', {}).get('namespace', 'default')  # Default to 'default' namespace

# Initialize the API client
v1 = client.CoreV1Api()

# Set up the watch
w = watch.Watch()
count = 30

# Stream pod events only in the current context's namespace
for event in w.stream(v1.list_namespaced_pod, namespace=namespace, timeout_seconds=10):
    print("Event: %s %s" % (event['type'], event['object'].metadata.name))
    count -= 1
    if count == 0:
        w.stop()

print("Ended.")
