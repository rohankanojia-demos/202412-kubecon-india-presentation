use kube::{Client, Api, config::Config};
use std::fs;
use k8s_openapi::api::apps::v1::Deployment;
use kube::api::PostParams;
use std::error::Error;

async fn read_deployment_yaml(yaml_file: &str) -> Result<Deployment, Box<dyn Error>> {
    // Read the YAML file to a string
    let yaml_content = fs::read_to_string(yaml_file)?;

    // Deserialize the YAML content into a Deployment object
    let deployment: Deployment = serde_yaml::from_str(&yaml_content)?;

    Ok(deployment)
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    // Load kubeconfig and get namespace
    let config = Config::infer().await?;
    let namespace = config.default_namespace().to_string();

    // Initialize the Kubernetes client
    let client = Client::try_from(config)?;

    // Read deployment YAML
    let deployment = read_deployment_yaml("../../artifacts/deployment.yaml").await?;

    let post_params = PostParams::default();
    let deployments: Api<Deployment> = Api::namespaced(client.clone(), &namespace);
    deployments.create(&post_params, &deployment).await?;

    println!("Deployment created successfully in namespace '{}'.", namespace);

    Ok(())
}
