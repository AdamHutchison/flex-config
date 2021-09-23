# Flux Config
The Flux Config package is designed to handle the loading and retrival of configurations values from within the Flux framework.

Flux config will load all values set in the `.env` file located in the root of the project aswell as any values contained within a `config/config.yaml` file.

install flex config using `go get github.com/AdamHutchison/flux-config`.


## Loading and Retrieving Config Values
if we have a `config/config.yaml` containing the following:
```yaml
app:
    name: "Flux App"
```

Then we can load and access the values like this:
```go
import(
    config "github.com/AdamHutchison/flux-config"
)

func main() {
    config.Load()
    name := config.Get("app.name")
}
```

## Setting Config Values Using Env Variables
Env variables may be used in the `config.yaml` file using the `${ENV_VARIABLE_NAME}` syntax:

.env
```bash
APP_NAME="Flux App"
```

config.yaml
```yaml
app:
    name: ${APP_NAME}
```

