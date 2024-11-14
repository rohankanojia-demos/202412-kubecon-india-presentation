package io.fabric8;

import io.fabric8.kubernetes.api.model.apps.Deployment;
import io.fabric8.kubernetes.client.KubernetesClient;
import io.fabric8.kubernetes.client.KubernetesClientBuilder;

public class LoadAndCreateDeployment {
    public static void main(String[] args) {
        try (KubernetesClient client = new KubernetesClientBuilder().build()) {
            Deployment deployment = client.apps().deployments()
                .load(LoadAndCreateDeployment.class.getResourceAsStream("/deployment.yaml"))
                .item();

            deployment = client.apps().deployments().inNamespace("default").resource(deployment).create();
            System.out.println("Deployment " + deployment.getMetadata().getName() + " created");
        }
    }
}
