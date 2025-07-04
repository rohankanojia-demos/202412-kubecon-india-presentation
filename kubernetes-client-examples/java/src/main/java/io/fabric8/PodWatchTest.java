package io.fabric8;

import io.fabric8.kubernetes.api.model.Pod;
import io.fabric8.kubernetes.client.KubernetesClient;
import io.fabric8.kubernetes.client.KubernetesClientBuilder;
import io.fabric8.kubernetes.client.Watch;
import io.fabric8.kubernetes.client.Watcher;
import io.fabric8.kubernetes.client.WatcherException;

public class PodWatchTest {
    public static void main(String[] args) {
        try (KubernetesClient client = new KubernetesClientBuilder().build()) {
            Watch watch = client.pods().watch(new Watcher<Pod>() {
                @Override
                public void eventReceived(Action action, Pod pod) {
                    System.out.printf("%s : %s\n", action.name(), pod.getMetadata().getName());
                }

                @Override
                public void onClose(WatcherException e) {
                    System.out.printf("onClose : %s\n", e.getMessage());
                }

            });

            System.out.println("Watch open for 10 seconds");
            // Watch till 10 seconds
            Thread.sleep(10 * 1000);

            // Close Watch
            watch.close();
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            e.printStackTrace();
        }
    }
}
