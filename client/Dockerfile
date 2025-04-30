# Use the official Deno image, specifically the alpine variant for smaller size
FROM denoland/deno:alpine

# The base image already has a 'deno' user and group.
# Create app directory owned by the deno user.
RUN mkdir /app && chown deno:deno /app

# Set the working directory
WORKDIR /app

# Switch to non-root user *before* copying files
USER deno

# Copy dependency manifests first for better layer caching
COPY --chown=deno:deno deno.json fresh.config.ts ./

# Cache dependencies. This will cache inside the default location relative to WORKDIR
# Using dev.ts as it might pull more deps than main.ts for Fresh projects
RUN deno cache deps dev.ts --no-lock

# Copy the rest of the application code
COPY --chown=deno:deno . .

EXPOSE 8000

# Define the runtime command for production using main.ts.
# Use DENO_NO_PROMPT=1 as recommended for non-interactive environments
ENV DENO_NO_PROMPT=1
CMD ["deno", "run", "--allow-net", "--allow-read", "--allow-env", "main.ts"] 