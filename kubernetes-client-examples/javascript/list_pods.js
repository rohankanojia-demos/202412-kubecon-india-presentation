const k8s = require('@kubernetes/client-node');

const kc = new k8s.KubeConfig();
kc.loadFromDefault();

const k8sApi = kc.makeApiClient(k8s.CoreV1Api);

const main = async () => {
    try {
        const podsResponse = await k8sApi.listNamespacedPod('default');
        podsResponse.body.items.forEach(pod => {
            console.log('Pod Name: ' + pod.metadata.name);
        })
    } catch (err) {
        console.error(err);
    }
};

main();
