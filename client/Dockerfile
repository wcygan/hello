# Build stage
FROM denoland/deno:alpine AS builder

# Create app directory
WORKDIR /app

# The base image already has a 'deno' user and group.
# Ensure the app directory is owned by the deno user.
RUN chown deno:deno /app

# Switch to non-root user *before* copying files
USER deno

# Copy dependency manifests
COPY --chown=deno:deno deno.json fresh.config.ts ./

# Cache dependencies. This will cache inside the default location relative to WORKDIR
# Using dev.ts as it might pull more deps than main.ts for Fresh projects
RUN deno cache deps dev.ts --no-lock

# Copy the rest of the application code
COPY --chown=deno:deno . .

# --- Final stage ---
FROM denoland/deno:alpine

# Create app directory
WORKDIR /app

# Copy the entire application directory (including cache) from the build stage
COPY --from=builder --chown=deno:deno /app .

# Switch to non-root user
USER deno

EXPOSE 8000

# Define the runtime command for production using main.ts.
# Use DENO_NO_PROMPT=1 as recommended for non-interactive environments
ENV DENO_NO_PROMPT=1
CMD ["run", "--allow-net", "--allow-read", "main.ts"] 