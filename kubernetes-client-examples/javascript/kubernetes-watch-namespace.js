const k8s = require('@kubernetes/client-node');

const kubeconfig = new k8s.KubeConfig();
kubeconfig.loadFromDefault();

const k8sApi = kubeconfig.makeApiClient(k8s.CoreV1Api);

const main = async () => {
    try {
        const watch = new k8s.Watch(kubeconfig);
        req = await watch.watch(
            `/api/v1/namespaces/default/pods`,
            {},
            (type, pod) => {
              console.log(`Event Type: ${type}`);
              console.log(`Pod Name: ${pod.metadata.name}`);
              console.log('---');
            },
            (err) => {
              if (err) {
                console.error('Error watching Pods:', err);
              }
            }
          );
        await sleep(10000);

        // watch returns a request object which you can use to abort the watch.
        req.abort();
    } catch (err) {
        console.error(err);
    }
};


function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms || DEF_DELAY));
}

main();