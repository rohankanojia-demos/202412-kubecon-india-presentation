package io.official;

import com.google.gson.reflect.TypeToken;
import io.kubernetes.client.openapi.ApiClient;
import io.kubernetes.client.openapi.ApiException;
import io.kubernetes.client.openapi.Configuration;
import io.kubernetes.client.openapi.apis.CoreV1Api;
import io.kubernetes.client.openapi.models.V1Pod;
import io.kubernetes.client.util.Config;
import io.kubernetes.client.util.Watch;

import java.io.IOException;

public class PodWatchTest {
    public static void main(String[] args) throws IOException, ApiException {
        ApiClient client = Config.defaultClient();
        Configuration.setDefaultApiClient(client);
        CoreV1Api api = new CoreV1Api();
        Watch<V1Pod> watch =
                Watch.createWatch(
                        client,
                        api.listPodForAllNamespaces()
                          .watch(true)
                          .buildCall(null),
                        new TypeToken<Watch.Response<V1Pod>>() {}.getType());
        try {
            for (Watch.Response<V1Pod> item : watch) {
                System.out.printf("%s : %s%n", item.type, item.object.getMetadata().getName());
            }
        } finally {
            watch.close();
        }
    }
}
