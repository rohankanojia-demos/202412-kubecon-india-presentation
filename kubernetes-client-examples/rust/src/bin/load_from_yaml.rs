use kube::Client;
use kube::Api;
use k8s_openapi::api::apps::v1::Deployment;
use kube::api::PostParams;
use std::error::Error;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    // Initialize the Kubernetes client
    let client = Client::try_default().await?;

    // Define the namespace and API object for Deployments
    let namespace = "default";
    let deployments: Api<Deployment> = Api::namespaced(client.clone(), namespace);

    // Load the YAML file into a serde_yaml::Value object
    let yaml_file = "../../artifacts/deployment.yaml";
    let yaml_content = std::fs::read_to_string(yaml_file)?;

    // Parse the YAML content into a Kubernetes Deployment resource
    let deployment: Deployment = serde_yaml::from_str(&yaml_content)?;
    let post_params = PostParams::default();
    deployments.create(&post_params, &deployment).await?;
    println!("Deployment created successfully.");

    Ok(())
}