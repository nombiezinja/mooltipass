# Mooltipass

Mooltipass is a personal project meant to solve the problem of scaffolding every single service from scratch. Think of it as a one-command set up for new services, if you will. It comes with platform-ready features, ready as the foundation for adding business logic:

- Authn with email/password, OAuth SSO (configurable identity providers)
- ABAC with OPA
- Base Openapi definition with rudimentary endpoints
- oapi-codegen set up 
- Observability (customizable logger, wrapper over slog/zero -- configurable; otel)
- DB setup 
- build and release 
- standardized dev experience 

## Usage (WIP)
make push-artifacts SSH_ALIAS=github-public TARGET_REPO=OWNER/REPO