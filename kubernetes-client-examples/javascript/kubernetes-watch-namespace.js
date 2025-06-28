const k8s = require('@kubernetes/client-node');

const kubeconfig = new k8s.KubeConfig();
kubeconfig.loadFromDefault();

// Get namespace from current context
let namespace = 'default';
const currentContext = kubeconfig.getCurrentContext();
const contextObj = kubeconfig.contexts.find(ctx => ctx.name === currentContext);
if (contextObj && contextObj.namespace) {
    namespace = contextObj.namespace;
}

const main = async () => {
    try {
        const watch = new k8s.Watch(kubeconfig);

        const req = await watch.watch(
            `/api/v1/namespaces/${namespace}/pods`,
            {},
            (type, pod) => {
                console.log(`Event ${type} ${pod.metadata.name}`);
            },
            (err) => {
                if (err) {
                    console.error('Error watching Pods:', err);
                }
            }
        );

        await sleep(10000);

        // Abort the watch
        req.abort();
    } catch (err) {
        console.error(err);
    }
};

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms || 10000));
}

main();
