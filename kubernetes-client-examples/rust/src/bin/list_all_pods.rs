use kube::{Client, Api, api::ListParams, config::Config};
use k8s_openapi::api::core::v1::Pod;
use std::error::Error;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    // Load config and get namespace
    let config = Config::infer().await?;
    let namespace = config.default_namespace.to_string();

    let client = Client::try_from(config)?;
    let pods: Api<Pod> = Api::namespaced(client, &namespace);
    let pods_list = pods.list(&ListParams::default()).await?;

    println!("Listing pods in namespace: {}", namespace);
    for pod in pods_list.items {
        println!("Pod Name: {}", pod.metadata.name.unwrap_or_default());
    }

    Ok(())
}
