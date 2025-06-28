const k8s = require('@kubernetes/client-node');
const fs = require('fs');

const kc = new k8s.KubeConfig();
kc.loadFromDefault();

// Get namespace from current context
let namespace = 'default'; // fallback
const currentContext = kc.getCurrentContext();
const contextObj = kc.contexts.find(ctx => ctx.name === currentContext);
if (contextObj && contextObj.namespace) {
    namespace = contextObj.namespace;
}

const main = async () => {
    try {
        const appsV1Api = kc.makeApiClient(k8s.AppsV1Api);

        const yamlContent = fs.readFileSync('../../artifacts/deployment.yaml', 'utf8');
        const resources = k8s.loadYaml(yamlContent);

        await appsV1Api.createNamespacedDeployment(namespace, resources);

        console.log(`Deployment created successfully in namespace "${namespace}"`);
    } catch (err) {
        console.error(err.body || err);
    }
};

main();

