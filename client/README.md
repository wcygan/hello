# Fresh project

Your new Fresh project is ready to go. You can follow the Fresh "Getting
Started" guide here: https://fresh.deno.dev/docs/getting-started

## Local Development

### Prerequisites

Make sure you have Deno installed. You can install it using one of the commands below or find more options at the [official Deno installation guide](https://deno.land/manual/getting_started/installation).

**Shell (macOS, Linux):**
```sh
curl -fsSL https://deno.land/x/install/install.sh | sh
```

**PowerShell (Windows):**
```powershell
irm https://deno.land/install.ps1 | iex
```

### Running the Client

Navigate to the `client` directory and start the development server:

```bash
deno task start
```

This command uses Deno's task runner (defined in `deno.json`) to start the Fresh application.

The client will be accessible at [http://localhost:8000](http://localhost:8000) by default.
The server will automatically watch for file changes and restart as necessary.
