const k8s = require('@kubernetes/client-node');
const fs = require('fs');

const kc = new k8s.KubeConfig();
kc.loadFromDefault();



const main = async () => {
    try {

        const appsV1Api = kc.makeApiClient(k8s.AppsV1Api);
        const resources = k8s.loadYaml(fs.readFileSync("../../artifacts/deployment.yaml", 'utf8'));
        await appsV1Api.createNamespacedDeployment("default", resources);

        console.log("Deployment created successfully");
    } catch (err) {
        console.error(err);
    }
};

main();
