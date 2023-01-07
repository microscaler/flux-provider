# Terraform Provider flux
A provider to install flux to a cluster on completion of cluster creation by terraform. Once the install is completed, the flux controllers self manage their versions using native kubernetes resources. 

This folder encapsulates the Terraform Provider that issues API calls to gitops filesystem CMDB.

## Disclaimer

See license for details.

This project is still in early stage development and is not yet ready for production use.

## Running the example

To run the Terraform Provider locally there are a few steps to complete:

Step 1: Build the source code locally

```
go build -o terraform-provider-flux_v0.1.0
```

Step 2: Move the executable into the local terraform plugin folder:

```
mv terraform-provider-flux_v0.1.0 ~/terraform.d/plugins/$(uname)/
```

> Note: The plugin folder may need to be created.

Step 3: From within this directory, initialize Terraform:

```
terraform init
```

Step 4: Run an apply via Terraform:

```
terraform apply
```

Step 5: coming soon
