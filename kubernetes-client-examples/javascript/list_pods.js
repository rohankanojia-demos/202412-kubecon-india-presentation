const k8s = require('@kubernetes/client-node');

const kc = new k8s.KubeConfig();
kc.loadFromDefault();

// Get namespace from current context
let namespace = 'default'; // fallback
const currentContext = kc.getCurrentContext();
const contextObj = kc.contexts.find(ctx => ctx.name === currentContext);
if (contextObj && contextObj.namespace) {
    namespace = contextObj.namespace;
}

const k8sApi = kc.makeApiClient(k8s.CoreV1Api);

const main = async () => {
    try {
        const podsResponse = await k8sApi.listNamespacedPod(namespace);
        podsResponse.body.items.forEach(pod => {
            console.log('Pod Name: ' + pod.metadata.name);
        });
    } catch (err) {
        console.error(err.body || err);
    }
};

main();

