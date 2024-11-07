from kubernetes import client, config

# Configs can be set in Configuration class directly or using helper utility
config.load_kube_config()

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
    appsV1.create_namespaced_deployment(namespace="default", body=deployment)
    print("Deployment client-python-deploy created successfully.")
except client.exceptions.ApiException as e:
    print(f"Exception when calling AppsV1Api->create_namespaced_deployment: {e}")