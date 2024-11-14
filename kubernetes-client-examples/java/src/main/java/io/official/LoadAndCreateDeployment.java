package io.official;

import io.kubernetes.client.openapi.ApiClient;
import io.kubernetes.client.openapi.ApiException;
import io.kubernetes.client.openapi.Configuration;
import io.kubernetes.client.openapi.apis.AppsV1Api;
import io.kubernetes.client.openapi.models.V1Deployment;
import io.kubernetes.client.util.Config;
import io.kubernetes.client.util.Yaml;
import org.snakeyaml.engine.v2.api.Load;

import java.io.File;
import java.io.IOException;
import java.net.URISyntaxException;
import java.nio.file.Paths;

public class LoadAndCreateDeployment {
  public static void main(String[] args) throws URISyntaxException {
    try {
      ApiClient client = Config.defaultClient();
      Configuration.setDefaultApiClient(client);
      File deploymentYaml = Paths.get(LoadAndCreateDeployment.class.getResource("/deployment.yaml").toURI()).toFile();
      V1Deployment deployment = (V1Deployment) Yaml.load(deploymentYaml);

      AppsV1Api appsV1Api = new AppsV1Api(client);
      deployment = appsV1Api.createNamespacedDeployment("default", deployment).execute();
      System.out.println(deployment.getMetadata().getName() + " created");
    } catch (IOException | ApiException e) {
      throw new IllegalStateException("failed to create Deployment : " + e.getMessage());
    }
  }
}
