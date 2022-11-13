# Helm Plugin

## Requirements

- Helm
- Helm s3 (`helm plugin install https://github.com/hypnoglow/helm-s3.git`)
- Go

## Resources

- https://github.com/chartmuseum/helm-push

## Required Config

The following properties must be defined in secret manager or as environment variables:

- `org`: The Github org for manging the dev envs.
- `github_token`: Access token for a Github user with access to the org.

## Testing

### E2E

`make install`

`make test-helm`
