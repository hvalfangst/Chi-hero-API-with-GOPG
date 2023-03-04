#!/bin/sh

  echo "Build image [rollespill] from Dockerfile"
  if ! docker build -t hardokkerdocker/hvalfangst:rollespill .; then
      echo
      echo "[Error building image 'rollespill' - Exiting script]"
      exit 1
  fi

echo -e "\n\n"

# Create our k8s resources
kubectl apply -f k8s/manifest.yaml > /dev/null 2>&1

wait

# Fetch IP associated with entrypoint load balancer
cluster_ip=$(kubectl get svc entrypoint -n default -o jsonpath='{.spec.clusterIP}')

# Get the output of the "kubectl describe svc entrypoint" command
service_definition=$(kubectl describe svc entrypoint)

# Extract the first occurrence of line associated with port definition for the db deployment
db_port_line=$(echo "$service_definition" | grep "db" | head -n 1)

# Extract the port number for db
db_port=$(echo "$db_port_line" | awk '{print $3}' | cut -d "/" -f1)

# Set fully qualified DB URL
db_url="postgresql://postgres:admin@$cluster_ip:$db_port/postgres?sslmode=disable"

# Base64 encode URLs
b64_db_url=$(echo -n "$db_url" | base64 | tr -d '\n')

# Use 'sed' to overwrite value of the field "database-url" contained in yaml file "secrets"
sed -i "s|^\(.*database-url: \)\(.*\)|\1$b64_db_url|" k8s/secrets.yaml

# Create a k8s secret based on contents of manifest file "secrets.yaml"
kubectl apply -f k8s/secrets.yaml > /dev/null 2>&1

# Reset deployment so that the updated secret is being used
kubectl scale deployment api --replicas=0 && kubectl scale deployment api --replicas=1  > /dev/null 2>&1

echo "Preparing pods..."
./misc/progress_bar.sh 0.125

echo -e "\n\n"

# List pods
kubectl get pods

