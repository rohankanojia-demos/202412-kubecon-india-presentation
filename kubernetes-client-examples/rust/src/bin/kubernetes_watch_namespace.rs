use futures::TryStreamExt;
use k8s_openapi::api::core::v1::Pod;
use kube::{
    api::Api,
    client::Client,
    runtime::{watcher, WatchStreamExt},
};
use kube::ResourceExt;
use std::error::Error;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>>  {
    let client = Client::try_default().await?;
    
let pods: Api<Pod> = Api::default_namespaced(client.clone());
watcher(pods, watcher::Config::default()).applied_objects().default_backoff()
  .try_for_each(|p| async move {
      println!("Event {}", p.name_any());
      Ok(())
   })
  .await?;
    
    
    Ok(())
}