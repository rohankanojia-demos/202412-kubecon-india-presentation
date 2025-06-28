const k8s = require('@kubernetes/client-node');

const kc = new k8s.KubeConfig();
kc.loadFromDefault();

// Get namespace from the current context in the kubeconfig
let namespace = 'default'; // fallback
const currentContext = kc.getCurrentContext();
const contextObj = kc.contexts.find(ctx => ctx.name === currentContext);
if (contextObj && contextObj.namespace) {
  namespace = contextObj.namespace;
}

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

const main = async () => {
  try {
    const AppsV1Api = kc.makeApiClient(k8s.AppsV1Api);
    const createDeploymentResponse = await AppsV1Api.createNamespacedDeployment(namespace, deploymentSpec);
    console.log("Deployment %s created in namespace %s", createDeploymentResponse.body.metadata.name, namespace);
  } catch (err) {
    console.error(err.body || err);
  }
};

main();