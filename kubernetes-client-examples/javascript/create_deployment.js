const k8s = require('@kubernetes/client-node');

const kc = new k8s.KubeConfig();
kc.loadFromDefault();

const deploymentName = 'client-javascript-deployment';

const deploymentSpec = {
  metadata: {
    name: deploymentName,
    labels: {
      app: 'client-javascript-app'
    }
  },
  spec: {
    replicas: 2,
    selector: {
      matchLabels: {
        app: 'client-javascript-app'
      }
    },
    template: {
      metadata: {
        labels: {
          app: 'client-javascript-app'
        }
      },
      spec: {
        containers: [
          {
            name: 'nginx-container',
            image: 'nginx:latest',
            ports: [
              {
                containerPort: 80
              }
            ]
          }
        ]
      }
    }
  }
};

const k8sApi = kc.makeApiClient(k8s.AppsV1Api);

const main = async () => {
    try {
        const createDeploymentResponse = await k8sApi.createNamespacedDeployment('default', deploymentSpec);
        console.log("Deployment %s created", createDeploymentResponse.body.metadata.name)
    } catch (err) {
        console.error(err);
    }
};

main();
