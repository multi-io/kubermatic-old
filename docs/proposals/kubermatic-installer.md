# Kubermatic Installer

**Author**: Patrick Kavajin

**Status**: designing

*short description of the topic e.g.*
kubermatic-installer should be a go application which allows customers to easily setup new seed/master clusters.

## Motivation and Background

Our current approach to installing kubermatic involves a lot of manual and errorprone manual steps:
* Manually provision the nodes  
* Setting up the values.yaml  
* Setting up the datacenters.yaml  
* Helm deploying (order is important as is also which components need to be deployed on which cluster type(seed/master), ...)  

The idea is to make the installation way more simple by offering the customer a graphical editor to setup values.yaml/datacenters.yaml configuration and abstract away the provisioning and manual helm deployment.

## Implementation

The kubermatic installer is one application consisting of these major components:

* Wizard  
* Installer  
	* Node Provisioner (should use the machine controller for provisioning)  
	* HA-Kubernetes installer (aka former seed installer)  
	* Helm based installer for e.g. seed/master charts.  
* Manifest  

The **wizard** is a fancy HTML5 based wizard containing:  
* A graphical editor used for configuring a manifest file.  
* A summary page with a "install now" button which executed the installer component with the generated manifest file passed and a "export" button for downloading the generated manifest for offline installations.  
* A installation progress pane (as in: installer runs, reports this and that output, etc.)  

The wizard will be compiled into the kubermatic installer application which then can be used as a webserver serving the wizard.

The kubermatic installer is supposed to be hosted (in a later step) online on install.kubermatic.io and also be available as a download for offline installations.

The **manifest** is a yaml file containing:
* The generated values.yaml
* The generated datacenters.yaml
* A flag which cloud provider will be used.
* The wanted network configuration
* A list of provisioned nodes
	* IPs, Users, SSH Keys, ...
* A list of to-be-provisioned nodes
	* Which size,
	* Cloudprovider credentials
	* SSH keys to deploy
	* OS,
	* ...

The **installer** contains all the magic to bootstrap nodes, kubernetes clusters and kubermatic installations.

**Node provisioner**

If the manifest contains nodes which should be provisioned, the installer should use the kubermatic machine controller to provision the corresponding machines. The userdata/cloudinit script should only deploy the needed ssh keys to the machine. No further installation should happen there.

**HA-Kubernetes installer**

The installer should connect (as in: using ssh) to the provisioned nodes and install a HA Kubernetes cluster onto it.
It will do the same as the former seed installer.

**Helm based installer**

After nodes have been provisioned and a kubernetes cluster was set up, the kubermatic helm charts should be installed onto that cluster.
Therefore the installer should also contain the most recent (as in: time of building the installer) kubermatic charts baked into the installer binary. Optionally it should be possible to manually specify a path to the helm charts in case a custom should be used.

## MVP

* Graphical Editor can set up basic settings
* Installer can install kubernetes on already provisioned machines
* Installer can install kubermatic charts from a specified fs path