# Commands

Authentication:
- Name: auth
- Description: Store the `api_key` on a config file.
- Flags: file

Configuration:
- Name: config
- Description: Display configuration to stdout

Not for now, it is failing raising an unathorized
Admin-Configuration:
- Name: admin-config
- Description: Retrieve or update admin configuration from dev.to
- Subcommands: [retrieve, update]
- Responses: {
    retrive: [200],
    update: [400, 401, 403, 422],
}


# Folder structure
- cmd
- api
- models
- util

Let's implement this first.

# Question during implementation
1. In case of an existing .devto what to do?
Answer: will truncate the existing file and replace the content
2. Where to put this .devto file? On HOME?
3. Need to define the Makefile for compilation

