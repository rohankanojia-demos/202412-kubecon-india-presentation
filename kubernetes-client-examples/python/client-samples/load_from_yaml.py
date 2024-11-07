import yaml

yaml_file = 'k8s/deployment.yaml'
with open(yaml_file, 'r') as f:
    try:
        yaml_content = yaml.safe_load(f)
    except yaml.YAMLError as exc:
        print(f"Error reading YAML file: {exc}")

print(f"Deployment {yaml_content} loaded successfully")