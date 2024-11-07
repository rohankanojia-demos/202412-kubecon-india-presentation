const k8s = require('@kubernetes/client-node');
const fs = require('fs');

const kc = new k8s.KubeConfig();
kc.loadFromDefault();

const k8sApi = kc.makeApiClient(k8s.AppsV1Api);

const main = async () => {
    try {
        const filePath = "../../artifacts/deployment.yaml"
        const fileContents = fs.readFileSync(filePath, 'utf8');

        const resources = k8s.loadYaml(fileContents);

        await k8sApi.createNamespacedDeployment("default", resources);
        console.log("Deployment created successfully");
    } catch (err) {
        console.error(err);
    }
};

main();
