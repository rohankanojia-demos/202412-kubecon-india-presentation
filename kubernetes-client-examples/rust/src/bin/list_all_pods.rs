use kube::Client;
use kube::Api;
use kube::api::ListParams;
use k8s_openapi::api::core::v1::Pod;
use std::error::Error;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    let client = Client::try_default().await?;
    let pods: Api<Pod> = Api::namespaced(client, "default");
    let pods_list = pods.list(&ListParams::default()).await?;

    for pod in pods_list.items {
        println!("Pod Name: {}", pod.metadata.name.unwrap_or_default());
    }

    Ok(())
}

