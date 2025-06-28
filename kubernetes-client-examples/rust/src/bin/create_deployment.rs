use kube::Api;
use kube::Client;
use kube::config::Config;
use k8s_openapi::api::apps::v1::Deployment;
use kube::api::PostParams;
use k8s_openapi::api::core::v1::{Container, ContainerPort, PodSpec, PodTemplateSpec};
use k8s_openapi::apimachinery::pkg::apis::meta::v1::{ObjectMeta, LabelSelector};
use k8s_openapi::api::apps::v1::DeploymentSpec;

use std::error::Error;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    // Load kubeconfig and get client
    let config = Config::infer().await?;
    let namespace = config.default_namespace.to_string();

    let client = Client::try_from(config)?;

    // Use the namespace from the kubeconfig context dynamically
    let deployments: Api<Deployment> = Api::namespaced(client.clone(), &namespace);

    let deployment_name = "nginx-deployment";
    let deployment = Deployment {
        metadata: ObjectMeta {
            name: Some(deployment_name.to_string()),
            labels: Some(
                vec![("app".to_string(), "nginx".to_string())]
                    .into_iter()
                    .collect(),
            ),
            ..Default::default()
        },
        spec: Some(DeploymentSpec {
            replicas: Some(3),
            selector: LabelSelector {
                match_labels: Some(
                    vec![("app".to_string(), "nginx".to_string())]
                        .into_iter()
                        .collect(),
                ),
                ..Default::default()
            },
            template: PodTemplateSpec {
                metadata: Some(ObjectMeta {
                    labels: Some(
                        vec![("app".to_string(), "nginx".to_string())]
                            .into_iter()
                            .collect(),
                    ),
                    ..Default::default()
                }),
                spec: Some(PodSpec {
                    containers: vec![
                        Container {
                            name: "nginx".to_string(),
                            image: Some("nginx:latest".to_string()),
                            ports: Some(vec![ContainerPort {
                                container_port: 80,
                                ..Default::default()
                            }]),
                            ..Default::default()
                        },
                    ],
                    ..Default::default()
                }),
            },
            ..Default::default()
        }),
        ..Default::default()
    };

    // Apply the deployment (create it)
    let post_params = PostParams::default();
    let result = deployments.create(&post_params, &deployment).await;

    match result {
        Ok(_) => println!("Deployment '{}' created successfully in namespace '{}'", deployment_name, namespace),
        Err(err) => eprintln!("Error creating deployment: {:?}", err),
    }

    Ok(())
}
