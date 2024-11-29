from kubernetes import client, config

# Configs can be set in Configuration class directly or using helper utility
config.load_kube_config()

pods = client.CoreV1Api()
podsList = pods.list_namespaced_pod(namespace="default")
for i in podsList.items:
    print("%s\t%s\t%s" % (i.status.pod_ip, i.metadata.namespace, i.metadata.name))