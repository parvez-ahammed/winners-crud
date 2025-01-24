# Day 2

## Initial planning

There will be theee environments. `Dev`, `QA`, `Prod`.

For the `Dev` environment the setup will be as follows:

```text
crud app (service) + sql server (db)
```

## Resource Group

The name resource group name should be clear to find out details about the resource group an example can be. Creating a resource group does not cost anything. The resource group is a logical container for resources. The resource group can contain resources like web apps, databases, and storage accounts.

```jsx
devops - learning - dev - rg;
```

Region is required to be set for the resource group.This can be the closest region to the team. For the current context the `(Asia Pacific) Central Inda` region is the closest.

## First Resource

### Details

- Web App
- Resource Group: `devops - learning - dev - rg`
- Name: `winners-crud-client-dev`
- Runtime Stack: Node 20 LTS

### App Service Plan

- Name: devop-learning-dev-linux-service-plan
- operating system: Linux
- Region: Central India
- SKU: Free
- ACU: Shared

## Prices

- prices vary based on region on OS

## Environments

- Dev ( Dhaka team)
- QA (QA memeber)
- Prod(master) (team lead read)

dev environment
