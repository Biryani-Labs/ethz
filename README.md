# EZETH Documentation

## Overview

EZETH is a powerful setup wizard designed to streamline the deployment and management of Ethereum-based applications. Operating over SSH, EZETH enables users to draft and execute blueprints on remote machines effortlessly. It integrates with various system components, including Systemd Dynamic Users and Docker, to ensure seamless execution of commands.

## Key Features

- **Blueprint Drafting**: Easily create and manage deployment blueprints.
- **Blueprint Execution**: Run blueprints on remote machines with a single command.
- **SSH Integration**: Execute commands on remote systems securely.
- **Comprehensive Support**: Works with Systemd Dynamic Users and Docker.

## Commands

### Draft a Blueprint

To create a new blueprint, use the following command:

```
ezeth draft [blueprint-name]
```

**Parameters:**
- `[blueprint-name]`: The name you wish to assign to your new blueprint.

**Description:**
This command generates a blueprint with the specified name. Blueprints define the configuration and commands needed to deploy your Ethereum-based applications.

### Execute a Blueprint

To run an existing blueprint on a remote machine, use the following command:

```
ezeth exec [blueprint-name]
```

**Parameters:**
- `[blueprint-name]`: The name of the blueprint you wish to execute.

**Description:**
This command executes the specified blueprint on a remote machine via SSH. It applies the configurations and commands outlined in the blueprint to set up or manage your Ethereum-based application.

## Setup and Usage

1. **Install EZETH**: Follow the installation instructions provided for your operating system to set up EZETH on your local machine.
   
2. **Configure SSH Access**: Ensure you have SSH access to the remote machines where you intend to deploy your applications.

3. **Draft a Blueprint**: Use the `ezeth draft` command to create a blueprint. Customize the blueprint file as needed to define the desired configuration and commands.

4. **Execute the Blueprint**: Use the `ezeth exec` command to apply the blueprint to your remote machine. Monitor the process to ensure successful execution.

## Examples

### Drafting a Blueprint

To draft a blueprint named `my-ethereum-app`, run:

```
ezeth draft my-ethereum-app
```

### Executing a Blueprint

To execute the `my-ethereum-app` blueprint on a remote machine, run:

```
ezeth exec my-ethereum-app
```

## Troubleshooting

- **Connection Issues**: Ensure that your SSH configuration is correct and that you have the necessary permissions to access the remote machine.
- **Blueprint Errors**: Verify that the blueprint is correctly formatted and contains valid configurations and commands.
- **Execution Failures**: Check the logs and output messages for any errors that may indicate issues with the remote machine or command execution.

## Conclusion

EZETH simplifies the process of deploying and managing Ethereum-based applications by providing an intuitive interface for drafting and executing blueprints. With its SSH integration and support for Systemd Dynamic Users and Docker, EZETH is a versatile tool for developers and operators alike. For further assistance, refer to the [EZETH support documentation](#) or contact our support team.