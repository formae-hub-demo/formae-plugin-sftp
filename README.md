# SFTP Plugin for formae

[![CI](https://github.com/platform-engineering-labs/formae-plugin-sftp/actions/workflows/ci.yml/badge.svg)](https://github.com/platform-engineering-labs/formae-plugin-sftp/actions/workflows/ci.yml)

A formae plugin for managing files on SFTP servers. This plugin was created as part of the [Plugin SDK Tutorial](https://docs.formae.io/plugin-sdk/tutorial/01-scaffold/).

## Supported Resources

| Resource Type | Description |
|---------------|-------------|
| `SFTP::Files::File` | Manages files on an SFTP server |

## Configuration

Configure a target in your forma file:

```pkl
import "@sftp/sftp.pkl"

new formae.Target {
  label = "sftp-server"
  config = new sftp.Config {
    url = "sftp://hostname:22"
  }
}
```

### Credentials

The plugin reads SFTP credentials from environment variables:

| Variable | Description |
|----------|-------------|
| `SFTP_USERNAME` | SFTP username |
| `SFTP_PASSWORD` | SFTP password |

Set these environment variables before starting the formae agent.

## Examples

See the [examples/](examples/) directory for usage examples.

```pkl
import "@sftp/sftp.pkl"

new sftp.File {
  label = "hello"
  path = "/upload/hello.txt"
  content = "Hello from formae!"
  permissions = "0644"
}
```

```bash
# Apply resources
formae apply --mode reconcile examples/basic/main.pkl
```

## License

This plugin is licensed under [FSL-1.1-ALv2](LICENSE).
