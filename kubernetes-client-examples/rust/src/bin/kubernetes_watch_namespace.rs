use futures::TryStreamExt;
use k8s_openapi::api::core::v1::Pod;
use kube::{
    api::Api,
    client::Client,
    config::Config,
    runtime::{watcher, WatchStreamExt},
};
use kube::ResourceExt;
use std::error::Error;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    // Load kubeconfig and get namespace
    let config = Config::infer().await?;
    let namespace = config.default_namespace.to_string();

    let client = Client::try_from(config)?;

    let pods: Api<Pod> = Api::namespaced(client.clone(), &namespace);

    watcher(pods, watcher::Config::default())
        .applied_objects()
        .default_backoff()
        .try_for_each(|p| async move {
            println!("Event {}", p.name_any());
            Ok(())
        })
        .await?;

    Ok(())
}
